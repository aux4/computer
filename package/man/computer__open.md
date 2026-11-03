#### Description

Open an app, a file or a URL without bringing anything to the front (by default); optionally wait for its window.

`aux4 computer open` runs `aux4 computer <provider> open` on the provider for this operating system (`mac` on macOS, `linux` on Linux, `win` on Windows), or the one given with `--provider` / `AUX4_COMPUTER_PROVIDER`. Only the flags you give are passed on, unchanged, so the provider's defaults apply (shown below as "provider default"); its JSON output, errors and exit code are passed straight through. A missing provider, or one that speaks another contract version, fails with a JSON error that names the package to install or update.

It sends input or changes state, so it is refused (exit code `1`, `"stopped": true`) while the stop switch is on (`aux4 computer stop`), and the provider records it in the action log (`aux4 computer log`).

`--mode auto` (default) acts in the background and falls back to real foreground input only when the platform can't do it in the background (the result then has `fallbackReason` and `restored`); `--mode background` returns an error instead; `--mode foreground` always uses real input.

Platform specifics: `aux4 computer notes` and the provider's own page, e.g. `aux4 aux4 man computer mac open`.

#### Usage

```bash
aux4 computer open [app] [--mode <value>] [--file <value>] [--url <value>] [--wait <value>] [--hidden <value>] [--provider <value>]
```

--app       Application name or platform id (macOS bundle id, Linux desktop id, Windows process/AppUserModelID); optional with --file/--url (their default app is used)
--mode      auto (default): open without activating, put the frontmost app back if the app activates itself; background: same, error instead of bringing anything forward; foreground: open and activate the app (provider default: auto)
--file      File to open (in --app, or its default app)
--url       URL to open (in --app, or the default handler), e.g. https://aux4.io
--wait      Wait up to this long for the app's window (the file's window with --file) and return its frame, e.g. 10s
--hidden    Launch the app hidden (provider default: false)
--provider  Provider to use: mac, linux or win (default: the one for this operating system; env AUX4_COMPUTER_PROVIDER)

#### Example

```bash
aux4 computer open --file /tmp/notes.txt --app com.apple.TextEdit --wait 10s
```

```json
{"app":"TextEdit","bundleId":"com.apple.TextEdit","file":"/tmp/notes.txt","mode":"background","opened":"com.apple.TextEdit","pid":71767,"waited":0.04,"window":{"frame":{"height":422,"width":656,"x":152,"y":73},"ref":"0","title":"notes.txt"}}
```
