# conformance: action log

Every state-changing action appends one JSON line to `<state dir>/actions.jsonl`, which `aux4 computer log` reads.

```beforeEach
rm -rf state
```

```afterAll
rm -rf state
```

## recording actions

### should record a failed action with the contract fields

```execute
export AUX4_COMPUTER_STATE_DIR=$PWD/state; aux4 computer "$AUX4_COMPUTER_PROVIDER" click --x 10 --y 20 --app ThisAppDoesNotExist12345 2>/dev/null; aux4 computer log --limit 1 | jq -c '.entries[0] | {ts: (.ts | test("^20[0-9][0-9]-[0-9][0-9]-[0-9][0-9]T")), provider: (.provider == env.AUX4_COMPUTER_PROVIDER), command, app, args: (.args | type), x: .args.x, result, error: (.error | type)}'
```

```expect
{"ts":true,"provider":true,"command":"click","app":"ThisAppDoesNotExist12345","args":"object","x":"10","result":"error","error":"string"}
```

### should truncate typed text to 16 characters plus its length

```execute
export AUX4_COMPUTER_STATE_DIR=$PWD/state; aux4 computer "$AUX4_COMPUTER_PROVIDER" type "a long and private sentence typed by an agent" --app ThisAppDoesNotExist12345 2>/dev/null; aux4 computer log --limit 1 | jq -r '.entries[0].args.text'
```

```expect
a long and priva… (45 chars)
```

### should record a refusal by the stop switch

```execute
export AUX4_COMPUTER_STATE_DIR=$PWD/state; AUX4_COMPUTER_STOP=1 aux4 computer "$AUX4_COMPUTER_PROVIDER" click --x 1 --y 1 --app ThisAppDoesNotExist12345 2>/dev/null; aux4 computer log --limit 1 | jq -c '.entries[0] | {command, result}'
```

```expect
{"command":"click","result":"error"}
```

### should not record read-only commands or dry runs

```execute
export AUX4_COMPUTER_STATE_DIR=$PWD/state; aux4 computer "$AUX4_COMPUTER_PROVIDER" cursor > /dev/null; AUX4_COMPUTER_DRY_RUN=1 aux4 computer "$AUX4_COMPUTER_PROVIDER" click --x 1 --y 1 > /dev/null; aux4 computer log | jq '.entries | length'
```

```expect
0
```

### should keep the log private (mode 0600; skipped on Windows, where the file inherits the user-only profile ACL)

```execute
export AUX4_COMPUTER_STATE_DIR=$PWD/state; aux4 computer "$AUX4_COMPUTER_PROVIDER" click --x 10 --y 20 --app ThisAppDoesNotExist12345 2>/dev/null; if [ "$OS" = "Windows_NT" ]; then echo "-rw-------"; else ls -l state/actions.jsonl | cut -c1-10; fi
```

```expect
-rw-------
```
