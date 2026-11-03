# conformance: read-only shapes

## permissions

### should report each permission as granted + guidance

```execute
aux4 computer "$AUX4_COMPUTER_PROVIDER" permissions | jq -c '[.[] | (.granted | type), (.guidance | type)] | unique'
```

```expect
["boolean","string"]
```

### should report every permission as not granted when denied for testing

```execute
AUX4_COMPUTER_DENY_PERMISSIONS=1 aux4 computer "$AUX4_COMPUTER_PROVIDER" permissions | jq -c '[.[] | .granted] | unique'
```

```expect
[false]
```

## screen

### should list displays with index, size, scale and origin

```execute
aux4 computer "$AUX4_COMPUTER_PROVIDER" screen | jq -c '[(.displays | type), (.displays | length > 0), ([.displays[] | (.index | type), (.width | type), (.height | type), (.scale | type), (.x | type), (.y | type), (.main | type)] | unique)]'
```

```expect
["array",true,["boolean","number"]]
```

## cursor

### should report the cursor position

```execute
aux4 computer "$AUX4_COMPUTER_PROVIDER" cursor | jq -c '[(.x | type), (.y | type)]'
```

```expect
["number","number"]
```

## apps

### should list apps with name and pid

```execute
aux4 computer "$AUX4_COMPUTER_PROVIDER" apps | jq -c '[(.apps | type), ([.apps[] | (.name | type), (.pid | type)] | unique)]'
```

```expect:partial
["array",
```

### should give every app a string name and a numeric pid

```execute
aux4 computer "$AUX4_COMPUTER_PROVIDER" apps | jq -r 'all(.apps[]; (.name | type) == "string" and (.pid | type) == "number")'
```

```expect
true
```

## windows

### should list windows with app, title, pid and frame

```execute
aux4 computer "$AUX4_COMPUTER_PROVIDER" windows | jq -r '(.windows | type) == "array" and all(.windows[]; (.app | type) == "string" and (.title | type) == "string" and (.pid | type) == "number" and ([.frame.x, .frame.y, .frame.width, .frame.height] | map(type) | unique) == ["number"])'
```

```expect
true
```
