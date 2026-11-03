#### Description

Clipboard commands, dispatched to the provider for this operating system: `copy` (put text on the clipboard, or copy
an app's selection), `paste` (paste into an app) and `read`. `copy` and `paste` are refused by the stop switch;
`read` is read-only.

#### Usage

```bash
aux4 computer clipboard <copy|paste|read> [flags]
```

#### Example

```bash
aux4 computer clipboard read
```
