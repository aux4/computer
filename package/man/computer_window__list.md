#### Description

List an app's windows (index, ref, title, frame, minimized, fullscreen, main, focused).

`aux4 computer window list` runs `aux4 computer <provider> window list` on the provider for this operating system (`mac` on macOS, `linux` on Linux, `win` on Windows), or the one given with `--provider` / `AUX4_COMPUTER_PROVIDER`. Only the flags you give are passed on, unchanged, so the provider's defaults apply (shown below as "provider default"); its JSON output, errors and exit code are passed straight through. A missing provider, or one that speaks another contract version, fails with a JSON error that names the package to install or update.

This command is read-only: it keeps working while the stop switch is on.

Platform specifics: `aux4 computer notes` and the provider's own page, e.g. `aux4 aux4 man computer mac window list`.

#### Usage

```bash
aux4 computer window list [app] [--provider <value>]
```

--app       App whose windows to list: name or platform id (default: the frontmost app)
--provider  Provider to use: mac, linux or win (default: the one for this operating system; env AUX4_COMPUTER_PROVIDER)

#### Example

```bash
aux4 computer window list TextEdit
```

```json
{
  "app": "TextEdit",
  "pid": 86813,
  "windows": [
    {
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
    }
  ]
}
```
