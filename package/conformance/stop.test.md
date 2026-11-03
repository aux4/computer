# conformance: stop switch

The stop switch is shared: the core's `stop` must stop the provider too. AUX4_COMPUTER_STATE_DIR keeps the test away
from the real switch.

```beforeEach
rm -rf state
```

```afterAll
rm -rf state
```

## the shared STOP file

### should refuse input after 'aux4 computer stop'

```execute
export AUX4_COMPUTER_STATE_DIR=$PWD/state; aux4 computer stop > /dev/null; AUX4_COMPUTER_DRY_RUN=1 aux4 computer "$AUX4_COMPUTER_PROVIDER" click --x 1 --y 1 2>&1 >/dev/null | tail -1 | jq -c '{stopped, error: (.error | type)}'
```

```expect
{"stopped":true,"error":"string"}
```

### should exit 1 when refusing

```execute
export AUX4_COMPUTER_STATE_DIR=$PWD/state; aux4 computer stop > /dev/null; AUX4_COMPUTER_DRY_RUN=1 aux4 computer "$AUX4_COMPUTER_PROVIDER" type hello --app ThisAppDoesNotExist12345 2>/dev/null; echo "exit=$?"
```

```expect
exit=1
```

### should keep read-only commands working

```execute
export AUX4_COMPUTER_STATE_DIR=$PWD/state; aux4 computer stop > /dev/null; aux4 computer "$AUX4_COMPUTER_PROVIDER" cursor | jq -c '[(.x | type), (.y | type)]'
```

```expect
["number","number"]
```

### should allow input again after 'aux4 computer resume'

```execute
export AUX4_COMPUTER_STATE_DIR=$PWD/state; aux4 computer stop > /dev/null; aux4 computer resume > /dev/null; AUX4_COMPUTER_DRY_RUN=1 aux4 computer "$AUX4_COMPUTER_PROVIDER" click --x 1 --y 1 | jq -c '{dryRun}'
```

```expect
{"dryRun":true}
```

## the environment switch

### should refuse input when AUX4_COMPUTER_STOP=1

```execute
AUX4_COMPUTER_STATE_DIR=$PWD/state AUX4_COMPUTER_STOP=1 AUX4_COMPUTER_DRY_RUN=1 aux4 computer "$AUX4_COMPUTER_PROVIDER" click --x 1 --y 1 2>&1 >/dev/null | tail -1 | jq -c '{stopped}'
```

```expect
{"stopped":true}
```
