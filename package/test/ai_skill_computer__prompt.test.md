# ai skill computer prompt

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

## the method

### should print the generic method

```execute
aux4 ai skill computer prompt --provider fake | head -1
```

```expect
# Using a computer with aux4/computer
```

### should cover the loop, background, stop and data-loss warnings

```execute
aux4 ai skill computer prompt --provider fake | grep -c -E '^## (The loop|Background and foreground|Safety)'
```

```expect
3
```

## provider notes

### should append the current provider's notes

```execute
aux4 ai skill computer prompt --provider fake | tail -3
```

```expect
## Platform notes: fake (aux4/computer-fake)

- Fake notes: nothing here is real.
```

### should still work for a provider without notes

```execute
aux4 ai skill computer prompt --provider legacy | tail -1
```

```expect
The installed aux4/computer-legacy predates the provider contract and has no `notes` command. Run `aux4 computer legacy --help` and `aux4 aux4 pkger man aux4/computer-legacy` for its platform specifics, and update it when you can.
```

### should say which package to install when the provider is missing

```execute
aux4 ai skill computer prompt --provider nosuch | tail -1
```

```expect
The provider for this computer is not installed, so every command will fail until it is: `aux4 aux4 pkger install aux4/computer-nosuch`.
```
