# conformance: errors

Errors are one JSON object on stderr with an `error` string, and exit code 1. Validation runs under
AUX4_COMPUTER_DRY_RUN=1 so a provider that validated too late still sends nothing.

## argument validation

### should reject an invalid --mode

```execute
AUX4_COMPUTER_DRY_RUN=1 aux4 computer "$AUX4_COMPUTER_PROVIDER" click --x 1 --y 1 --mode bogus 2>&1 >/dev/null | tail -1 | jq -c '{error: (.error | test("--mode"))}'; AUX4_COMPUTER_DRY_RUN=1 aux4 computer "$AUX4_COMPUTER_PROVIDER" click --x 1 --y 1 --mode bogus 2>/dev/null; echo "exit=$?"
```

```expect
{"error":true}
exit=1
```

### should reject a click without a target

```execute
AUX4_COMPUTER_DRY_RUN=1 aux4 computer "$AUX4_COMPUTER_PROVIDER" click 2>&1 >/dev/null | tail -1 | jq -c '{error: (.error | type)}'; AUX4_COMPUTER_DRY_RUN=1 aux4 computer "$AUX4_COMPUTER_PROVIDER" click 2>/dev/null; echo "exit=$?"
```

```expect
{"error":"string"}
exit=1
```

### should print nothing on stdout when failing

```execute
AUX4_COMPUTER_DRY_RUN=1 aux4 computer "$AUX4_COMPUTER_PROVIDER" click --x 1 --y 1 --mode bogus 2>/dev/null | wc -c | tr -d ' '
```

```expect
0
```

## missing app

### should fail with a JSON error for an app that isn't running

```execute
aux4 computer "$AUX4_COMPUTER_PROVIDER" find Save --app ThisAppDoesNotExist12345 2>&1 >/dev/null | tail -1 | jq -c '{error: (.error | type)}'
```

```expect
{"error":"string"}
```

## permissions denied

### should refuse input with a JSON error when permissions are denied

```execute
AUX4_COMPUTER_DENY_PERMISSIONS=1 AUX4_COMPUTER_DRY_RUN=1 aux4 computer "$AUX4_COMPUTER_PROVIDER" click --x 1 --y 1 2>&1 >/dev/null | tail -1 | jq -c '{error: (.error | type)}'; AUX4_COMPUTER_DENY_PERMISSIONS=1 AUX4_COMPUTER_DRY_RUN=1 aux4 computer "$AUX4_COMPUTER_PROVIDER" click --x 1 --y 1 2>/dev/null; echo "exit=$?"
```

```expect
{"error":"string"}
exit=1
```
