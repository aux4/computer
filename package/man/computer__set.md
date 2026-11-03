#### Description

Set an element: a text field or slider value, a checkbox state or a pop-up option; in the background.

`aux4 computer set` runs `aux4 computer <provider> set` on the provider for this operating system (`mac` on macOS, `linux` on Linux, `win` on Windows), or the one given with `--provider` / `AUX4_COMPUTER_PROVIDER`. Only the flags you give are passed on, unchanged, so the provider's defaults apply (shown below as "provider default"); its JSON output, errors and exit code are passed straight through. A missing provider, or one that speaks another contract version, fails with a JSON error that names the package to install or update.

It sends input or changes state, so it is refused (exit code `1`, `"stopped": true`) while the stop switch is on (`aux4 computer stop`), and the provider records it in the action log (`aux4 computer log`).

`--mode auto` (default) acts in the background and falls back to real foreground input only when the platform can't do it in the background (the result then has `fallbackReason` and `restored`); `--mode background` returns an error instead; `--mode foreground` always uses real input.

Platform specifics: `aux4 computer notes` and the provider's own page, e.g. `aux4 aux4 man computer mac set`.

#### Usage

```bash
aux4 computer set [--app <value>] [--element <value>] [--value <value>] [--checked <value>] [--option <value>] [--mode <value>] [--provider <value>]
```

--app       App the element belongs to: name or platform id (default: the frontmost app)
--element   Element ref from 'find' or 'state'
--value     New value for a text field, text area, combo box or slider
--checked   true or false: press the checkbox/radio button until it is in that state
--option    Option title to pick in a pop-up button, combo box or radio group
--mode      auto (default): act in the background, fall back to real input only if that isn't possible; background: never fall back (error instead); foreground: always use real input (provider default: auto)
--provider  Provider to use: mac, linux or win (default: the one for this operating system; env AUX4_COMPUTER_PROVIDER)

#### Example

```bash
aux4 computer set --app TextEdit --element 0.0.0 --value "Replaced by aux4 computer set"
```

```json
{"after":"Replaced by aux4 computer set","app":"TextEdit","before":"Hello from aux4 computer-mac local check","element":{"label":"","ref":"0.0.0","role":"AXTextArea"},"method":"AXValue","mode":"background"}
```
