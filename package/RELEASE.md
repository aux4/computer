# Release notes

- New package: `aux4 computer <command>` runs the computer-use command surface on the provider for this OS (`mac`, `linux`, `win`), with `--provider` to pick one; flags pass through unchanged.
- Clear errors when the provider is missing (names the package to install) or speaks another contract version; providers without `info` are used and listed as `legacy`.
- `aux4 computer providers`, `info` and `notes`.
- Shared stop switch and action log across every provider: `aux4 computer stop`, `resume`, `log`; the switch is checked before any state-changing command is dispatched.
- Provider contract v1 conformance suite and `aux4 computer conformance` to run it from a provider repository.
- AI skill: `aux4 ai skill computer prompt` (the generic method plus the current provider's notes).
