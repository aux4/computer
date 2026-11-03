#### Description

Shows the most recent entries of the **shared action log**, oldest first. Every provider appends one JSON line to
`actions.jsonl` in the state directory for each command that sends input or changes state — whether it succeeded,
failed or was refused by the stop switch — and this package adds its own refusals and `stop` / `resume`. Read-only
commands and dry runs are not recorded.

Each entry has `ts` (ISO 8601), `provider`, `command`, `app`, `args` (the arguments given; `--text` and `--value`
are cut to 16 characters plus their length, since typed text can be private), `mode`, `result` (`ok` or `error`),
`error`, and, depending on the provider, `method`, `fallbackReason` and `ms`. The file is private to the user (mode
`0600`) and rotated at 10MB; the previous file (`actions.jsonl.1`) is read too.

On macOS the log of `aux4/computer-mac` under `~/Library/Logs` (including releases that predate the shared state
directory) is merged in. `paths` lists the files read, `path` is the shared log file. The output also says whether
the stop switch is on (`stopped`, `stopReason`).

#### Usage

```bash
aux4 computer log [--limit 20]
```

--limit  Number of recent entries to show, oldest first (default: `20`)

#### Example

```bash
aux4 computer log --limit 2
```

```json
{"entries":[{"app":"TextEdit","args":{"all":"true","app":"TextEdit","mode":"background"},"command":"select","method":"accessibility","mode":"background","ms":62,"provider":"mac","result":"ok","ts":"2026-11-02T22:52:10.114Z"},{"app":"Calculator","args":{"app":"Calculator","x":"34","y":"211"},"command":"click","error":"stopped: /Users/sally/Library/Application Support/aux4-computer/STOP exists","provider":"mac","result":"error","source":"core","stopped":true,"ts":"2026-11-02T22:52:12.208Z"}],"path":"/Users/sally/Library/Application Support/aux4-computer/actions.jsonl","paths":["/Users/sally/Library/Application Support/aux4-computer/actions.jsonl"],"stopped":true,"stopReason":"/Users/sally/Library/Application Support/aux4-computer/STOP exists"}
```
