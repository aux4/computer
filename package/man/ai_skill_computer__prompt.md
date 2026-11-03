#### Description

Prints the computer-use method for agents, in markdown, followed by the current provider's platform notes:

- **The loop** — state, find, OCR fallback, screenshot `--annotate`, act by ref (click, menu, set, select, type,
  dialog, window), wait; refs shift when the UI changes
- **Background and foreground** — `--mode auto|background|foreground`, the fallback and what it restores
- **Coordinates** — screen points, window-relative with `--app`, image pixels versus points
- **Safety** — the stop switch (never resume on your own), the action log, the silent data-loss warning for
  background edits (`"documentEdited": false`), not quitting or closing the person's windows
- **Permissions and troubleshooting**
- **Platform notes** — the output of `aux4 computer <provider> notes`. When the provider is missing or predates the
  contract, a short note says what to install or update instead, so the prompt always prints.

Give it to an agent as instructions, e.g. with `aux4 ai agent ask --instructions`.

#### Usage

```bash
aux4 ai skill computer prompt [--provider <name>]
```

--provider  Provider whose notes to append (default: the one for this operating system; env AUX4_COMPUTER_PROVIDER)

#### Example

```bash
aux4 ai skill computer prompt > computer-skill.md
```

```text
# Using a computer with aux4/computer

Drive desktop apps with `aux4 computer <command>`: look (state, find, ocr, screenshot --annotate), act by ref, wait.
...

## Platform notes: mac (aux4/computer-mac)
...
```
