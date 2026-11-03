# computer click

Generic commands dispatch to `aux4 computer <provider> <command>`. These tests use fake providers defined in a
test-local `.aux4`, so nothing real is clicked or typed.

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
        },
        {
          "name": "legacy",
          "execute": [
            "profile:computer:legacy"
          ],
          "help": {
            "text": "Fake provider that predates the contract (no info)"
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
        },
        {
          "name": "badinfo",
          "execute": [
            "profile:computer:badinfo"
          ],
          "help": {
            "text": "Fake provider whose info is not JSON"
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
          "name": "click",
          "execute": [
            "printf 'fake click x=%s y=%s element=%s app=%s mode=%s\\n' values(x, y, element, app, mode)"
          ],
          "help": {
            "text": "Echo a click",
            "variables": [
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
                "name": "element",
                "text": "element",
                "default": ""
              },
              {
                "name": "app",
                "text": "app",
                "default": ""
              },
              {
                "name": "mode",
                "text": "mode",
                "default": "auto"
              }
            ]
          }
        },
        {
          "name": "scroll",
          "execute": [
            "echo '{\"error\":\"App not running: Nope\"}' >&2; exit 3"
          ],
          "help": {
            "text": "Always fails",
            "variables": [
              {
                "name": "app",
                "text": "app",
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
          "name": "click",
          "execute": [
            "printf 'legacy click app=%s\\n' values(app)"
          ],
          "help": {
            "text": "Echo a click",
            "variables": [
              {
                "name": "app",
                "text": "app",
                "default": ""
              }
            ]
          }
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
          "name": "click",
          "execute": [
            "echo 'future click'"
          ]
        }
      ]
    },
    {
      "name": "computer:badinfo",
      "commands": [
        {
          "name": "info",
          "execute": [
            "echo 'not json'"
          ]
        },
        {
          "name": "click",
          "execute": [
            "echo 'badinfo click'"
          ]
        }
      ]
    }
  ]
}
```

## dispatching

### should pass the given flags to the provider and let it apply its own defaults

```execute
aux4 computer click --provider fake --x 34 --y 211 --app Calculator
```

```expect
fake click x=34 y=211 element= app=Calculator mode=auto
```

### should take the provider from AUX4_COMPUTER_PROVIDER

```execute
AUX4_COMPUTER_PROVIDER=fake aux4 computer click --element 0.1.2 --mode background
```

```expect
fake click x= y= element=0.1.2 app= mode=background
```

### should pass values with quotes and shell syntax through unchanged

```execute
aux4 computer click --provider fake --app "it's \$(touch pwned) \`touch pwned\`; touch pwned" && ls pwned 2>&1 | grep -c 'No such file'
```

```expect
fake click x= y= element= app=it's $(touch pwned) `touch pwned`; touch pwned mode=auto
1
```

### should pass negative coordinates

```execute
aux4 computer click --provider fake --x=-200 --y 10
```

```expect
fake click x=-200 y=10 element= app= mode=auto
```

### should pass values that look like flags or contain '='

```execute
aux4 computer click --provider fake --app=--mode=foreground --element "a=b=c"
```

```expect
fake click x= y= element=a=b=c app=--mode=foreground mode=auto
```

### should pass the provider's error and exit code through

```execute
aux4 computer scroll --provider fake --app Nope; echo "exit=$?"
```

```expect
exit=3
```

```error
{"error":"App not running: Nope"}
```

## provider checks

### should dispatch to a provider that predates the contract (no info)

```execute
aux4 computer click --provider legacy --app Calculator
```

```expect
legacy click app=Calculator
```

### should refuse a provider with another contract version

```execute
aux4 computer click --provider future; echo "exit=$?"
```

```expect
exit=1
```

```error
{"contract":2,"error":"Computer provider 'future' (aux4/computer-future) is not compatible: provider implements computer contract 2; aux4/computer implements contract 1. Update both with: aux4 aux4 pkger install aux4/computer aux4/computer-future","expected":1,"package":"aux4/computer-future","provider":"future"}
```

### should refuse a provider whose info is not JSON

```execute
aux4 computer click --provider badinfo; echo "exit=$?"
```

```expect
exit=1
```

```error:partial
'info' did not return a JSON object
```

### should name the package to install when the provider is missing

```execute
aux4 computer click --provider nosuch --x 1 --y 1; echo "exit=$?"
```

```expect
exit=1
```

```error
{"error":"No computer provider 'nosuch' is installed. Install it with: aux4 aux4 pkger install aux4/computer-nosuch","package":"aux4/computer-nosuch","provider":"nosuch"}
```

### should pick the Linux provider on Linux

```execute
AUX4_COMPUTER_OS=linux aux4 computer click --x 1 --y 1; echo "exit=$?"
```

```expect
exit=1
```

```error
{"error":"No computer provider 'linux' is installed. Install it with: aux4 aux4 pkger install aux4/computer-linux","package":"aux4/computer-linux","provider":"linux"}
```

### should pick the win provider on Windows

```execute
AUX4_COMPUTER_OS=windows aux4 computer click --x 1 --y 1; echo "exit=$?"
```

```expect
exit=1
```

```error
{"error":"No computer provider 'win' is installed. Install it with: aux4 aux4 pkger install aux4/computer-windows","package":"aux4/computer-windows","provider":"win"}
```

### should fail on an operating system without a provider

```execute
AUX4_COMPUTER_OS=plan9 aux4 computer click --x 1 --y 1; echo "exit=$?"
```

```expect
exit=1
```

```error
{"error":"No computer provider for this operating system (plan9). Pass --provider <name> to use one explicitly.","os":"plan9","providers":["mac","linux","win"]}
```

### should reject an invalid provider name

```execute
aux4 computer click --provider 'mac; rm -rf x'; echo "exit=$?"
```

```expect
exit=1
```

```error
{"error":"Invalid --provider 'mac; rm -rf x': use a provider name such as mac, linux or win"}
```

### should explain that the Windows provider is win

```execute
aux4 computer click --provider windows; echo "exit=$?"
```

```expect
exit=1
```

```error
{"error":"Invalid --provider 'windows': the Windows provider is named 'win' ('windows' is the command that lists windows)"}
```

## stop switch

```beforeEach
rm -rf state
```

```afterAll
rm -rf state
```

### should refuse to dispatch while the stop switch is on

```execute
export AUX4_COMPUTER_STATE_DIR=$PWD/state; aux4 computer stop > /dev/null; aux4 computer click --provider fake --x 1 --y 2 --app Calculator; echo "exit=$?"
```

```expect
exit=1
```

```error:partial
{"command":"click","error":"Stopped: aux4 computer is not sending input or changing state (*/state/STOP exists). Run 'aux4 computer resume' to allow it again.","provider":"fake","stopped":true}
```

### should refuse when AUX4_COMPUTER_STOP=1

```execute
AUX4_COMPUTER_STATE_DIR=$PWD/state AUX4_COMPUTER_STOP=1 aux4 computer click --provider fake --x 1 --y 2; echo "exit=$?"
```

```expect
exit=1
```

```error
{"command":"click","error":"Stopped: aux4 computer is not sending input or changing state (AUX4_COMPUTER_STOP=1 is set). Run 'aux4 computer resume' to allow it again.","provider":"fake","stopped":true}
```

### should record the refusal in the action log

```execute
export AUX4_COMPUTER_STATE_DIR=$PWD/state; aux4 computer stop > /dev/null; aux4 computer click --provider fake --x 1 --y 2 --app Calculator --mode background 2>/dev/null; aux4 computer log --limit 1 | jq -c '.entries[0] | {provider, command, app, args, mode, result, stopped, source}'
```

```expect
{"provider":"fake","command":"click","app":"Calculator","args":{"app":"Calculator","mode":"background","x":"1","y":"2"},"mode":"background","result":"error","stopped":true,"source":"core"}
```

### should dispatch again after resume

```execute
export AUX4_COMPUTER_STATE_DIR=$PWD/state; aux4 computer stop > /dev/null; aux4 computer resume > /dev/null; aux4 computer click --provider fake --x 1 --y 2
```

```expect
fake click x=1 y=2 element= app= mode=auto
```
