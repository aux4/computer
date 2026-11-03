#### Description

Find UI elements by text and/or role.

`aux4 computer find` runs `aux4 computer <provider> find` on the provider for this operating system (`mac` on macOS, `linux` on Linux, `win` on Windows), or the one given with `--provider` / `AUX4_COMPUTER_PROVIDER`. Only the flags you give are passed on, unchanged, so the provider's defaults apply (shown below as "provider default"); its JSON output, errors and exit code are passed straight through. A missing provider, or one that speaks another contract version, fails with a JSON error that names the package to install or update.

This command is read-only: it keeps working while the stop switch is on.

Platform specifics: `aux4 computer notes` and the provider's own page, e.g. `aux4 aux4 man computer mac find`.

#### Usage

```bash
aux4 computer find [text] [--role <value>] [--app <value>] [--provider <value>]
```

--text      Text to search for in element label/value
--role      Filter by role: the provider's native name or the mac name (AXButton, AXTextField, AXStaticText...)
--app       App to search (defaults to the frontmost app)
--provider  Provider to use: mac, linux or win (default: the one for this operating system; env AUX4_COMPUTER_PROVIDER)

#### Example

```bash
aux4 computer find "Save" --role AXButton --app "TextEdit"
```

```json
{"app":"TextEdit","matches":[{"ref":"0.1.3","role":"AXButton","label":"Save","value":"","frame":{"x":200,"y":600,"width":80,"height":24},"center":{"x":240,"y":612},"enabled":true,"focused":false}]}
```
