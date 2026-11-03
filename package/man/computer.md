#### Description

Computer-use commands for AI agents. `aux4 computer <command>` runs the command on the provider for this operating
system — `aux4/computer-mac` (`mac`), `aux4/computer-linux` (`linux`) or `aux4/computer-windows` (`win`) — so
the same commands, flags and JSON work everywhere. `--provider <name>` picks a provider explicitly.

The core also owns what every provider shares: `providers` (what's installed and compatible), the stop switch
(`stop`, `resume`), the action log (`log`), the conformance suite (`conformance`) and the AI skill
(`aux4 ai skill computer prompt`). Provider commands stay available directly as `aux4 computer <provider> <command>`.

#### Usage

```bash
aux4 computer <command> [--provider <mac|linux|win>] [flags]
```

#### Example

```bash
aux4 computer find "7" --app Calculator
```

```json
{"app":"Calculator","matches":[{"center":{"x":496,"y":610},"frame":{"height":48,"width":56,"x":468,"y":586},"label":"7","ref":"0.0.3.1","role":"AXButton","windowCenter":{"x":34,"y":211}}]}
```
