#### Description

Prints the current provider's description, as defined by the provider contract: `provider` (its name), `contract`
(the contract version it implements), `platform` (`os/arch`), `version` (the provider package version),
`capabilities` (booleans such as `background`, `ocr`, `annotate`) and `permissions` (the same shape as
`aux4 computer permissions`).

The provider only has to be installed: `info` is how the contract version is discovered, so it is shown even for a
provider this package would refuse. A provider that predates the contract has no `info` command.

#### Usage

```bash
aux4 computer info [--provider <name>]
```

--provider  Provider to use: mac, linux or win (default: the one for this operating system; env AUX4_COMPUTER_PROVIDER)

#### Example

```bash
aux4 computer info
```

```json
{"capabilities":{"annotate":true,"background":true,"ocr":true},"contract":1,"permissions":{"accessibility":{"granted":true,"guidance":"Accessibility access is granted."}},"platform":"darwin/arm64","provider":"mac","version":"0.7.0"}
```
