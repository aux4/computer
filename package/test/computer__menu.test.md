# computer menu

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
            "text": "Fake provider"
          }
        },
        {
          "name": "legacy",
          "execute": [
            "profile:computer:legacy"
          ],
          "help": {
            "text": "Fake provider without info"
          }
        },
        {
          "name": "future",
          "execute": [
            "profile:computer:future"
          ],
          "help": {
            "text": "Fake provider speaking contract 2"
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
          "name": "notes",
          "execute": [
            "echo '- Fake notes: nothing here is real.'"
          ]
        },
        {
          "name": "screenshot",
          "execute": [
            "printf 'fake screenshot app=%s maxWidth=%s\\n' values(app, maxWidth)"
          ],
          "help": {
            "text": "Echo",
            "variables": [
              {
                "name": "app",
                "text": "app",
                "default": ""
              },
              {
                "name": "maxWidth",
                "text": "max",
                "default": "1568"
              }
            ]
          }
        },
        {
          "name": "menu",
          "execute": [
            "printf 'fake menu app=%s path=%s\\n' values(app, path)"
          ],
          "help": {
            "text": "Echo",
            "variables": [
              {
                "name": "app",
                "text": "app",
                "default": ""
              },
              {
                "name": "path",
                "text": "path",
                "default": ""
              }
            ]
          }
        },
        {
          "name": "window",
          "execute": [
            "profile:computer:fake:window"
          ],
          "help": {
            "text": "Windows"
          }
        },
        {
          "name": "clipboard",
          "execute": [
            "profile:computer:fake:clipboard"
          ],
          "help": {
            "text": "Clipboard"
          }
        }
      ]
    },
    {
      "name": "computer:fake:window",
      "commands": [
        {
          "name": "list",
          "execute": [
            "printf 'fake window list app=%s\\n' values(app)"
          ],
          "help": {
            "text": "Echo",
            "variables": [
              {
                "name": "app",
                "text": "app",
                "default": "",
                "arg": true
              }
            ]
          }
        },
        {
          "name": "move",
          "execute": [
            "printf 'fake window move app=%s title=%s x=%s y=%s mode=%s\\n' values(app, title, x, y, mode)"
          ],
          "help": {
            "text": "Echo",
            "variables": [
              {
                "name": "app",
                "text": "app",
                "default": ""
              },
              {
                "name": "title",
                "text": "t",
                "default": ""
              },
              {
                "name": "x",
                "text": "x",
                "default": ""
              },
              {
                "name": "y",
                "text": "y",
                "default": ""
              },
              {
                "name": "mode",
                "text": "m",
                "default": "auto"
              }
            ]
          }
        },
        {
          "name": "fullscreen",
          "execute": [
            "printf 'fake window fullscreen app=%s enabled=%s\\n' values(app, enabled)"
          ],
          "help": {
            "text": "Echo",
            "variables": [
              {
                "name": "app",
                "text": "app",
                "default": ""
              },
              {
                "name": "enabled",
                "text": "e",
                "default": "true",
                "arg": true
              }
            ]
          }
        }
      ]
    },
    {
      "name": "computer:fake:clipboard",
      "commands": [
        {
          "name": "read",
          "execute": [
            "echo '{\"text\":\"fake clipboard\"}'"
          ]
        },
        {
          "name": "paste",
          "execute": [
            "printf 'fake clipboard paste text=%s\\n' values(text)"
          ],
          "help": {
            "text": "Echo",
            "variables": [
              {
                "name": "text",
                "text": "t",
                "default": ""
              }
            ]
          }
        }
      ]
    },
    {
      "name": "computer:legacy",
      "commands": [
        {
          "name": "screen",
          "execute": [
            "echo '{\"displays\":[]}'"
          ]
        }
      ]
    },
    {
      "name": "computer:future",
      "commands": [
        {
          "name": "info",
          "execute": [
            "echo '{\"provider\":\"future\",\"contract\":2}'"
          ]
        },
        {
          "name": "notes",
          "execute": [
            "echo '- future notes'"
          ]
        }
      ]
    }
  ]
}
```

```beforeEach
rm -rf state
```

```afterAll
rm -rf state
```

## stop switch

### should list menus while stopped (no --path is read-only)

```execute
export AUX4_COMPUTER_STATE_DIR=$PWD/state; aux4 computer stop > /dev/null; aux4 computer menu --provider fake --app Calculator
```

```expect
fake menu app=Calculator path=
```

### should refuse to press a menu item while stopped

```execute
export AUX4_COMPUTER_STATE_DIR=$PWD/state; aux4 computer stop > /dev/null; aux4 computer menu --provider fake --app Calculator --path "View > Scientific"; echo "exit=$?"
```

```expect
exit=1
```

```error:partial
{"command":"menu","error":"Stopped: *","provider":"fake","stopped":true}
```

### should press a menu item when not stopped

```execute
AUX4_COMPUTER_STATE_DIR=$PWD/state aux4 computer menu --provider fake --app Calculator --path "View > Scientific"
```

```expect
fake menu app=Calculator path=View > Scientific
```
