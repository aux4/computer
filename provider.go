package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"
)

// knownProviders maps a provider name to the package that implements it.
var knownProviders = []struct{ name, pkg, goos string }{
	{"mac", "aux4/computer-mac", "darwin"},
	{"linux", "aux4/computer-linux", "linux"},
	{"win", "aux4/computer-windows", "windows"},
}

var providerNamePattern = regexp.MustCompile(`^[a-z0-9][a-z0-9-]*$`)

// coreCommands are the names in the core's `computer` profile. A provider can't use one of them as its name:
// `aux4 computer <name>` would run the core command instead (that is why the Windows provider is `win`).
var coreCommands = map[string]bool{
	"providers": true, "info": true, "notes": true, "permissions": true, "screenshot": true, "ocr": true,
	"wait": true, "read": true, "find": true, "click": true, "move": true, "drag": true, "type": true,
	"shortcut": true, "scroll": true, "cursor": true, "screen": true, "apps": true, "windows": true,
	"quit": true, "hide": true, "unhide": true, "open": true, "focus": true, "state": true, "menu": true,
	"select": true, "set": true, "dialog": true, "window": true, "clipboard": true, "stop": true,
	"resume": true, "log": true, "conformance": true,
}

// validateProviderName fails unless name can be a provider name.
func validateProviderName(name string) {
	if name == "windows" {
		fail(map[string]any{"error": "Invalid --provider 'windows': the Windows provider is named 'win' ('windows' is the command that lists windows)"})
	}
	if !providerNamePattern.MatchString(name) || coreCommands[name] {
		fail(map[string]any{"error": fmt.Sprintf("Invalid --provider '%s': use a provider name such as mac, linux or win", name)})
	}
}

func providerPackage(name string) string {
	for _, p := range knownProviders {
		if p.name == name {
			return p.pkg
		}
	}
	return "aux4/computer-" + name
}

// defaultProvider is the provider for this OS, or "" when the OS has none.
func defaultProvider() string {
	goos := hostOS()
	for _, p := range knownProviders {
		if p.goos == goos {
			return p.name
		}
	}
	return ""
}

// resolveProvider returns the provider to use: --provider / AUX4_COMPUTER_PROVIDER, else the one for this OS.
func resolveProvider(given string) string {
	given = strings.TrimSpace(given)
	if given == "" {
		p := defaultProvider()
		if p == "" {
			fail(map[string]any{
				"error":     fmt.Sprintf("No computer provider for this operating system (%s). Pass --provider <name> to use one explicitly.", hostOS()),
				"os":        hostOS(),
				"providers": []string{"mac", "linux", "win"},
			})
		}
		return p
	}
	validateProviderName(given)
	return given
}

// runAux4 runs `aux4 <args...>` with stdin closed and a deadline, capturing its output.
func runAux4(timeout time.Duration, args ...string) (stdout, stderr string, code int) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, aux4Binary(), args...)
	var out, errb bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errb
	err := cmd.Run()
	code = 0
	if err != nil {
		code = 1
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() >= 0 {
			code = exitErr.ExitCode()
		}
		if ctx.Err() == context.DeadlineExceeded {
			code = 124
			errb.WriteString(fmt.Sprintf("timed out after %s", timeout))
		}
	}
	return out.String(), errb.String(), code
}

func aux4Binary() string {
	if v := os.Getenv("AUX4_COMPUTER_AUX4"); v != "" {
		return v
	}
	return "aux4"
}

func commandNotFound(stdout, stderr string) bool {
	return strings.Contains(stdout+stderr, "Command not found")
}

// probe describes what was learned about one provider.
type probe struct {
	Name      string         `json:"provider"`
	Package   string         `json:"package"`
	Installed bool           `json:"installed"`
	Status    string         `json:"status"` // ok | legacy | mismatch | error | missing
	Contract  any            `json:"contract"`
	Info      map[string]any `json:"info,omitempty"`
	Error     string         `json:"error,omitempty"`
}

const probeTimeout = 20 * time.Second

func installed(name string) bool {
	_, _, code := runAux4(probeTimeout, "computer", name, "--help")
	return code == 0
}

// probeProvider calls `aux4 computer <name> info` and classifies the provider.
func probeProvider(name string) probe {
	p := probe{Name: name, Package: providerPackage(name)}
	stdout, stderr, code := runAux4(probeTimeout, "computer", name, "info")
	if code != 0 {
		if !installed(name) {
			p.Status = "missing"
			return p
		}
		p.Installed = true
		if commandNotFound(stdout, stderr) {
			// Installed but predates the contract: it has no `info` command.
			p.Status = "legacy"
			p.Error = "provider has no 'info' command (it predates computer contract " + fmt.Sprint(ContractVersion) + "); update " + p.Package
			return p
		}
		p.Status = "error"
		p.Error = strings.TrimSpace(firstNonEmpty(stderr, stdout))
		return p
	}
	p.Installed = true
	var info map[string]any
	if err := json.Unmarshal([]byte(strings.TrimSpace(stdout)), &info); err != nil {
		p.Status = "mismatch"
		p.Error = "'info' did not return a JSON object"
		return p
	}
	p.Info = info
	p.Contract = info["contract"]
	if n, ok := info["contract"].(float64); ok && int(n) == ContractVersion && n == float64(int(n)) {
		p.Status = "ok"
	} else {
		p.Status = "mismatch"
		p.Error = fmt.Sprintf("provider implements computer contract %v; aux4/computer implements contract %d", contractText(info["contract"]), ContractVersion)
	}
	return p
}

func contractText(v any) string {
	if v == nil {
		return "(none)"
	}
	return fmt.Sprint(v)
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if strings.TrimSpace(v) != "" {
			return v
		}
	}
	return ""
}

func missingError(name string) map[string]any {
	pkg := providerPackage(name)
	return map[string]any{
		"error":    fmt.Sprintf("No computer provider '%s' is installed. Install it with: aux4 aux4 pkger install %s", name, pkg),
		"provider": name,
		"package":  pkg,
	}
}

// requireProvider fails unless the provider is installed and speaks this core's contract.
// Providers that predate the contract (no `info`) are allowed through: their contract is unknown.
func requireProvider(name string, needContract bool) probe {
	if !needContract {
		if !installed(name) {
			fail(missingError(name))
		}
		return probe{Name: name, Package: providerPackage(name), Installed: true, Status: "unchecked"}
	}
	p := probeProvider(name)
	switch p.Status {
	case "missing":
		fail(missingError(name))
	case "mismatch":
		fail(map[string]any{
			"error":    fmt.Sprintf("Computer provider '%s' (%s) is not compatible: %s. Update both with: aux4 aux4 pkger install aux4/computer %s", name, p.Package, p.Error, p.Package),
			"provider": name,
			"package":  p.Package,
			"contract": p.Contract,
			"expected": ContractVersion,
		})
	case "error":
		fail(map[string]any{
			"error":    fmt.Sprintf("Computer provider '%s' failed to describe itself ('info'): %s", name, p.Error),
			"provider": name,
			"package":  p.Package,
		})
	}
	return p
}

func cmdProviders(args []string) {
	extra := ""
	if len(args) > 0 {
		extra = strings.TrimSpace(args[0])
	}
	names := []string{}
	for _, k := range knownProviders {
		names = append(names, k.name)
	}
	if extra != "" {
		validateProviderName(extra)
		seen := false
		for _, n := range names {
			seen = seen || n == extra
		}
		if !seen {
			names = append(names, extra)
		}
	}
	def := defaultProvider()
	installedList := []map[string]any{}
	notInstalled := []map[string]any{}
	for _, n := range names {
		p := probeProvider(n)
		if !p.Installed {
			notInstalled = append(notInstalled, map[string]any{"provider": n, "package": p.Package})
			continue
		}
		entry := map[string]any{
			"provider": p.Name,
			"package":  p.Package,
			"status":   p.Status,
			"contract": p.Contract,
			"default":  n == def,
		}
		if p.Info != nil {
			entry["info"] = p.Info
		}
		if p.Error != "" {
			entry["error"] = p.Error
		}
		installedList = append(installedList, entry)
	}
	platform := hostOS() + "/" + runtime.GOARCH
	out := map[string]any{
		"contract":     ContractVersion,
		"platform":     platform,
		"default":      def,
		"providers":    installedList,
		"notInstalled": notInstalled,
	}
	if def != "" {
		found := false
		for _, e := range installedList {
			found = found || e["provider"] == def
		}
		if !found {
			out["hint"] = fmt.Sprintf("No provider for this OS is installed. Install it with: aux4 aux4 pkger install %s", providerPackage(def))
		}
	}
	printJSON(out)
}

// cmdNotes prints the current provider's `notes` (markdown) for the AI skill. It never fails: the skill
// prompt should still be usable when the provider is missing or predates the contract.
func cmdNotes(args []string) {
	given := ""
	if len(args) > 0 {
		given = args[0]
	}
	name := strings.TrimSpace(given)
	if name == "" {
		name = defaultProvider()
	}
	if name == "" {
		fmt.Printf("\n## Platform notes\n\nThis operating system (%s) has no computer provider. Pass `--provider <name>` to every command to use one explicitly.\n", hostOS())
		return
	}
	validateProviderName(name)
	pkg := providerPackage(name)
	stdout, stderr, code := runAux4(probeTimeout, "computer", name, "notes")
	fmt.Printf("\n## Platform notes: %s (%s)\n\n", name, pkg)
	if code == 0 && strings.TrimSpace(stdout) != "" {
		fmt.Print(strings.TrimRight(stdout, "\n") + "\n")
		return
	}
	if !installed(name) {
		fmt.Printf("The provider for this computer is not installed, so every command will fail until it is: `aux4 aux4 pkger install %s`.\n", pkg)
		return
	}
	if commandNotFound(stdout, stderr) {
		fmt.Printf("The installed %s predates the provider contract and has no `notes` command. Run `aux4 computer %s --help` and `aux4 aux4 pkger man %s` for its platform specifics, and update it when you can.\n", pkg, name, pkg)
		return
	}
	fmt.Printf("`aux4 computer %s notes` failed: %s\n", name, strings.TrimSpace(firstNonEmpty(stderr, stdout)))
}
