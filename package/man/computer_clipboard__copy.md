#### Description

Copy to the clipboard: with --text, write it directly; without, copy the app's current selection (background by default) and return it.

`aux4 computer clipboard copy` runs `aux4 computer <provider> clipboard copy` on the provider for this operating system (`mac` on macOS, `linux` on Linux, `win` on Windows), or the one given with `--provider` / `AUX4_COMPUTER_PROVIDER`. Only the flags you give are passed on, unchanged, so the provider's defaults apply (shown below as "provider default"); its JSON output, errors and exit code are passed straight through. A missing provider, or one that speaks another contract version, fails with a JSON error that names the package to install or update.

It sends input or changes state, so it is refused (exit code `1`, `"stopped": true`) while the stop switch is on (`aux4 computer stop`), and the provider records it in the action log (`aux4 computer log`).

`--mode auto` (default) acts in the background and falls back to real foreground input only when the platform can't do it in the background (the result then has `fallbackReason` and `restored`); `--mode background` returns an error instead; `--mode foreground` always uses real input.

Platform specifics: `aux4 computer notes` and the provider's own page, e.g. `aux4 aux4 man computer mac clipboard copy`.

#### Usage

```bash
aux4 computer clipboard copy [--text <value>] [--app <value>] [--mode <value>] [--provider <value>]
```

--text      Text to put on the clipboard; omit to copy the current selection
--app       App whose selection to copy (default: the frontmost app)
--mode      auto (default): act in the background, fall back to real input only if that isn't possible; background: never fall back (error instead); foreground: always use real input (provider default: auto)
--provider  Provider to use: mac, linux or win (default: the one for this operating system; env AUX4_COMPUTER_PROVIDER)

#### Example

```bash
aux4 computer clipboard copy --app TextEdit
```

```json
{"app":"TextEdit","method":"accessibility","mode":"background","text":"the selected text"}
```

```bash
aux4 computer clipboard copy --text "hello"
```

```json
{"text":"hello"}
```
