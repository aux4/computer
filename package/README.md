# aux4/computer

Computer-use for AI agents on any operating system. `aux4 computer <command>` is one command surface — screenshot,
read the accessibility tree, find, OCR, click, type, shortcuts, scroll, drag, menus, dialogs, windows, clipboard and
app control — that runs on the provider for the current OS: `aux4/computer-mac` on macOS, `aux4/computer-linux` on
Linux and `aux4/computer-windows` on Windows. The package also owns what every provider shares: the stop switch, the
action log, the provider contract and its conformance suite, and an AI skill that teaches agents the method.

## Installation

```bash
aux4 aux4 pkger install aux4/computer
```

Then install the provider for your operating system (providers depend on `aux4/computer`, so installing a provider
also installs this package):

```bash
aux4 aux4 pkger install aux4/computer-mac      # macOS
aux4 aux4 pkger install aux4/computer-linux    # Linux
aux4 aux4 pkger install aux4/computer-windows  # Windows
```

## Quick Start

```bash
aux4 computer providers                         # which provider this OS uses, its contract and capabilities
aux4 computer state --app Calculator            # controls of the app, each with a ref
aux4 computer find "7" --app Calculator         # elements whose label contains "7"
aux4 computer click --element 0.0.3.1 --app Calculator
aux4 computer screenshot --app Calculator --annotate true
```

Every command prints one JSON object on stdout; errors are a JSON object on stderr (`{"error": "..."}`) with exit
code `1`.

## How commands reach the provider

`aux4 computer <command> [flags]` runs `aux4 computer <provider> <command> [flags]`:

- **Provider choice** — `mac` on macOS, `linux` on Linux, `win` on Windows. `--provider <name>` (or the
  `AUX4_COMPUTER_PROVIDER` environment variable) picks one explicitly.
- **Flags** — only the flags you give are passed on, unchanged, so the provider's own defaults apply. Values reach
  the provider byte for byte: quotes, `$(...)`, spaces and leading dashes are never re-interpreted. Negative numbers
  need `=` (`--x=-200`).
- **Output and exit code** — the provider's stdout, stderr and exit code are passed straight through.
- **Missing provider** — a clear error names the package to install:

  ```json
  {"error":"No computer provider 'linux' is installed. Install it with: aux4 aux4 pkger install aux4/computer-linux","package":"aux4/computer-linux","provider":"linux"}
  ```

- **Incompatible provider** — a provider whose `info` reports another contract version is refused, with the
  contract it reports (`contract`) and the one this package needs (`expected`). A provider that predates the contract
  (it has no `info` command) is still used; `providers` lists it as `legacy`.
- **Stop switch** — while it's on, input and state-changing commands are refused before they reach the provider.

The Windows provider is named `win` because `aux4 computer windows` is the command that lists windows.

## Commands

| Command | Description |
|---|---|
| `aux4 computer providers` | Installed providers, their status (`ok`, `legacy`, `mismatch`, `error`), contract, capabilities |
| `aux4 computer info` | The current provider's info: name, contract, platform, version, capabilities, permissions |
| `aux4 computer notes` | The current provider's platform notes for agents (markdown) |
| `aux4 computer permissions` | Permissions the platform needs and whether each is granted |
| `aux4 computer state` | App, windows, focused element, dialog and a compact control tree with refs |
| `aux4 computer read` | The full accessibility tree of an app |
| `aux4 computer find` | Elements whose label or value contains a text |
| `aux4 computer screenshot` | Screenshot of a display or an app's window; `--annotate true` numbers the elements |
| `aux4 computer ocr` | Recognize text in an app window, image or the screen |
| `aux4 computer wait` | Wait until a text or role appears (or disappears) |
| `aux4 computer click` | Click coordinates or press an element |
| `aux4 computer move` | Move the mouse cursor |
| `aux4 computer drag` | Drag from one point to another |
| `aux4 computer type` | Type text into the focused element |
| `aux4 computer shortcut` | Press a keyboard shortcut |
| `aux4 computer scroll` | Scroll at a point |
| `aux4 computer cursor` | The cursor position |
| `aux4 computer screen` | Displays: index, size, scale, origin |
| `aux4 computer apps` | Running apps |
| `aux4 computer windows` | On-screen windows |
| `aux4 computer open` | Open an app, file or URL |
| `aux4 computer focus` | Bring an app to the front |
| `aux4 computer quit` | Quit an app |
| `aux4 computer hide` / `unhide` | Hide or show an app |
| `aux4 computer menu` | List or press menu bar items |
| `aux4 computer select` | Select text in an element |
| `aux4 computer set` | Set a value, checkbox or pop-up option |
| `aux4 computer dialog` | Read or answer an open dialog |
| `aux4 computer window list\|move\|resize\|minimize\|restore\|close\|fullscreen\|raise` | Window control |
| `aux4 computer clipboard copy\|paste\|read` | Clipboard |
| `aux4 computer stop` / `resume` | Turn the shared stop switch on or off |
| `aux4 computer log` | Recent entries of the shared action log |
| `aux4 computer conformance` | Run the provider contract conformance suite against a provider |
| `aux4 ai skill computer prompt` | The computer-use method for agents plus the current provider's notes |

Every command except `stop`, `resume` and `log` takes `--provider`. The flags of each
command are those of the provider commands; see `aux4 aux4 man computer <command>`.

## Stop switch and action log

The stop switch and the action log are shared by this package and every provider, so a direct provider call
(`aux4 computer mac click ...`) can't bypass them.

- `aux4 computer stop` creates the `STOP` file in the state directory. While it exists — or while the environment
  variable `AUX4_COMPUTER_STOP=1` is set — every command that sends input or changes state is refused with exit code
  `1` and `"stopped": true`. Read-only commands (`state`, `find`, `screenshot`, `ocr`, `wait`, `menu` without
  `--path`, `dialog` without `--press`, `window list`, `clipboard read`...) keep working.
- `aux4 computer resume` removes it.
- `aux4 computer log` shows the most recent entries of `actions.jsonl` — one JSON line per input or state-changing
  action, from any provider: `ts`, `provider`, `command`, `app`, `args` (typed text and values are cut to 16
  characters plus their length), `mode`, `result`, `error`. The file is private to the user (mode `0600`) and rotated
  at 10MB.

| OS | State directory |
|---|---|
| macOS | `~/Library/Application Support/aux4-computer` |
| Linux | `${XDG_STATE_HOME:-~/.local/state}/aux4-computer` |
| Windows | `%LOCALAPPDATA%\aux4-computer` |

On macOS `stop` also turns on the switch of `aux4/computer-mac` releases that predate the shared state directory,
and `log` also reads their log under `~/Library/Logs`.

## AI skill

```bash
aux4 ai skill computer prompt
```

Prints the method for agents — state, find, OCR, annotate, act by ref, menu, set, select, wait; background versus
foreground and the fallback; coordinates; stop and log; the silent data-loss warning for background edits — followed
by the current provider's platform notes (`aux4 computer <provider> notes`). `--provider` picks another provider's
notes.

## Provider contract

A provider is a package (`aux4/computer-<os>`) that adds `aux4 computer <provider> <command>` for the whole command
surface above plus:

- `aux4 computer <provider> info` — `{"provider":"linux","contract":1,"platform":"linux/amd64","version":"x.y.z","capabilities":{...},"permissions":{...}}`
- `aux4 computer <provider> notes` — markdown platform notes for agents.

It follows the same JSON shapes, `--mode auto|background|foreground` semantics, coordinates, stop switch and action
log, and honours the test switches below. `aux4/computer-mac` is the reference implementation. This package
implements contract version `1`.

## Conformance suite

The suite checks a provider against the contract without sending any real input: `info` fields and contract version,
`permissions` / `screen` / `cursor` / `apps` / `windows` shapes, argument validation and the JSON error shape,
dry-run background/foreground decisions, the shared stop switch, and the action log line shape. It runs on any OS.

With the provider installed (or from its `package/` directory, so the working copy's `.aux4` is used without
installing it):

```bash
aux4 aux4 pkger install aux4/computer aux4/test
aux4 computer conformance --provider linux
```

`aux4 computer conformance` copies the suite into `.aux4-computer-conformance/` under `--dir` (default: the current
directory), runs `aux4 test run` there with `AUX4_COMPUTER_PROVIDER` set, removes the copy and exits with the test
result. It refuses a provider whose `info` doesn't report contract `1`, since the suite relies on the test switches to
stay headless.

In a provider's GitHub workflow, add a step after the `aux4/action@v1` test step — that step already installed aux4
(on `PATH` for later steps), `aux4/test`, the provider and its dependency `aux4/computer`:

```yaml
- name: Test
  uses: aux4/action@v1
  with:
    command: test

- name: Conformance
  run: aux4 computer conformance --provider linux
```

## Environment variables

| Variable | Effect |
|---|---|
| `AUX4_COMPUTER_PROVIDER` | Provider to use instead of the one for this OS (same as `--provider`) |
| `AUX4_COMPUTER_STOP=1` | Stop switch for the processes that see it |
| `AUX4_COMPUTER_STATE_DIR` | Use this directory for the `STOP` file and `actions.jsonl` instead of the per-OS one |
| `AUX4_COMPUTER_DRY_RUN=1` | Providers decide how an action would run without sending it (test switch) |
| `AUX4_COMPUTER_DENY_PERMISSIONS=1` | Providers behave as if no permission were granted (test switch) |
| `AUX4_COMPUTER_SIMULATE_BACKGROUND=unavailable` | Providers behave as if the background path were unavailable (test switch) |

## Development notes

The command surface is declared in `package/.aux4` (each command lists the provider's flags with empty defaults, so
`params()` passes on only the flags that were given and `--help` documents them); a small Go helper,
`aux4-computer`, does the rest. A helper rather than pure DSL because the core must: work the same on Windows (no
POSIX tools to test files, compute per-OS state paths or set file modes), parse provider `info` JSON to check the
contract version, build the `providers` and `log` JSON, check the stop switch, and run the provider with an explicit
argv — `--name=value` pairs, no shell — so no flag value can be re-interpreted on the way. It cross-compiles to the
usual 8 platform/arch targets.

```bash
aux4 build            # all platforms into package/dist
go build -o package/aux4-computer . && go test ./...
cd package && aux4 test run --dir test
```

## License

Apache-2.0. See [LICENSE](./LICENSE).
