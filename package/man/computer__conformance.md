#### Description

Runs the provider contract conformance suite against a provider. The suite checks, without sending any real input:

- **info** — `provider`, `contract` (`1`), `platform` (`os/arch`), `version`, boolean `capabilities`, `permissions`;
  accepted by `aux4 computer providers`; `notes` prints markdown
- **shapes** — `permissions` (`granted` + `guidance`), `screen`, `cursor`, `apps`, `windows`
- **errors** — argument validation, JSON error on stderr with exit code `1` and nothing on stdout, missing app,
  denied permissions
- **dry-run decisions** — background versus foreground, the fallback and its reason, `--mode background` refusing
  to fall back, `move` always foreground
- **stop switch** — the core's `stop` stops the provider, read-only commands keep working, `resume`,
  `AUX4_COMPUTER_STOP=1`
- **action log** — the line shape, text truncation, refusals recorded, read-only commands and dry runs not recorded,
  mode `0600`

It copies the suite into `.aux4-computer-conformance/` under `--dir`, runs `aux4 test run` there with
`AUX4_COMPUTER_PROVIDER` set, removes the copy and exits with the test result. Run it from a provider's `package/`
directory to test that working copy without installing it (aux4 loads the `.aux4` of the current directory and its
parents). It needs `aux4/test`.

It refuses a provider whose `info` doesn't report contract `1`: the suite relies on `AUX4_COMPUTER_DRY_RUN` and
`AUX4_COMPUTER_STATE_DIR` to stay headless and away from the real stop switch, and a provider that ignores them would
send real input.

#### Usage

```bash
aux4 computer conformance [--provider <name>] [--dir <path>]
```

--provider  Provider to test (default: the one for this operating system; env AUX4_COMPUTER_PROVIDER)
--dir       Directory to run from (default: `.`)

#### Example

```bash
cd computer-linux/package
aux4 computer conformance --provider linux
```

```text
Running the computer contract 1 conformance suite against provider 'linux' (6 files)
Tests:       32 passed, 32 total
```
