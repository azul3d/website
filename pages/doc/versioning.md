# Versioning

Azul3D packages are [semantically versioned](http://semver.org). The implementation allows you to import a specific version of a package, so that future API-incompatible changes do not break your programs.

* [The Basics](#the-basics)
* [Major Versions Restriction](#major-versions-restriction)
* [Version Zero](#version-zero)
* [Pre-Release Versions](#pre-release-versions)
* [Implementation](#implementation)

# The Basics

To import a Go package from GitHub you would write something like this: 

```
import "github.com/user/example"
```

If you ran `go get github.com/user/example` it would download the most recent version of that package. This is problematic because if future API-incompatible changes are made to that package, your program will no longer build.

Instead of importing packages in the above way, the `azul3d.org` website hosts a special type of proxy application which internally does a bit of magic -- ultimitely it forwards your requests to one of [our GitHub](https://github.com/azul3d) repositories.

For instance to import the `gfx` package you can simply write:

```
import "azul3d.org/gfx.v1"
```

If you run `go get azul3d.org/gfx.v1` the proxy application running at `azul3d.org` will do some magic such that you will always get an API-compatable version of that package.


Even though the import path ends with the version number `.v1`, *the package is still referenced in code the same exact way*:

```
gfx.Something()
```

This is because Go analyzes the `package gfx` statement from the Go source files -- and doesn't care about the file path.

# Major Versions Restriction

If a program imports two seperate versions of the same package -- bad things can happen.

Generally speaking this is *not allowed* -- but it's ultimately your responsibility to ensure it doesn't happen. To remedy this issue slightly, **only major versions are allowed in the package path**:

```
// allowed
import "azul3d.org/gfx.v1"

// not allowed
import "azul3d.org/gfx.v1.2"
import "azul3d.org/gfx.v1.2.3"
```

This ensures that all applications using the `v1` API version are importing the same package.

Note that the major version in an import *always points to the latest minor and patch version*, for instance if four versions exist:

```
azul3d.org/gfx.v1
  -> v2.0.0 (not chosen, major version is not the same)
  -> v1.2.1 (chosen, it's the latest version)
  -> v1.2.0
  -> v1.0.0
```

# Version Zero

In accordance with section 2. of the [semantic versioning](http://semver.org) specification, which reads:

> Major version zero (0.y.z) is for initial development. Anything may change at any time. The public API should not be considered stable.

A major version zero is for very initial development: *A package whose major version is zero should not be considered stable*.

```
// unstable
import "azul3d.org/pkg.v0"

// stable
import "azul3d.org/pkg.v1"
```

# Pre-Release Versions

In accordance with section 9. of the [semantic versioning](http://semver.org) specification, which reads:

> A pre-release version MAY be denoted by appending a hyphen and a series of dot separated identifiers immediately following the patch version. Identifiers MUST comprise only ASCII alphanumerics and hyphen [0-9A-Za-z-]. Identifiers MUST NOT be empty. Numeric identifiers MUST NOT include leading zeroes. Pre-release versions have a lower precedence than the associated normal version. A pre-release version indicates that the version is unstable and might not satisfy the intended compatibility requirements as denoted by its associated normal version. Examples: 1.0.0-alpha, 1.0.0-alpha.1, 1.0.0-0.3.7, 1.0.0-x.7.z.92.

A *single pre-release version* is supported by prefixing `-dev` onto the existing version number, like so:

```
// in-development
import "azul3d.org/pkg.v2-dev"

// not released yet
import "azul3d.org/pkg.v2"

// released already
import "azul3d.org/pkg.v1"
```

This special extension should only be used if you need features only found in the in-development versions of packages, but most of the time you should never use it and instead stick with the most-recently released version.

# Implementation

In the implementation we make use of Git tags and branches named after the version (e.g. `v1` or `v1.2.1`).

If both a branch and tag of the same name and version exist (e.g. tag `v1` and branch `v1` both exist), then the first choice *is always the branch*, and the second is the tag. Otherwise the latest version is always used (i.e. tag `v1.1` will be chosen over branch `v1`).

Note that the application at `azul3d.org` is truly just a special proxy application -- all of the code is still hosted [on GitHub](https://github.com/azul3d).

If you want to see how all of this happens you can [view the source](http://github.com/azul3d/appengine).

