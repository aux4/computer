# computer log

```beforeEach
rm -rf state
```

```afterAll
rm -rf state
```

## reading the shared log

### should be empty when nothing was recorded

```execute
AUX4_COMPUTER_STATE_DIR=$PWD/state aux4 computer log | jq -c '{entries, paths, stopped}'
```

```expect
{"entries":[],"paths":[],"stopped":false}
```

### should merge provider lines and return the most recent, oldest first

```execute
mkdir -p state && printf '%s\n' '{"ts":"2026-11-02T10:00:00.000Z","provider":"linux","command":"click","result":"ok"}' '{"ts":"2026-11-02T10:00:02.000Z","provider":"linux","command":"type","result":"ok"}' 'not json' '{"ts":"2026-11-02T10:00:01.000Z","provider":"linux","command":"scroll","result":"error"}' > state/actions.jsonl && AUX4_COMPUTER_STATE_DIR=$PWD/state aux4 computer log --limit 2 | jq -c '[.entries[].command]'
```

```expect
["scroll","type"]
```

### should read the rotated file too

```execute
mkdir -p state && echo '{"ts":"2026-11-01T09:00:00.000Z","provider":"linux","command":"open","result":"ok"}' > state/actions.jsonl.1 && echo '{"ts":"2026-11-02T09:00:00.000Z","provider":"linux","command":"quit","result":"ok"}' > state/actions.jsonl && AUX4_COMPUTER_STATE_DIR=$PWD/state aux4 computer log | jq -c '[.entries[].command]'
```

```expect
["open","quit"]
```

### should create the log file with mode 0600

```execute
AUX4_COMPUTER_STATE_DIR=$PWD/state aux4 computer stop > /dev/null; ls -l state/actions.jsonl | cut -c1-10
```

```expect
-rw-------
```

### should rotate at the size limit

```execute
mkdir -p state && echo '{"ts":"2026-11-01T09:00:00.000Z","provider":"linux","command":"open","result":"ok"}' > state/actions.jsonl && AUX4_COMPUTER_LOG_MAX_BYTES=10 AUX4_COMPUTER_STATE_DIR=$PWD/state aux4 computer stop > /dev/null; wc -l < state/actions.jsonl | tr -d ' '; jq -r .command state/actions.jsonl.1
```

```expect
1
open
```

## argument validation

### should reject an invalid limit

```execute
AUX4_COMPUTER_STATE_DIR=$PWD/state aux4 computer log --limit 0
```

```error
{"error":"Invalid --limit '0': use a positive number"}
```
