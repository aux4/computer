# playbook actions computer

Verifies the `computer` playbook context (`package/actions/computer.yaml`): every action calls the front door
`aux4 computer <command>`, `with app` scopes the default target, a per-step `in {app}` wins, and `click on {name}`
resolves through `find`. Dispatches to a fake provider (via `AUX4_COMPUTER_PROVIDER`) so nothing real is driven.

```file:.aux4
{
  "profiles": [
    {
      "name": "main",
      "commands": [
        {
          "name": "computer",
          "execute": [
            "profile:computer"
          ]
        }
      ]
    },
    {
      "name": "computer",
      "commands": [
        {
          "name": "fake",
          "execute": [
            "profile:computer:fake"
          ],
          "help": {
            "text": "Fake provider that echoes what it receives"
          }
        }
      ]
    },
    {
      "name": "computer:fake",
      "commands": [
        {
          "name": "info",
          "execute": [
            "echo '{\"provider\":\"fake\",\"contract\":1,\"platform\":\"test/none\",\"version\":\"0.0.0\",\"capabilities\":{\"background\":true}}'"
          ]
        },
        {
          "name": "open",
          "execute": [
            "printf 'open app=%s\\n' values(app)"
          ],
          "help": {
            "text": "Echo an open",
            "variables": [
              { "name": "app", "text": "app", "default": "" }
            ]
          }
        },
        {
          "name": "wait",
          "execute": [
            "printf 'wait text=%s app=%s\\n' values(text, app)"
          ],
          "help": {
            "text": "Echo a wait",
            "variables": [
              { "name": "text", "text": "text", "default": "" },
              { "name": "app", "text": "app", "default": "" }
            ]
          }
        },
        {
          "name": "find",
          "execute": [
            "if [ \"$AUX4_COMPUTER_TEST_EMPTY\" = \"1\" ]; then echo '{\"matches\":[]}'; else printf '{\"matches\":[{\"label\":\"%s\",\"center\":{\"x\":500,\"y\":600},\"windowCenter\":{\"x\":34,\"y\":211},\"frame\":{\"width\":48,\"height\":48}}]}\\n' values(text); fi"
          ],
          "help": {
            "text": "Echo one matching element, or none when AUX4_COMPUTER_TEST_EMPTY=1",
            "variables": [
              { "name": "text", "text": "text", "default": "" },
              { "name": "app", "text": "app", "default": "" }
            ]
          }
        },
        {
          "name": "click",
          "execute": [
            "printf 'click x=%s y=%s app=%s\\n' values(x, y, app)"
          ],
          "help": {
            "text": "Echo a click",
            "variables": [
              { "name": "x", "text": "x", "default": "" },
              { "name": "y", "text": "y", "default": "" },
              { "name": "app", "text": "app", "default": "" }
            ]
          }
        }
      ]
    }
  ]
}
```

## with app

### open and click target the block's app, using find's window-relative coordinates

```file:with-app.txt
on context computer with app "Calculator"
  open "Calculator"
  wait for "0"
  click on "7"
end of context
```

```execute
AUX4_COMPUTER_PROVIDER=fake aux4 playbook execute with-app.txt
```

```expect
open app=Calculator
wait text=0 app=Calculator
click x=34 y=211 app=Calculator
```

## without app

### no "with app" clause leaves app empty and uses find's screen-absolute coordinates

```file:no-app.txt
on context computer
  click on "7"
end of context
```

```execute
AUX4_COMPUTER_PROVIDER=fake aux4 playbook execute no-app.txt
```

```expect
click x=500 y=600 app=
```

## per-step override

### "in {app}" on a single step wins over the block's app for that line only

```file:override.txt
on context computer with app "Calculator"
  click on "7"
  click on "OK" in "Finder"
  click on "8"
end of context
```

```execute
AUX4_COMPUTER_PROVIDER=fake aux4 playbook execute override.txt
```

```expect
click x=34 y=211 app=Calculator
click x=34 y=211 app=Finder
click x=34 y=211 app=Calculator
```

## no element found

### "click on" fails clearly when find returns no match

```file:nomatch.txt
on context computer with app "Calculator"
  click on "Nope"
end of context
```

```execute
AUX4_COMPUTER_PROVIDER=fake AUX4_COMPUTER_TEST_EMPTY=1 aux4 playbook execute nomatch.txt; echo "exit=$?"
```

```expect:partial
exit=1
```

```error:partial
No element matching "Nope" found
```
