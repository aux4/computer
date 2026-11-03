# computer providers

Lists the installed providers. A test-local `.aux4` adds fake providers; `--provider` probes an extra name.

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

## a provider that speaks the contract

### should report its info and contract

```execute
aux4 computer providers --provider fake | jq -c '.providers[] | select(.provider == "fake")'
```

```expect
{"contract":1,"default":false,"info":{"capabilities":{"background":true},"contract":1,"platform":"test/none","provider":"fake","version":"0.0.0"},"package":"aux4/computer-fake","provider":"fake","status":"ok"}
```

## a provider without info

### should be listed as legacy instead of failing

```execute
aux4 computer providers --provider legacy | jq -c '.providers[] | select(.provider == "legacy")'
```

```expect
{"contract":null,"default":false,"error":"provider has no 'info' command (it predates computer contract 1); update aux4/computer-legacy","package":"aux4/computer-legacy","provider":"legacy","status":"legacy"}
```

## a provider with another contract

### should be listed as a mismatch

```execute
aux4 computer providers --provider future | jq -c '.providers[] | select(.provider == "future") | {provider, status, contract, error}'
```

```expect
{"provider":"future","status":"mismatch","contract":2,"error":"provider implements computer contract 2; aux4/computer implements contract 1"}
```

## providers that are not installed

### should list their packages

```execute
aux4 computer providers --provider nosuch | jq -c '.notInstalled[] | select(.provider == "nosuch")'
```

```expect
{"package":"aux4/computer-nosuch","provider":"nosuch"}
```

### should hint at the package for this OS when it is missing

```execute
AUX4_COMPUTER_OS=windows aux4 computer providers | jq -c '{default, contract, platform: (.platform | startswith("windows/")), hint}'
```

```expect
{"default":"win","contract":1,"platform":true,"hint":"No provider for this OS is installed. Install it with: aux4 aux4 pkger install aux4/computer-windows"}
```

### should have no default on an OS without a provider

```execute
AUX4_COMPUTER_OS=plan9 aux4 computer providers | jq -c '{default, hint}'
```

```expect
{"default":"","hint":null}
```

## argument validation

### should reject an invalid provider name

```execute
aux4 computer providers --provider 'Bad Name'
```

```error
{"error":"Invalid --provider 'Bad Name': use a provider name such as mac, linux or win"}
```
