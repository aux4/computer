#### Description

Recognize text on the screen, an app's window or a PNG, with positions in click coordinates.

`aux4 computer ocr` runs `aux4 computer <provider> ocr` on the provider for this operating system (`mac` on macOS, `linux` on Linux, `win` on Windows), or the one given with `--provider` / `AUX4_COMPUTER_PROVIDER`. Only the flags you give are passed on, unchanged, so the provider's defaults apply (shown below as "provider default"); its JSON output, errors and exit code are passed straight through. A missing provider, or one that speaks another contract version, fails with a JSON error that names the package to install or update.

This command is read-only: it keeps working while the stop switch is on.

- **`--format`** -- `json` (default) is the provider's response unchanged: every recognized line in `matches`, each with `confidence`, `frame`, `center` and (with `--text`) `match`. `text` is for reading, not clicking: the core reformats the same matches as plain text, no coordinates -- it groups them into rows by vertical position (so a few pixels of jitter between words on the same line don't split it) and sorts each row left-to-right, printing one line per row, items separated by two spaces. `--format` never reaches the provider; errors and the exit code are unaffected either way. `--text` still filters which lines come back (ignoring case, whitespace and commas) -- with `--format text`, only the matching lines are printed, reformatted the same way.

`--mode auto` (default) acts in the background and falls back to real foreground input only when the platform can't do it in the background (the result then has `fallbackReason` and `restored`); `--mode background` returns an error instead; `--mode foreground` always uses real input.

Platform specifics: `aux4 computer notes` and the provider's own page, e.g. `aux4 aux4 man computer mac ocr`.

#### Usage

```bash
aux4 computer ocr [--app <value>] [--image <value>] [--text <value>] [--level <value>] [--languages <value>] [--mode <value>] [--scale <value>] [--timeout <value>] [--format json|text] [--provider <value>]
```

--app        Read this app's front window (even when covered by other windows); positions are also given relative to it as windowCenter, the space 'click --app' uses
--image      Recognize an existing PNG instead of capturing the screen (no Screen Recording permission needed)
--text       Only return lines containing this text (case-insensitive); each line also gets a 'match' with the position of just that text
--level      accurate (default) or fast recognition (provider default: accurate)
--languages  Comma-separated recognition languages in priority order, e.g. en-US,pt-BR (default: the system's default)
--mode       auto (default): capture in the background, bringing the app to the front only if its window isn't on screen; background: never bring it forward (error instead); foreground: bring the app to the front for the capture, then restore (provider default: auto)
--scale      With --image: image pixels per point (the 'scale' from 'screenshot'), so positions come out in points (provider default: 1)
--timeout    Give up with a JSON error after this long, e.g. 30s, 2m (default 60s: the first recognition after installing or updating takes about 30s) (provider default: 60s)
--format     json (default): the full matches array with positions; text: the recognized text as plain lines in reading order, grouped into rows and sorted left-to-right, no coordinates (handled by this command, same for every provider)
--provider   Provider to use: mac, linux or win (default: the one for this operating system; env AUX4_COMPUTER_PROVIDER)

#### Example

Find Calculator's result even though Calculator is behind other windows, then press a button that OCR
located:

```bash
aux4 computer ocr --app Calculator --text "1,792"
```

```json
{"app":"Calculator","height":408,"level":"accurate","matches":[{"center":{"x":645,"y":511.5},"confidence":1,"frame":{"height":33,"width":74,"x":608,"y":495},"match":{"center":{"x":645,"y":511.5},"frame":{"height":33,"width":74,"x":608,"y":495},"windowCenter":{"x":183,"y":112.5}},"text":"1,792","windowCenter":{"x":183,"y":112.5}}],"mode":"background","scale":2,"source":"window","text":"1,792","width":230,"window":{"frame":{"height":408,"width":230,"x":462,"y":399},"id":331,"title":"Calculator"}}
```

```bash
aux4 computer click --app Calculator --x 34 --y 211.5
```

```json
{"action":"AXPress","app":"Calculator","button":"left","count":1,"element":{"label":"7","role":"AXButton"},"mode":"background","screen":{"x":496,"y":610.5},"x":34,"y":211.5}
```

Recognize a saved PNG:

```bash
aux4 computer ocr --image receipt.png --text total
```

```json
{"height":200,"level":"accurate","matches":[{"center":{"x":404,"y":149},"confidence":1,"frame":{"height":45.4,"width":208.7,"x":299.7,"y":126.3},"match":{"center":{"x":347.3,"y":149.9},"frame":{"height":43.5,"width":95.2,"x":299.7,"y":128.2}},"text":"Total: 1,792"}],"mode":"background","path":"receipt.png","scale":1,"source":"image","text":"total","width":600}
```

See what Calculator currently shows, with no coordinates to wade through:

```bash
aux4 computer ocr --app Calculator --format text
```

```text
42+75×100
7,542
AC  %  7  8  9  ÷
```

Stop after 30 seconds at most:

```bash
aux4 computer ocr --app Calculator --text "1,792" --timeout 30s
```

```json
{"error":"Text recognition did not finish within 30s","note":"macOS may still be preparing its text-recognition model: the first recognition after installing/updating this package or restarting can take about 30s, and that preparation continues in the background, so retry shortly. If every call is slow, macOS's Neural Engine compiler is stuck: 'sudo killall ANECompilerService' or a restart clears it","phase":"recognition","timeout":30,"waited":30,"warmup":true}
```
