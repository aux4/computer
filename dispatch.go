package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
)

// stateChanging lists the commands that send input or change state. menu and dialog only act with
// --path / --press, window and clipboard depend on the sub-command (see isStateChanging).
var stateChanging = map[string]bool{
	"click": true, "move": true, "drag": true, "type": true, "shortcut": true, "scroll": true,
	"open": true, "focus": true, "quit": true, "hide": true, "unhide": true,
	"select": true, "set": true,
	"window move": true, "window resize": true, "window minimize": true, "window restore": true,
	"window close": true, "window fullscreen": true, "window raise": true,
	"clipboard copy": true, "clipboard paste": true,
}

func isStateChanging(command string, flags map[string]string) bool {
	switch command {
	case "menu":
		return flags["path"] != ""
	case "dialog":
		return flags["press"] != ""
	}
	return stateChanging[command]
}

// parseFlags reads the `--name value` pairs produced by params() (empty values are already dropped).
func parseFlags(argv []string) (map[string]string, error) {
	flags := map[string]string{}
	for i := 0; i < len(argv); i++ {
		a := argv[i]
		if !strings.HasPrefix(a, "--") || len(a) < 3 {
			return nil, fmt.Errorf("unexpected argument '%s'", a)
		}
		name := a[2:]
		if i+1 >= len(argv) {
			return nil, fmt.Errorf("flag --%s has no value", name)
		}
		flags[name] = argv[i+1]
		i++
	}
	return flags, nil
}

func logArgs(flags map[string]string) map[string]any {
	out := map[string]any{}
	for k, v := range flags {
		if k == "text" || k == "value" {
			out[k] = truncateForLog(v)
		} else {
			out[k] = v
		}
	}
	return out
}

// cmdDispatch: dispatch <command path, ':'-separated> <provider or ""> [--flag value ...]
func cmdDispatch(args []string) int {
	if len(args) < 2 {
		fail(map[string]any{"error": "Usage: aux4-computer dispatch <command> <provider> [--flag value ...]"})
	}
	path := strings.Split(args[0], ":")
	command := strings.Join(path, " ")
	flagArgv := args[2:]
	flags, err := parseFlags(flagArgv)
	if err != nil {
		fail(map[string]any{"error": fmt.Sprintf("Invalid arguments for '%s': %v", command, err)})
	}

	provider := resolveProvider(args[1])

	// info/notes are how the contract is discovered, so they only need the provider to be installed.
	needContract := command != "info" && command != "notes"
	requireProvider(provider, needContract)

	if isStateChanging(command, flags) {
		if reason := stopReason(); reason != "" {
			entry := map[string]any{
				"ts": nowTS(), "provider": provider, "command": command, "args": logArgs(flags),
				"result": "error", "error": "stopped: " + reason, "stopped": true, "source": "core",
			}
			if flags["app"] != "" {
				entry["app"] = flags["app"]
			}
			if flags["mode"] != "" {
				entry["mode"] = flags["mode"]
			}
			appendActionLog(entry)
			fail(map[string]any{
				"error":    fmt.Sprintf("Stopped: aux4 computer is not sending input or changing state (%s). Run 'aux4 computer resume' to allow it again.", reason),
				"stopped":  true,
				"command":  command,
				"provider": provider,
			})
		}
	}

	// ocr --format text is reformatted by the core after the provider replies; providers never see --format
	// at all (they only know json), so it's stripped from both the flag map and the raw argv before the
	// provider's argv is built below.
	ocrFormat := ""
	if command == "ocr" {
		ocrFormat = flags["format"]
		delete(flags, "format")
		flagArgv = removeFlagPair(flagArgv, "format")
	}

	argv := append([]string{"computer", provider}, path...)
	// --name=value keeps values that start with '-' (negative coordinates, "--text") from being read as flags.
	for i := 0; i+1 < len(flagArgv); i += 2 {
		argv = append(argv, flagArgv[i]+"="+flagArgv[i+1])
	}

	if ocrFormat == "text" {
		return execOcrText(argv)
	}
	return execPassthrough(argv)
}

// removeFlagPair drops a "--name value" pair from a flat argv slice (as produced by params()), leaving
// everything else in order.
func removeFlagPair(argv []string, name string) []string {
	target := "--" + name
	out := make([]string, 0, len(argv))
	for i := 0; i < len(argv); i++ {
		if argv[i] == target && i+1 < len(argv) {
			i++
			continue
		}
		out = append(out, argv[i])
	}
	return out
}

// execPassthrough runs aux4 with an explicit argv (no shell parses it, so no value can break out), wires
// stdin/stdout/stderr straight through and returns the child's exit code.
func execPassthrough(argv []string) int {
	cmd := exec.Command(aux4Binary(), argv...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// Let the child handle Ctrl-C itself; we only wait for it.
	signal.Ignore(os.Interrupt)
	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if code := exitErr.ExitCode(); code >= 0 {
				return code
			}
			return 1
		}
		fail(map[string]any{"error": fmt.Sprintf("Could not run aux4: %v", err)})
	}
	return 0
}
