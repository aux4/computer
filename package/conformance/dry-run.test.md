# conformance: dry-run decisions

AUX4_COMPUTER_DRY_RUN=1 decides how an action would run (background or foreground) without sending anything.
AUX4_COMPUTER_SIMULATE_BACKGROUND=unavailable makes the background path unavailable.

## click

### should report a dry run with a mode

```execute
AUX4_COMPUTER_DRY_RUN=1 aux4 computer "$AUX4_COMPUTER_PROVIDER" click --x 100 --y 200 | jq -c '{dryRun, x, y, mode: (.mode | IN("background", "foreground"))}'
```

```expect
{"dryRun":true,"x":100,"y":200,"mode":true}
```

### should use real input with --mode foreground, without a fallback reason

```execute
AUX4_COMPUTER_DRY_RUN=1 aux4 computer "$AUX4_COMPUTER_PROVIDER" click --x 100 --y 200 --mode foreground | jq -c '{dryRun, mode, fallbackReason}'
```

```expect
{"dryRun":true,"mode":"foreground","fallbackReason":null}
```

### should fall back to the foreground with a reason when the background is unavailable

```execute
AUX4_COMPUTER_DRY_RUN=1 AUX4_COMPUTER_SIMULATE_BACKGROUND=unavailable aux4 computer "$AUX4_COMPUTER_PROVIDER" click --x 100 --y 200 | jq -c '{dryRun, mode, fallbackReason: (.fallbackReason | type)}'
```

```expect
{"dryRun":true,"mode":"foreground","fallbackReason":"string"}
```

### should refuse to fall back with --mode background

```execute
AUX4_COMPUTER_DRY_RUN=1 AUX4_COMPUTER_SIMULATE_BACKGROUND=unavailable aux4 computer "$AUX4_COMPUTER_PROVIDER" click --x 100 --y 200 --mode background 2>&1 >/dev/null | tail -1 | jq -c '{error: (.error | type)}'; AUX4_COMPUTER_DRY_RUN=1 AUX4_COMPUTER_SIMULATE_BACKGROUND=unavailable aux4 computer "$AUX4_COMPUTER_PROVIDER" click --x 100 --y 200 --mode background 2>/dev/null; echo "exit=$?"
```

```expect
{"error":"string"}
exit=1
```

## move

### should always use the foreground (it moves the real cursor)

```execute
AUX4_COMPUTER_DRY_RUN=1 aux4 computer "$AUX4_COMPUTER_PROVIDER" move --x 10 --y 20 | jq -r .mode
```

```expect
foreground
```
