#### Description

Prints the `computer` [aux4/playbook](https://hub.aux4.io/package/aux4/playbook) action definitions: plain-English
verbs that drive this package's computer-use commands from a playbook script. This command is `private` — it is
hidden from `aux4 playbook actions --help` and from the engine's default global-load discovery — because the
action file declares a top-level `context: computer` key, which makes it **context-bound**: it is never registered
globally, only loaded lazily (and cached for the run) the first time a script enters a matching `on context
computer` block. Calling it directly still works, exactly as `aux4 playbook execute` does internally.

Every action it defines calls the front door `aux4 computer <command>` (never a provider binary directly), so the
shared stop switch and action log always apply.

- `open {app}`, `focus {app}`, `quit {app}` — app lifecycle.
- `state`, `find {text}`, `wait for {text}` — reading the UI.
- `click on {name}`, `click at {x} {y}`, `type {text}`, `shortcut {keys}` / `press {keys}`, `scroll down {n}` /
  `scroll up {n}`, `drag from {x} {y} to {x} {y}` — input.
- `menu {path}`, `dialog`, `press {button} dialog button` — menus and dialogs.
- `screenshot [{file}]`, `ocr [{text}]` — capture.
- `clipboard copy`, `clipboard paste {text}`, `clipboard read` — clipboard.

Every verb above (except `clipboard read`) has two forms: a bare one that targets `{{app}}` — the block's `with app
"<name>"` setting, or the frontmost app / whole screen if the block set none — and one ending in `in {app}` that
overrides the target for just that line.

`click on {name}` has no direct equivalent on `aux4 computer click` (which only takes coordinates or an element
ref), so it resolves the label to coordinates with `aux4 computer find` first — `click --app` takes window-relative
coordinates, so the action picks `find`'s `windowCenter` when an app is set and `center` (screen-absolute) when it
is not.

#### Usage

```bash
aux4 playbook actions computer
```

#### Example

```text
on context computer with app "Calculator"
  open "Calculator"
  wait for "0"
  click on "7"
  click on "Add"
  click on "3"
  click on "Equals"
  wait for "10"
  screenshot "result.png"
end of context
```

```bash
aux4 playbook execute script.txt
```

```text
{"app":"Calculator","bundleId":"com.apple.calculator","mode":"background","opened":"Calculator","pid":12345}
...
{"app":"Calculator","height":408,"path":"result.png","scale":1,"width":230,...}
```
