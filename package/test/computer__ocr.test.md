# computer ocr

Generic commands dispatch to `aux4 computer <provider> <command>`. This fake provider returns the same
JSON shape a real provider does, so these tests cover `--format text` (and that the default `json`
format is unchanged) without a real screen.

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
            "text": "Fake provider that returns canned ocr JSON"
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
          "name": "ocr",
          "execute": [
            "if [ \"${app}\" = \"Calculator\" ]; then echo '{\"app\":\"Calculator\",\"matches\":[{\"text\":\"AC\",\"center\":{\"x\":100,\"y\":500},\"frame\":{\"x\":80,\"y\":480,\"width\":40,\"height\":40}},{\"text\":\"%\",\"center\":{\"x\":220,\"y\":500},\"frame\":{\"x\":200,\"y\":480,\"width\":40,\"height\":40}},{\"text\":\"7,542\",\"center\":{\"x\":300,\"y\":100},\"frame\":{\"x\":260,\"y\":85,\"width\":80,\"height\":30}}]}'; elif [ \"${app}\" = \"NoSuchApp\" ]; then echo '{\"matches\":[]}'; else echo '{\"error\":\"Text recognition did not finish within 1s\",\"phase\":\"recognition\",\"timeout\":1,\"waited\":1}' >&2; exit 1; fi"
          ],
          "help": {
            "text": "Echo canned ocr JSON, or a JSON error (and exit 1) for any other app",
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
    }
  ]
}
```

## default format

### should return the raw json unchanged when --format is not given

```execute
aux4 computer ocr --provider fake --app Calculator
```

```expect
{"app":"Calculator","matches":[{"text":"AC","center":{"x":100,"y":500},"frame":{"x":80,"y":480,"width":40,"height":40}},{"text":"%","center":{"x":220,"y":500},"frame":{"x":200,"y":480,"width":40,"height":40}},{"text":"7,542","center":{"x":300,"y":100},"frame":{"x":260,"y":85,"width":80,"height":30}}]}
```

### should return the same raw json when --format json is given explicitly

```execute
aux4 computer ocr --provider fake --app Calculator --format json
```

```expect
{"app":"Calculator","matches":[{"text":"AC","center":{"x":100,"y":500},"frame":{"x":80,"y":480,"width":40,"height":40}},{"text":"%","center":{"x":220,"y":500},"frame":{"x":200,"y":480,"width":40,"height":40}},{"text":"7,542","center":{"x":300,"y":100},"frame":{"x":260,"y":85,"width":80,"height":30}}]}
```

## --format text

### should group matches into rows and sort each row left-to-right, with no coordinates

```execute
aux4 computer ocr --provider fake --app Calculator --format text
```

```expect
7,542
AC  %
```

### should print nothing (one blank line) when there are no matches

```execute
aux4 computer ocr --provider fake --app NoSuchApp --format text
```

```expect

```

### should pass a provider error through unchanged, same as the json format

```execute
aux4 computer ocr --provider fake --app Stuck --format text; echo "exit=$?"
```

```expect
exit=1
```

```error
{"error":"Text recognition did not finish within 1s","phase":"recognition","timeout":1,"waited":1}
```
