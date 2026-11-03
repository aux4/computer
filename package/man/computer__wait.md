#### Description

Wait until text, an element or an app window appears (or is gone), checking the accessibility tree first, then OCR.

`aux4 computer wait` runs `aux4 computer <provider> wait` on the provider for this operating system (`mac` on macOS, `linux` on Linux, `win` on Windows), or the one given with `--provider` / `AUX4_COMPUTER_PROVIDER`. Only the flags you give are passed on, unchanged, so the provider's defaults apply (shown below as "provider default"); its JSON output, errors and exit code are passed straight through. A missing provider, or one that speaks another contract version, fails with a JSON error that names the package to install or update.

This command is read-only: it keeps working while the stop switch is on.

Platform specifics: `aux4 computer notes` and the provider's own page, e.g. `aux4 aux4 man computer mac wait`.

#### Usage

```bash
aux4 computer wait [text] [--role <value>] [--app <value>] [--gone <value>] [--source <value>] [--timeout <value>] [--interval <value>] [--provider <value>]
```

--text      Text to wait for in element labels/values or on the pixels (ignores case, whitespace and commas, so 1792 matches 1,792)
--role      Only elements with this role: the provider's native name or the mac name (AXButton, AXTextField...), accessibility only
--app       App to watch (defaults to the frontmost app / the whole screen); alone, wait until it is running with an on-screen window
--gone      true: wait until the text, element or app window is no longer there (provider default: false)
--source    auto (default): accessibility first, OCR of the app window (or screen) when accessibility has no match; ax (accessibility) or ocr: only that one (provider default: auto)
--timeout   How long to wait before failing, e.g. 500ms, 10s, 2m (provider default: 10s)
--interval  Pause between checks, e.g. 250ms, 1s (checks never overlap) (provider default: 500ms)
--provider  Provider to use: mac, linux or win (default: the one for this operating system; env AUX4_COMPUTER_PROVIDER)

#### Example

Wait for Calculator to show the result after pressing `=`:

```bash
aux4 computer wait --app Calculator --text 1792
```

```json
{"app":"Calculator","matches":[{"center":{"x":645,"y":511},"enabled":true,"focused":false,"frame":{"height":36,"width":206,"x":542,"y":493},"label":"Edit field","ref":"0.0.0.0.1.0.0.0.0.0","role":"AXStaticText","value":"1,792","windowCenter":{"x":183,"y":112}}],"source":"ax","waited":1.7}
```

Wait for an app to open a window:

```bash
aux4 computer wait --app TextEdit --timeout 30s
```

```json
{"app":"TextEdit","matches":[{"app":"TextEdit","frame":{"height":488,"width":586,"x":152,"y":73},"id":258,"pid":15584,"title":"Untitled.rtf"}],"source":"app","waited":0.01}
```

Wait for a spinner's text to go away:

```bash
aux4 computer wait --app Safari --text "Loading" --gone true --timeout 1m
```

When it times out:

```bash
aux4 computer wait --app Calculator --text 1793 --timeout 2s
```

```json
{"condition":{"app":"Calculator","gone":false,"source":"auto","text":"1793"},"error":"Timed out after 2s waiting for text '1793' in Calculator to appear","seen":{"ax":["Calculator","256 × 7","1,792","Delete","All Clear","Percent","Divide","7","8","9","Multiply"],"ocr":["256×7","1,792","AC","%","7","8","9"],"running":true,"window":true},"timeout":2,"waited":2.1}
```
