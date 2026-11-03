# conformance: info and notes

Run with the provider name in AUX4_COMPUTER_PROVIDER (`aux4 computer conformance --provider <name>` does it).

## info

### should describe the provider with the contract fields

```execute
aux4 computer "$AUX4_COMPUTER_PROVIDER" info | jq -c '{provider: (.provider == env.AUX4_COMPUTER_PROVIDER), contract, platform: (.platform | type), version: (.version | type), capabilities: (.capabilities | type), permissions: (.permissions | type)}'
```

```expect
{"provider":true,"contract":1,"platform":"string","version":"string","capabilities":"object","permissions":"object"}
```

### should report a platform as os/arch

```execute
aux4 computer "$AUX4_COMPUTER_PROVIDER" info | jq -r '.platform | test("^(darwin|linux|windows)/[a-z0-9]+$")'
```

```expect
true
```

### should report capabilities as booleans

```execute
aux4 computer "$AUX4_COMPUTER_PROVIDER" info | jq -c '[.capabilities[] | type] | unique'
```

```expect
["boolean"]
```

### should be accepted by the core

```execute
aux4 computer providers --provider "$AUX4_COMPUTER_PROVIDER" | jq -r '.providers[] | select(.provider == env.AUX4_COMPUTER_PROVIDER) | .status'
```

```expect
ok
```

## notes

### should print markdown for agents

```execute
aux4 computer "$AUX4_COMPUTER_PROVIDER" notes | grep -c . | awk '{print ($1 > 0) ? "has notes" : "empty"}'
```

```expect
has notes
```
