#### Description

Click at coordinates or press an element; in the background by default.

`aux4 computer click` runs `aux4 computer <provider> click` on the provider for this operating system (`mac` on macOS, `linux` on Linux, `win` on Windows), or the one given with `--provider` / `AUX4_COMPUTER_PROVIDER`. Only the flags you give are passed on, unchanged, so the provider's defaults apply (shown below as "provider default"); its JSON output, errors and exit code are passed straight through. A missing provider, or one that speaks another contract version, fails with a JSON error that names the package to install or update.

It sends input or changes state, so it is refused (exit code `1`, `"stopped": true`) while the stop switch is on (`aux4 computer stop`), and the provider records it in the action log (`aux4 computer log`).

`--mode auto` (default) acts in the background and falls back to real foreground input only when the platform can't do it in the background (the result then has `fallbackReason` and `restored`); `--mode background` returns an error instead; `--mode foreground` always uses real input.

Platform specifics: `aux4 computer notes` and the provider's own page, e.g. `aux4 aux4 man computer mac click`.

#### Usage

```bash
aux4 computer click [--x <value>] [--y <value>] [--button <value>] [--count <value>] [--element <value>] [--app <value>] [--mode <value>] [--provider <value>]
```

--x         X coordinate: relative to the app's window with --app (same as 'screenshot --app'), screen points without
--y         Y coordinate: relative to the app's window with --app (same as 'screenshot --app'), screen points without
--button    left or right (provider default: left)
--count     Click count (2 for double-click) (provider default: 1)
--element   Element ref from 'read'/'find' to press instead of x/y
--app       Target app; x/y become relative to its window (default: the frontmost app, screen points)
--mode      auto (default): act in the background, fall back to real input only if that isn't possible; background: never fall back (error instead); foreground: always use real input (provider default: auto)
--provider  Provider to use: mac, linux or win (default: the one for this operating system; env AUX4_COMPUTER_PROVIDER)

#### Example

Press Calculator's `7` button at (34, 211) in its window, while another app is in front:

```bash
aux4 computer click --app Calculator --x 34 --y 211
```

```json
{"action":"AXPress","app":"Calculator","button":"left","count":1,"element":{"label":"7","role":"AXButton"},"mode":"background","screen":{"x":496,"y":610},"x":34,"y":211}
```

When background isn't possible:

```json
{"app":"Calculator","button":"left","count":2,"fallbackReason":"multi-click has no accessibility equivalent","mode":"foreground","restored":{"cursor":true,"frontmost":true},"screen":{"x":496,"y":610},"x":34,"y":211}
```
