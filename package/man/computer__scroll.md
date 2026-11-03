#### Description

Scroll at a position in an app; in the background by default.

`aux4 computer scroll` runs `aux4 computer <provider> scroll` on the provider for this operating system (`mac` on macOS, `linux` on Linux, `win` on Windows), or the one given with `--provider` / `AUX4_COMPUTER_PROVIDER`. Only the flags you give are passed on, unchanged, so the provider's defaults apply (shown below as "provider default"); its JSON output, errors and exit code are passed straight through. A missing provider, or one that speaks another contract version, fails with a JSON error that names the package to install or update.

It sends input or changes state, so it is refused (exit code `1`, `"stopped": true`) while the stop switch is on (`aux4 computer stop`), and the provider records it in the action log (`aux4 computer log`).

`--mode auto` (default) acts in the background and falls back to real foreground input only when the platform can't do it in the background (the result then has `fallbackReason` and `restored`); `--mode background` returns an error instead; `--mode foreground` always uses real input.

Platform specifics: `aux4 computer notes` and the provider's own page, e.g. `aux4 aux4 man computer mac scroll`.

#### Usage

```bash
aux4 computer scroll [--x <value>] [--y <value>] [--dx <value>] [--dy <value>] [--app <value>] [--mode <value>] [--provider <value>]
```

--x         X coordinate to scroll at: relative to the app's window with --app, screen points without
--y         Y coordinate to scroll at: relative to the app's window with --app, screen points without
--dx        Horizontal scroll delta in pixels (provider default: 0)
--dy        Vertical scroll delta in pixels (provider default: 0)
--app       App to scroll (default: the frontmost app)
--mode      auto (default): act in the background, fall back to real input only if that isn't possible; background: never fall back (error instead); foreground: always use real input (provider default: auto)
--provider  Provider to use: mac, linux or win (default: the one for this operating system; env AUX4_COMPUTER_PROVIDER)

#### Example

```bash
aux4 computer scroll --app TextEdit --x 290 --y 250 --dy=-400
```

```json
{"app":"TextEdit","dx":0,"dy":-400,"method":"accessibility","mode":"background","screen":{"x":442,"y":323},"vertical":0.317}
```
