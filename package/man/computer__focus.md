#### Description

Bring a running application to the foreground (always real foreground).

`aux4 computer focus` runs `aux4 computer <provider> focus` on the provider for this operating system (`mac` on macOS, `linux` on Linux, `win` on Windows), or the one given with `--provider` / `AUX4_COMPUTER_PROVIDER`. Only the flags you give are passed on, unchanged, so the provider's defaults apply (shown below as "provider default"); its JSON output, errors and exit code are passed straight through. A missing provider, or one that speaks another contract version, fails with a JSON error that names the package to install or update.

It sends input or changes state, so it is refused (exit code `1`, `"stopped": true`) while the stop switch is on (`aux4 computer stop`), and the provider records it in the action log (`aux4 computer log`).

Platform specifics: `aux4 computer notes` and the provider's own page, e.g. `aux4 aux4 man computer mac focus`.

#### Usage

```bash
aux4 computer focus <app> [--provider <value>]
```

--app       Application name (required)
--provider  Provider to use: mac, linux or win (default: the one for this operating system; env AUX4_COMPUTER_PROVIDER)

#### Example

```bash
aux4 computer focus "TextEdit"
```

```json
{"focused":"TextEdit","mode":"foreground","success":true}
```
