# Using a computer with aux4/computer

Drive desktop apps with `aux4 computer <command>`: look (state, find, ocr, screenshot --annotate), act by ref, wait.
The same commands work on macOS, Linux and Windows: aux4/computer sends each one to the provider for this operating
system (`aux4 computer providers` shows which one and what it can do). `--provider <name>` picks one explicitly.

Every command prints one JSON object on stdout (errors: JSON on stderr, exit code 1). Assume a person is using this
computer: actions go to the app in the background where the platform allows it, without moving their cursor or
bringing the app forward. Check `mode` in each result.

## The loop

1. **State** -- `aux4 computer state --app <App>`: app, windows, focused element, open dialog and a compact tree of
   controls, each with a `ref`. Start here, and re-read after anything that changes the UI.
2. **Find** -- `aux4 computer find "<text>" --app <App> [--role AXButton]`: elements whose label or value contains
   the text (case-insensitive), with `ref`, `frame`, `center` and `windowCenter`. `--role` takes the platform's own
   role names or the mac names (AXButton, AXTextField, AXStaticText, AXCheckBox, AXMenuItem, AXWindow).
3. **OCR fallback** -- when the text isn't in the accessibility tree (canvas, web content, images):
   `aux4 computer ocr --app <App> --text "<text>"` reads the window's pixels; use the match's `windowCenter`.
   To just see what an app currently shows (no coordinates, nothing to click), use
   `aux4 computer ocr --app <App> --format text`: plain text lines in reading order, one line per row on
   screen, instead of the full JSON with a position for every piece of text.
4. **Look** -- when you need to see the layout: `aux4 computer screenshot --app <App> --annotate true` draws a
   numbered box on each actionable element and returns `marks: [{n, ref, role, label, frame, center, windowCenter,
   image}]`. Look at the image, pick a number, act on that mark's `ref`. `--roles all` also marks text and images.
5. **Act** -- prefer, in order:
   - `click --element <ref> --app <App>` (presses the element; no coordinates)
   - `menu --app <App> --path "File > Save As…"` (app menu items; `menu --app <App>` alone lists them)
   - `set --app <App> --element <ref> --value "…"` / `--checked true` / `--option "…"` (fields, checkboxes, pop-ups)
   - `select --app <App> --all true` / `--text "…"`, then `clipboard copy --app <App>`
   - `type "…" --app <App>` (inserts at the caret), `dialog --app <App> --press "Don't Save"`
   - `window move|resize|minimize|restore|close --app <App>`, `hide`/`unhide <App>`, `quit <App>`
   - last resort: `click --x --y --app <App>` (window coordinates), `shortcut <keys> --app <App>`
6. **Wait** -- don't sleep and hope: `aux4 computer wait --app <App> --text "Saved" --timeout 10s` (or `--gone true`
   for something to disappear, `--role` for a sheet or dialog). Then go back to 1.

Refs are paths into the accessibility tree; they shift when the UI changes (a new window, a sheet, a dialog). Re-read
before acting on an old ref. Text matching in `wait` and `ocr --text` ignores case, whitespace and commas.

## Background and foreground

Every action takes `--mode auto|background|foreground`:

- `auto` (default): act in the background; when that isn't possible for this action, **fall back to the
  foreground** -- the app is brought forward for a moment, real input is sent, then the cursor and the frontmost app
  are put back. The result then says `mode: "foreground"`, a `fallbackReason` and `restored`.
- `background`: never fall back; get an error instead. Use it when the person must not be disturbed.
- `foreground`: always use real input.

What works in the background differs by platform -- read the platform notes below. Typically: pressing buttons and
menu items, setting values, selecting and inserting text, scrolling, reading and screenshots of covered windows work
in the background; keystrokes to an app that isn't frontmost, drag and multi-clicks usually need the foreground.
`focus` and `window raise` are always foreground. `move` and `drag` move the real cursor.

## Coordinates

- Screen points (logical pixels, independent of display scaling), origin top-left of the main display. With `--app`,
  `--x/--y` are relative to the app's front window -- the same space as `screenshot --app` and every `windowCenter`.
- Screenshot images may be scaled: `image` boxes in `marks` are pixels of the PNG; `center`/`windowCenter` are
  points. Never pass image pixels to `click`.
- Negative values need `=`: `--x=-200`.

## Safety

- `aux4 computer stop` turns on the stop switch: every input or state-changing command, on every provider, is refused
  (`"stopped": true`) until `aux4 computer resume`. Read-only commands keep working, so you can still look. Respect
  it; **never resume on your own** -- the person turned it on.
- `aux4 computer log` lists what was done (command, app, mode, result); typed text is recorded only as a short prefix
  and its length.
- **Silent data loss:** text written through accessibility (`set --value`, `type` or `clipboard paste` in the
  background) may not mark a document as edited, so closing or quitting can discard it **without asking**. When this
  applies the result includes `"documentEdited": false` and a `"warning"` -- save explicitly (the app's Save menu item
  or shortcut) before closing, or use `--mode foreground` for real keystrokes. Don't assume a background edit is safe
  to walk away from.
- Don't quit, close or move the person's windows unless asked. `quit` reports a "save changes?" dialog instead of
  quitting; answer it with `dialog --press` only as the person wants. `quit --force true` loses unsaved work.

## Permissions and troubleshooting

- `aux4 computer permissions` shows what the platform needs (accessibility, screen capture) and whether it's granted.
- `Element not found for ref` -- the UI changed: run `state` or `find` again.
- An action did nothing -- check the result's `mode`/`method`, then verify with `state`/`wait`; never repeat a press
  blindly (a toggle flips back).
- Nothing visible to annotate -- the app has no on-screen window (hidden or minimized): `unhide` or `window restore`.
- `No computer provider ... is installed` -- install the package the error names.
