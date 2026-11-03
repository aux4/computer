#### Description

Resize a window to --width/--height (points; give one to keep the other); returns the frame before and after.

`aux4 computer window resize` runs `aux4 computer <provider> window resize` on the provider for this operating system (`mac` on macOS, `linux` on Linux, `win` on Windows), or the one given with `--provider` / `AUX4_COMPUTER_PROVIDER`. Only the flags you give are passed on, unchanged, so the provider's defaults apply (shown below as "provider default"); its JSON output, errors and exit code are passed straight through. A missing provider, or one that speaks another contract version, fails with a JSON error that names the package to install or update.

It sends input or changes state, so it is refused (exit code `1`, `"stopped": true`) while the stop switch is on (`aux4 computer stop`), and the provider records it in the action log (`aux4 computer log`).

`--mode auto` (default) acts in the background and falls back to real foreground input only when the platform can't do it in the background (the result then has `fallbackReason` and `restored`); `--mode background` returns an error instead; `--mode foreground` always uses real input.

Platform specifics: `aux4 computer notes` and the provider's own page, e.g. `aux4 aux4 man computer mac window resize`.

#### Usage

```bash
aux4 computer window resize [--app <value>] [--title <value>] [--index <value>] [--width <value>] [--height <value>] [--mode <value>] [--provider <value>]
```

--app       App whose window to use: name or platform id (default: the frontmost app)
--title     Window title (exact, then contained; case-insensitive); default: the app's main window
--index     Window index from 'window list' (instead of --title)
--width     New width in points (default: keep)
--height    New height in points (default: keep)
--mode      auto (default): act in the background, fall back to bringing the app forward only if that isn't possible; background: never fall back (error instead); foreground: bring the app forward first (provider default: auto)
--provider  Provider to use: mac, linux or win (default: the one for this operating system; env AUX4_COMPUTER_PROVIDER)

#### Example

```bash
aux4 computer window resize --app TextEdit --title notes --width 50
```

```json
{
  "action": "resize",
  "after": {
    "frame": {
      "height": 422,
      "width": 115,
      "x": 152,
      "y": 73
    },
    "fullscreen": false,
    "index": 0,
    "minimized": false,
    "ref": "0",
    "title": "notes.txt"
  },
  "app": "TextEdit",
  "before": {
    "frame": {
      "height": 422,
      "width": 656,
      "x": 152,
      "y": 73
    },
    "fullscreen": false,
    "index": 0,
    "minimized": false,
    "ref": "0",
    "title": "notes.txt"
  },
  "method": "AXSize",
  "mode": "background",
  "note": "the app kept the window at 115x422 (its minimum/maximum size)",
  "requested": {
    "width": 50
  }
}
```
