// aux4-computer is the helper binary behind aux4/computer, the OS-independent core of the computer-use
// packages. The command surface itself (names, flags, help) lives in package/.aux4; every generic command
// calls this binary with the command path, the --provider value and the flags that were actually given
// (params() drops empty ones), and the binary:
//
//   - picks the provider for this OS (darwin -> mac, linux -> linux, windows -> windows) unless overridden,
//   - checks that the provider is installed and speaks the same contract version (via its `info` command),
//   - refuses state-changing commands while the shared stop switch is on (and records the refusal),
//   - runs `aux4 computer <provider> <command> <flags...>` with an explicit argv (no shell) and passes
//     stdin/stdout/stderr and the exit code straight through.
//
// It also implements the commands the core owns: providers, stop, resume, log, notes, conformance.
package main

import (
	"fmt"
	"os"
)

// ContractVersion is the provider contract this core implements (see the README, "Provider contract").
const ContractVersion = 1

func main() {
	if len(os.Args) < 2 {
		fail(map[string]any{"error": "Usage: aux4-computer <dispatch|providers|stop|resume|log|notes|conformance> ..."})
	}
	args := os.Args[2:]
	switch os.Args[1] {
	case "dispatch":
		os.Exit(cmdDispatch(args))
	case "providers":
		cmdProviders(args)
	case "stop":
		cmdStop()
	case "resume":
		cmdResume()
	case "log":
		cmdLog(args)
	case "notes":
		cmdNotes(args)
	case "conformance":
		os.Exit(cmdConformance(args))
	default:
		fail(map[string]any{"error": fmt.Sprintf("Unknown command '%s'", os.Args[1])})
	}
}
