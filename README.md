# Website

This repository holds the code for the server running at [azul3d.org](https://azul3d.org).

Changes accepted into the `master` branch of this repository are automatically served from `azul3d.org`. Just submit a PR, if your changes are accepted they will show up automatically at `azul3d.org` shortly thereafter.

## Overview

| Folder      | Description                                             |
|-------------|---------------------------------------------------------|
| `content`   | arbitrary files served at `azul3d.org/content`.         |
| `pages`     | Markdown files rendered into HTML and served under `/`. |
| `mdattr`    | Go package for parsing Markdown file attributes.        |
| `templates` | Go HTML templates used to render the markdown files.    |

## Running Locally

```
go get -u -d azul3d.org/website
cd $GOPATH/src/azul3d.org/website
go build
./website -https= -http=:8080 -update=false
```

Notes

- Conveniently, the `-d` flag instructs `go get` to not _install_, but rather just download, the `website` command.
- `go build` places the `website` binary into the current directory.
- `-https=` tells it not to serve over HTTPS, because you don't have the certificates.
- `-update=false` instructs the server to not pull changes from the remote Git repository.
