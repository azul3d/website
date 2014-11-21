# Installation

This page documents how to install Azul3D on any [supported platform](/doc/platform-support.html).

* [Dependencies](#dependencies)
* [Modularity](#modularity)
* [Examples](#examples)
* [From Scratch](#from-scratch)

# Dependencies

Azul3D requires a few system tools and libraries to be installed (like GCC/MinGW, FreeType, etc).

Detailed installation instructions for each [supported platform](/doc/platform-support.html):

* [Windows (386 / amd64)](/doc/install/windows.html)
* [Linux (386 / amd64)](/doc/install/linux.html)
* [Mac OS X (amd64)](/doc/install/osx.html)

# Modularity

Azul3D is very modular, each package is installed separately. This is very useful because you only need to install the specific Azul3D package that you intend to use (a full list of packages can be found on the [packages page](/packages.html)).

Want to use the `audio` package but not the `gfx` package? Simply just install the `audio` package:

```
go get -u azul3d.org/audio.v1
```

# Examples

After [installing the dependencies](#dependencies), you can install *most of* the Azul3D packages by simply installing the examples using the [go tool](https://golang.org/cmd/go/):

```
go get azul3d.org/examples.v1
go get azul3d.org/examples.v1/...
```

And the example binaries will be installed into the $GOPATH/bin directory.

# From Scratch

The best way to install Azul3D is by [installing the examples](#examples), optionally however you can install Azul3D packages from scratch.

For instance if you wrote an `example` package with these import statements:

```
import(
	_ "azul3d.org/audio/wav.v1"
	"azul3d.org/audio.v1"
)
```

If you tried to build that program and didn't have those packages installed yet, you would receive the following:

```
/godev/src/example$ go build
example.go:5:2: cannot find package "azul3d.org/audio.v1" in any of:
	/go/src/pkg/azul3d.org/audio.v1 (from $GOROOT)
	/godev/src/azul3d.org/audio.v1 (from $GOPATH)
example.go:4:2: cannot find package "azul3d.org/audio/wav.v1" in any of:
	/go/src/pkg/azul3d.org/audio/wav.v1 (from $GOROOT)
	/godev/src/azul3d.org/audio/wav.v1 (from $GOPATH)
```

These errors are telling you that you haven't installed the `audio` and `audio/wav` packages.

**Installing everything the above program needs is very easy**:

```
go get -u example
```

Every package needed by the `example` package will be downloaded and installed.

There are other ways to install Azul3D packages from scratch, like downloading each package independantly via `go get`, but using the above command is by far the easiest way to download and install everything your program depends on.

