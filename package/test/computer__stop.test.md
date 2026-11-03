# computer stop

```beforeEach
rm -rf state
```

```afterAll
rm -rf state
```

## turning the switch on

### should create the STOP file in the shared state dir

```execute
AUX4_COMPUTER_STATE_DIR=$PWD/state aux4 computer stop | jq -c '{stopped, file: (.file | endswith("/state/STOP")), files: (.files | length)}'; test -f state/STOP && echo exists
```

```expect
{"stopped":true,"file":true,"files":1}
exists
```

### should record stop in the action log

```execute
export AUX4_COMPUTER_STATE_DIR=$PWD/state; aux4 computer stop > /dev/null; aux4 computer log | jq -c '[.entries[] | {provider, command, result}]'
```

```expect
[{"provider":"core","command":"stop","result":"ok"}]
```

### should report the switch in the log output

```execute
export AUX4_COMPUTER_STATE_DIR=$PWD/state; aux4 computer stop > /dev/null; aux4 computer log | jq -r '.stopReason | endswith("/state/STOP exists")'
```

```expect
true
```
