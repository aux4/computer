#### Description

Detect an open alert, sheet, dialog or save/open panel (kind, title, text, buttons, fields); --press a button by title.

`aux4 computer dialog` runs `aux4 computer <provider> dialog` on the provider for this operating system (`mac` on macOS, `linux` on Linux, `win` on Windows), or the one given with `--provider` / `AUX4_COMPUTER_PROVIDER`. Only the flags you give are passed on, unchanged, so the provider's defaults apply (shown below as "provider default"); its JSON output, errors and exit code are passed straight through. A missing provider, or one that speaks another contract version, fails with a JSON error that names the package to install or update.

Without `--press` it only reads the dialog (read-only); with `--press` it answers it, which the stop switch refuses.

`--mode auto` (default) acts in the background and falls back to real foreground input only when the platform can't do it in the background (the result then has `fallbackReason` and `restored`); `--mode background` returns an error instead; `--mode foreground` always uses real input.

Platform specifics: `aux4 computer notes` and the provider's own page, e.g. `aux4 aux4 man computer mac dialog`.

#### Usage

```bash
aux4 computer dialog [--app <value>] [--press <value>] [--mode <value>] [--provider <value>]
```

--app       App to check: name or platform id (default: the frontmost app, then system alerts)
--press     Button to press, by title (case-insensitive, e.g. "Don't Save")
--mode      auto (default): act in the background, fall back to real input only if that isn't possible; background: never fall back (error instead); foreground: always use real input (provider default: auto)
--provider  Provider to use: mac, linux or win (default: the one for this operating system; env AUX4_COMPUTER_PROVIDER)

#### Example

```bash
aux4 computer dialog --app TextEdit
```

```json
{"app":"TextEdit","dialog":{"app":"TextEdit","buttons":[{"cancel":false,"center":{"x":389,"y":490},"default":false,"enabled":true,"focused":false,"frame":{"height":32,"width":86,"x":346,"y":474},"label":"Delete","ref":"0.8.13","role":"AXButton","title":"Delete","value":"","windowCenter":{"x":56,"y":359}},{"cancel":false,"center":{"x":559,"y":490},"default":false,"enabled":true,"focused":false,"frame":{"height":32,"width":86,"x":516,"y":474},"label":"Cancel","ref":"0.8.14","role":"AXButton","title":"Cancel","value":"","windowCenter":{"x":226,"y":359}},{"cancel":false,"center":{"x":639,"y":490},"default":false,"enabled":true,"focused":false,"frame":{"height":32,"width":86,"x":596,"y":474},"label":"Save","ref":"0.8.15","role":"AXButton","title":"Save","value":"","windowCenter":{"x":306,"y":359}}],"fields":[{"center":{"x":540.5,"y":312},"enabled":true,"focused":false,"frame":{"height":22,"width":235,"x":423,"y":301},"label":"","ref":"0.8.5","role":"AXTextField","value":"Untitled","windowCenter":{"x":207.5,"y":181}},{"center":{"x":527,"y":387.5},"enabled":true,"focused":false,"frame":{"height":25,"width":214,"x":420,"y":375},"label":"Where:","ref":"0.8.9","role":"AXPopUpButton","value":"Documents","windowCenter":{"x":194,"y":256.5}}],"frame":{"height":388,"width":362,"x":333,"y":131},"kind":"save-panel","ref":"0.8","text":["Do you want to keep this new document “Untitled”?","You can choose to save your changes, or delete this document immediately. You can’t undo this action.","Save As:","Where:"],"title":"save","window":"Untitled"}}
```

```bash
aux4 computer dialog --app TextEdit --press "Delete"
```

```json
{"app":"TextEdit","closed":true,"dialog":{"kind":"save-panel","title":"save"},"kind":"save-panel","method":"AXPress","mode":"background","pressed":"Delete"}
```
