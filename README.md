# Website

This repository holds the code for the server running at [azul3d.org](https://azul3d.org).

## Layout

| Folder | Description |
|--------|-------------|
| `content` | arbitrary files served at `azul3d.org/content`. |
| `pages` | Markdown files rendered into HTML and served under `/`. |
| `mdattr` | Go package for parsing Markdown file attributes. |

## Installation

```
go get -u -d azul3d.org/website
cd $GOPATH/src/azul3d.org/website
go build
./website -https=
```

Notes

- Conveniently, the `-d` flag instructs `go get` to not _install_, but rather just download, the `website` command.
- `go build` places the `website` binary into the current directory.
- `-https=` tells it not to serve over HTTPS, because you don't have the certificates.
