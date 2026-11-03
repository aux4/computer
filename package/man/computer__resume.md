#### Description

Turns the shared stop switch off: removes the `STOP` file(s) that `aux4 computer stop` created, so input and
state-changing commands run again on every provider. It succeeds when the switch is already off, and is recorded in
the action log.

An environment switch (`AUX4_COMPUTER_STOP=1`) can't be undone from here: the output then still says
`"stopped": true` with a `stopReason`.

#### Usage

```bash
aux4 computer resume
```

#### Example

```bash
aux4 computer resume
```

```json
{"file":"/Users/sally/Library/Application Support/aux4-computer/STOP","files":["/Users/sally/Library/Application Support/aux4-computer/STOP","/Users/sally/Library/Application Support/aux4-computer-mac/STOP"],"stopped":false}
```
