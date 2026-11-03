#### Description

One-call snapshot: frontmost app, the app's windows, focused element, open dialog and a compact tree of its actionable/text elements.

`aux4 computer state` runs `aux4 computer <provider> state` on the provider for this operating system (`mac` on macOS, `linux` on Linux, `win` on Windows), or the one given with `--provider` / `AUX4_COMPUTER_PROVIDER`. Only the flags you give are passed on, unchanged, so the provider's defaults apply (shown below as "provider default"); its JSON output, errors and exit code are passed straight through. A missing provider, or one that speaks another contract version, fails with a JSON error that names the package to install or update.

This command is read-only: it keeps working while the stop switch is on.

Platform specifics: `aux4 computer notes` and the provider's own page, e.g. `aux4 aux4 man computer mac state`.

#### Usage

```bash
aux4 computer state [--app <value>] [--depth <value>] [--limit <value>] [--provider <value>]
```

--app       App to describe: name or platform id (default: the frontmost app)
--depth     Levels of actionable/text elements to include (layout groups don't count; 0 = none) (provider default: 3)
--limit     Maximum number of elements in the compact tree (provider default: 120)
--provider  Provider to use: mac, linux or win (default: the one for this operating system; env AUX4_COMPUTER_PROVIDER)

#### Example

```bash
aux4 computer state --app Calculator
```

```json
{"app":{"bundleId":"com.apple.calculator","frontmost":false,"hidden":false,"name":"Calculator","pid":62830},"depth":3,"dialog":null,"elements":[{"frame":{"height":36,"width":73,"x":609,"y":488},"label":"Edit field","ref":"0.0.0.1.0.1.0","role":"AXStaticText","value":"1,792"},{"frame":{"height":48,"width":48,"x":526,"y":532},"label":"All Clear","ref":"0.0.0.1.0.3","role":"AXButton"}],"focused":{"center":{"x":645.5,"y":506},"enabled":true,"focused":false,"frame":{"height":36,"width":73,"x":609,"y":488},"label":"Edit field","role":"AXUnknown","value":"1,792","windowCenter":{"x":183.5,"y":107}},"frontmost":{"bundleId":"com.anthropic.claudefordesktop","name":"Claude","pid":49456},"root":"0","windows":[{"frame":{"height":408,"width":230,"x":462,"y":399},"main":true,"ref":"0","title":"Calculator"}]}
```
