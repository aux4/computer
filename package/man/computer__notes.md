#### Description

Prints the current provider's platform notes for agents, in markdown: the permissions it needs, what works in the
background on that platform, and known gotchas. `aux4 ai skill computer prompt` appends these notes to the generic
method. The provider only has to be installed.

#### Usage

```bash
aux4 computer notes [--provider <name>]
```

--provider  Provider to use: mac, linux or win (default: the one for this operating system; env AUX4_COMPUTER_PROVIDER)

#### Example

```bash
aux4 computer notes --provider linux
```

```text
- Needs an X11 session (DISPLAY) and the AT-SPI accessibility bus.
- Keystrokes go to the focused window; with --app they fall back to the foreground.
```
