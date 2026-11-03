#### Description

Read the current clipboard text (no side effects).

`aux4 computer clipboard read` runs `aux4 computer <provider> clipboard read` on the provider for this operating system (`mac` on macOS, `linux` on Linux, `win` on Windows), or the one given with `--provider` / `AUX4_COMPUTER_PROVIDER`. Only the flags you give are passed on, unchanged, so the provider's defaults apply (shown below as "provider default"); its JSON output, errors and exit code are passed straight through. A missing provider, or one that speaks another contract version, fails with a JSON error that names the package to install or update.

This command is read-only: it keeps working while the stop switch is on.

Platform specifics: `aux4 computer notes` and the provider's own page, e.g. `aux4 aux4 man computer mac clipboard read`.

#### Usage

```bash
aux4 computer clipboard read [--provider <value>]
```

--provider  Provider to use: mac, linux or win (default: the one for this operating system; env AUX4_COMPUTER_PROVIDER)

#### Example

```bash
aux4 computer clipboard read
```

```json
{"text":"hello"}
```
