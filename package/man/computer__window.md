#### Description

Window commands, dispatched to the provider for this operating system: `list`, `move`, `resize`, `minimize`,
`restore`, `close`, `fullscreen` and `raise`. Every command except `list` changes state, so the stop switch
refuses it.

#### Usage

```bash
aux4 computer window <list|move|resize|minimize|restore|close|fullscreen|raise> [flags]
```

#### Example

```bash
aux4 computer window list TextEdit
```
