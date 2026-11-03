#### Description

Lists the computer providers installed on this machine and whether this package can use them. For each known
provider (`mac`, `linux`, `win`, plus the one given with `--provider`) it calls `aux4 computer <provider> info`:

- **`ok`** — the provider reports contract `1`; its `info` (platform, version, capabilities, permissions) is included.
- **`legacy`** — the provider is installed but has no `info` command (it predates the contract). Commands are still
  dispatched to it; update the provider package.
- **`mismatch`** — its `info` reports another contract version (or isn't JSON); commands are refused until both
  packages are updated.
- **`error`** — `info` failed; the error is included.

`default` is the provider this OS uses (`""` on an OS without one) and each entry says whether it is the default.
Providers that aren't installed are listed in `notInstalled` with their package; when the provider for this OS is
missing, `hint` says how to install it. `contract` is the contract version this package implements.

#### Usage

```bash
aux4 computer providers [--provider <name>]
```

--provider  Also probe this provider name (beyond mac, linux and win)

#### Example

```bash
aux4 computer providers
```

```json
{"contract":1,"default":"mac","notInstalled":[{"package":"aux4/computer-linux","provider":"linux"},{"package":"aux4/computer-windows","provider":"win"}],"platform":"darwin/arm64","providers":[{"contract":1,"default":true,"info":{"capabilities":{"annotate":true,"background":true,"ocr":true},"contract":1,"platform":"darwin/arm64","provider":"mac","version":"0.7.0"},"package":"aux4/computer-mac","provider":"mac","status":"ok"}]}
```
