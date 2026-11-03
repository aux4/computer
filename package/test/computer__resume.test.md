# computer resume

```beforeEach
rm -rf state
```

```afterAll
rm -rf state
```

## turning the switch off

### should remove the STOP file

```execute
export AUX4_COMPUTER_STATE_DIR=$PWD/state; aux4 computer stop > /dev/null; aux4 computer resume | jq -c '{stopped}'; test -f state/STOP || echo removed
```

```expect
{"stopped":false}
removed
```

### should succeed when the switch is already off

```execute
AUX4_COMPUTER_STATE_DIR=$PWD/state aux4 computer resume | jq -c '{stopped}'
```

```expect
{"stopped":false}
```

### should say the environment switch is still on

```execute
AUX4_COMPUTER_STATE_DIR=$PWD/state AUX4_COMPUTER_STOP=1 aux4 computer resume | jq -c '{stopped, stopReason}'
```

```expect
{"stopped":true,"stopReason":"AUX4_COMPUTER_STOP=1 is set; unset it to allow input"}
```
