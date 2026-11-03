#### Description

List an app's menu bar, or press a menu item by path (e.g. 'Format > Font > Bold'); in the background by default.

`aux4 computer menu` runs `aux4 computer <provider> menu` on the provider for this operating system (`mac` on macOS, `linux` on Linux, `win` on Windows), or the one given with `--provider` / `AUX4_COMPUTER_PROVIDER`. Only the flags you give are passed on, unchanged, so the provider's defaults apply (shown below as "provider default"); its JSON output, errors and exit code are passed straight through. A missing provider, or one that speaks another contract version, fails with a JSON error that names the package to install or update.

Without `--path` it only lists the menus (read-only); with `--path` it presses an item, which the stop switch refuses.

`--mode auto` (default) acts in the background and falls back to real foreground input only when the platform can't do it in the background (the result then has `fallbackReason` and `restored`); `--mode background` returns an error instead; `--mode foreground` always uses real input.

Platform specifics: `aux4 computer notes` and the provider's own page, e.g. `aux4 aux4 man computer mac menu`.

#### Usage

```bash
aux4 computer menu [--app <value>] [--path <value>] [--depth <value>] [--mode <value>] [--provider <value>]
```

--app       App whose menu bar to use: name or platform id (default: the frontmost app)
--path      Menu item to press, titles separated by '>' (case-insensitive, trailing '…' optional); omit to list the menu bar
--depth     Menu levels to list: 1 = menu titles, 2 = their items (default), 3+ = submenus (provider default: 2)
--mode      auto (default): act in the background, fall back to real input only if that isn't possible; background: never fall back (error instead); foreground: always use real input (provider default: auto)
--provider  Provider to use: mac, linux or win (default: the one for this operating system; env AUX4_COMPUTER_PROVIDER)

#### Example

```bash
aux4 computer menu --app Calculator --path "View > Scientific"
```

```json
{"app":"Calculator","checked":true,"method":"AXPress","mode":"background","path":["View","Scientific"],"pressed":"Scientific","ref":"1.3.0.2"}
```

```bash
aux4 computer menu --app Calculator --depth 2
```

```json
{"app":"Calculator","depth":2,"menus":[{"ref":"1.3","title":"View","children":[{"checked":true,"ref":"1.3.0.1","shortcut":"cmd+1","title":"Basic"},{"ref":"1.3.0.2","shortcut":"cmd+2","title":"Scientific"},{"ref":"1.3.0.11","submenu":true,"title":"Decimal Places"}]}]}
```
