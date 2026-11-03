#### Description

Turns on the **shared stop switch**: creates the `STOP` file in the state directory. While it exists, every command
that sends input or changes state — on every provider, whether called as `aux4 computer <command>` or directly as
`aux4 computer <provider> <command>` — refuses with exit code `1` and `"stopped": true`, before touching anything.
Read-only commands (`state`, `read`, `find`, `screenshot`, `ocr`, `wait`, `menu` without `--path`, `dialog` without
`--press`, `window list`, `clipboard read`, `log`...) keep working, so an agent can still look but not act.

Use it to halt an agent immediately from any terminal; `aux4 computer resume` turns it off. The environment variable
`AUX4_COMPUTER_STOP=1` has the same effect for the processes that see it. `stop` is recorded in the action log, and
so is every refused command.

| OS | STOP file |
|---|---|
| macOS | `~/Library/Application Support/aux4-computer/STOP` |
| Linux | `${XDG_STATE_HOME:-~/.local/state}/aux4-computer/STOP` |
| Windows | `%LOCALAPPDATA%\aux4-computer\STOP` |

On macOS it also creates the stop file of `aux4/computer-mac` releases that predate the shared state directory
(`~/Library/Application Support/aux4-computer-mac/STOP`); `files` lists every file written. `AUX4_COMPUTER_STATE_DIR`
replaces the state directory (then only that directory is used).

#### Usage

```bash
aux4 computer stop
```

#### Example

```bash
aux4 computer stop
aux4 computer click --app Calculator --x 34 --y 211
```

```json
{"file":"/Users/sally/Library/Application Support/aux4-computer/STOP","files":["/Users/sally/Library/Application Support/aux4-computer/STOP","/Users/sally/Library/Application Support/aux4-computer-mac/STOP"],"stopped":true}
```

```json
{"command":"click","error":"Stopped: aux4 computer is not sending input or changing state (/Users/sally/Library/Application Support/aux4-computer/STOP exists). Run 'aux4 computer resume' to allow it again.","provider":"mac","stopped":true}
```
