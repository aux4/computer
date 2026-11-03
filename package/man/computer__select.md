#### Description

Select text in the focused (or given) text element through Accessibility; in the background.

`aux4 computer select` runs `aux4 computer <provider> select` on the provider for this operating system (`mac` on macOS, `linux` on Linux, `win` on Windows), or the one given with `--provider` / `AUX4_COMPUTER_PROVIDER`. Only the flags you give are passed on, unchanged, so the provider's defaults apply (shown below as "provider default"); its JSON output, errors and exit code are passed straight through. A missing provider, or one that speaks another contract version, fails with a JSON error that names the package to install or update.

It sends input or changes state, so it is refused (exit code `1`, `"stopped": true`) while the stop switch is on (`aux4 computer stop`), and the provider records it in the action log (`aux4 computer log`).

`--mode auto` (default) acts in the background and falls back to real foreground input only when the platform can't do it in the background (the result then has `fallbackReason` and `restored`); `--mode background` returns an error instead; `--mode foreground` always uses real input.

Platform specifics: `aux4 computer notes` and the provider's own page, e.g. `aux4 aux4 man computer mac select`.

#### Usage

```bash
aux4 computer select [--app <value>] [--all <value>] [--text <value>] [--occurrence <value>] [--start <value>] [--length <value>] [--element <value>] [--mode <value>] [--provider <value>]
```

--app         App to select in: name or platform id (default: the frontmost app)
--all         Select all the element's text (provider default: false)
--text        Select this text (case-insensitive)
--occurrence  Which match of --text to select (1 = first)
--start       Start of the range to select (character offset from 0)
--length      Length of the range to select (characters)
--element     Text element ref from 'find' or 'state' (default: the app's focused element)
--mode        auto (default): act in the background, fall back to real input only if that isn't possible; background: never fall back (error instead); foreground: always use real input (provider default: auto)
--provider    Provider to use: mac, linux or win (default: the one for this operating system; env AUX4_COMPUTER_PROVIDER)

#### Example

```bash
aux4 computer select --app TextEdit --text aux4
```

```json
{"app":"TextEdit","element":{"label":"","ref":"0.0.0","role":"AXTextArea"},"method":"accessibility","mode":"background","range":{"length":4,"start":11},"selected":"aux4"}
```
