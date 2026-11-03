package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// cmdConformance: conformance <packageDir> <provider or ""> <dir>
//
// Copies the contract's conformance suite into <dir>/.aux4-computer-conformance and runs `aux4 test run`
// there with AUX4_COMPUTER_PROVIDER set. Running it from a provider's package/ directory makes aux4 load
// that provider's local .aux4, so a provider's CI can check the working copy without installing it.
func cmdConformance(args []string) int {
	if len(args) < 3 {
		fail(map[string]any{"error": "Usage: aux4-computer conformance <packageDir> <provider> <dir>"})
	}
	packageDir, given, dir := args[0], args[1], args[2]
	if dir == "" {
		dir = "."
	}
	provider := resolveProvider(given)
	// Only a provider that speaks this contract may run the suite: it relies on AUX4_COMPUTER_DRY_RUN and
	// AUX4_COMPUTER_STATE_DIR, and a provider that ignores them would send real input or touch the real switch.
	p := probeProvider(provider)
	switch p.Status {
	case "missing":
		fail(missingError(provider))
	case "ok":
	default:
		fail(map[string]any{
			"error":    fmt.Sprintf("Refusing to run the conformance suite against '%s' (%s): status %s (%s). The suite needs a provider that reports computer contract %d in 'info', since it relies on AUX4_COMPUTER_DRY_RUN and AUX4_COMPUTER_STATE_DIR to stay headless and away from the real stop switch.", provider, p.Package, p.Status, p.Error, ContractVersion),
			"provider": provider,
			"package":  p.Package,
			"status":   p.Status,
		})
	}
	src := filepath.Join(packageDir, "conformance")
	files, err := filepath.Glob(filepath.Join(src, "*.test.md"))
	if err != nil || len(files) == 0 {
		fail(map[string]any{"error": fmt.Sprintf("No conformance tests found in %s", src)})
	}
	absDir, err := filepath.Abs(dir)
	if err != nil {
		fail(map[string]any{"error": fmt.Sprintf("Invalid --dir '%s': %v", dir, err)})
	}
	target := filepath.Join(absDir, ".aux4-computer-conformance")
	_ = os.RemoveAll(target)
	if err := os.MkdirAll(target, 0o755); err != nil {
		fail(map[string]any{"error": fmt.Sprintf("Could not create %s: %v", target, err)})
	}
	defer os.RemoveAll(target)
	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			fail(map[string]any{"error": fmt.Sprintf("Could not read %s: %v", f, err)})
		}
		if err := os.WriteFile(filepath.Join(target, filepath.Base(f)), data, 0o644); err != nil {
			fail(map[string]any{"error": fmt.Sprintf("Could not write %s: %v", target, err)})
		}
	}
	fmt.Fprintf(os.Stderr, "Running the computer contract %d conformance suite against provider '%s' (%d files)\n", ContractVersion, provider, len(files))
	cmd := exec.Command(aux4Binary(), "test", "run", "--dir", target)
	cmd.Dir = target
	env := []string{}
	for _, e := range os.Environ() {
		if !strings.HasPrefix(e, "AUX4_COMPUTER_PROVIDER=") {
			env = append(env, e)
		}
	}
	cmd.Env = append(env, "AUX4_COMPUTER_PROVIDER="+provider)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() > 0 {
			return exitErr.ExitCode()
		}
		return 1
	}
	return 0
}
