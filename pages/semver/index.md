+Title = semver: Semantic Versioning for Go packages

[![GoDoc](https://godoc.org/azul3d.org/semver.v2?status.svg)](https://godoc.org/azul3d.org/semver.v2) [![Build Status](https://travis-ci.org/azul3d/semver.svg)](https://travis-ci.org/azul3d/semver)

```
import "azul3d.org/semver.v2"
```

- [What is it?](#what-is-it)
- Version History
  - [Version 2.0.1](#version-201-changes)
  - [Version 2](#version-2-changes)
  - [Version 1.0.1](#version-101-changes)
  - [Version 1](#version-1-changes)

## What is it?

- [Semantic Versioning](http://semver.org/) for Go packages.
- Like [gopkg.in](http://gopkg.in), but it runs in your own Go HTTP server.
- Works great in combination with the [govers](https://github.com/rogpeppe/govers) tool.
- Folder-based packages (e.g. `mydomain.org/my/pkg.v1` -> `github.com/myuser/my-pkg`).
- Git tags and branches (e.g. `v1 -> tag/branch v1.3.2`).
- Unstable branches (e.g. `import "pkg.v2-unstable"`).

Also see the [versioning](/doc/versioning.html) and [updating](/doc/updating.html) pages.

## Version 2.0.1: Changes

- Fixed a documentation typo (see [#11](https://github.com/azul3d/semver/pull/11)).
- Fixed a critical bug that caused `v1` to be chosen over `v1.0.1` (see [#12](https://github.com/azul3d/semver/issues/12)).
- Added extensive tests for version choosing to avoid future issues (see [commit](https://github.com/azul3d/semver/commit/9f950773d7a302656368bae7747e1e7633429b17)).
- [Full Changes](https://github.com/azul3d/semver/compare/v2...v2.0.1)

## Version 2: Changes

- Development branches `pkg.v2-dev` replaced by _unstable_ branches `pkg.v2-unstable` (see [#7](https://github.com/azul3d/semver/issues/7)).
- A tool to migrate from `-dev` to `-unstable` is available, [see here](https://github.com/azul3d/semver/issues/7#issuecomment-70383909).
- Clicking on types on _godoc.org_ now brings you to the source code on GitHub (see [#10](https://github.com/azul3d/semver/issues/10)).
- Fixed a bug that caused import paths for packages with dashes (e.g. `go-thepkg`) to not work (see [#8](https://github.com/azul3d/semver/issues/8)).
- [Full Changes](https://github.com/azul3d/semver/compare/v1.0.1...v2)

## Version 1.0.1: Changes

- Fixed a bug that caused branches to resolve incorrectly (see [#2](https://github.com/azul3d/semver/issues/2)).
- [Full Changes](https://github.com/azul3d/semver/compare/v1...v1.0.1)

## Version 1: Changes

- Initial release.
- [Full Changes](https://github.com/azul3d/semver/commits/v1)
