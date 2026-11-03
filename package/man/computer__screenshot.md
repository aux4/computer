#### Description

Capture a screenshot of a display or a specific app's window.

`aux4 computer screenshot` runs `aux4 computer <provider> screenshot` on the provider for this operating system (`mac` on macOS, `linux` on Linux, `win` on Windows), or the one given with `--provider` / `AUX4_COMPUTER_PROVIDER`. Only the flags you give are passed on, unchanged, so the provider's defaults apply (shown below as "provider default"); its JSON output, errors and exit code are passed straight through. A missing provider, or one that speaks another contract version, fails with a JSON error that names the package to install or update.

This command is read-only: it keeps working while the stop switch is on.

Platform specifics: `aux4 computer notes` and the provider's own page, e.g. `aux4 aux4 man computer mac screenshot`.

#### Usage

```bash
aux4 computer screenshot [--output <value>] [--display <value>] [--app <value>] [--maxWidth <value>] [--annotate <value>] [--roles <value>] [--limit <value>] [--provider <value>]
```

--output    File path to save the PNG to (defaults to a temp file)
--display   Display index to capture (from 'computer screen') (provider default: 0)
--app       Capture this app's front window (even when covered by other windows) instead of a display
--maxWidth  Resize down if wider than this, preserving aspect ratio (provider default: 1568)
--annotate  Draw a numbered box on each element (set-of-marks) and return the marks: number, ref, role, label, frame, center, windowCenter (provider default: false)
--roles     With --annotate: actionable (default: buttons, links, fields, checkboxes, menus, pop-ups, tabs, rows) or all (also text, images, headings, cells) (provider default: actionable)
--limit     With --annotate: most marks to draw (the rest are reported as truncated) (provider default: 150)
--provider  Provider to use: mac, linux or win (default: the one for this operating system; env AUX4_COMPUTER_PROVIDER)

#### Example

```bash
aux4 computer screenshot --app "Calculator" --output /tmp/calculator.png
```

```json
{"app":"Calculator","height":408,"path":"/tmp/calculator.png","scale":1,"width":230,"window":{"frame":{"height":408,"width":230,"x":462,"y":399},"id":295,"title":"Calculator"}}
```

With `--annotate`, the image has a numbered box on every control and the output lists them:

```bash
aux4 computer screenshot --app Calculator --annotate true --output /tmp/calculator.png
```

```json
{
  "annotated": true,
  "app": "Calculator",
  "height": 408,
  "marks": [
    {
      "center": {"x": 496, "y": 610},
      "frame": {"height": 48, "width": 48, "x": 472, "y": 586},
      "image": {"height": 48, "width": 48, "x": 10, "y": 187},
      "label": "7",
      "n": 10,
      "ref": "0.0.0.1.0.6",
      "role": "AXButton",
      "windowCenter": {"x": 34, "y": 211}
    }
  ],
  "path": "/tmp/calculator.png",
  "roles": "actionable",
  "scale": 1,
  "width": 230,
  "window": {"frame": {"height": 408, "width": 230, "x": 462, "y": 399}, "id": 295, "title": "Calculator"}
}
```

(One of the 25 marks shown.) Press mark 10 with `aux4 computer click --element 0.0.0.1.0.6 --app Calculator`.
