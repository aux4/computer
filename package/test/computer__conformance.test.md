# computer conformance

## a missing provider

### should name the package to install

```execute
aux4 computer conformance --provider nosuch; echo "exit=$?"
```

```expect
exit=1
```

```error
{"error":"No computer provider 'nosuch' is installed. Install it with: aux4 aux4 pkger install aux4/computer-nosuch","package":"aux4/computer-nosuch","provider":"nosuch"}
```
