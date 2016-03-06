+Title = Installation

This page documents how to install Azul3D on any [supported platform](/doc/platform-support.html).

* [Dependencies](#dependencies)
* [Modularity](#modularity)
* [Examples](#examples)
* [From Scratch](#from-scratch)

## Dependencies

Azul3D requires a few system tools and libraries to be installed (like GCC/MinGW, FreeType, etc).

Detailed installation instructions for each [supported platform](/doc/platform-support.html):

* [Windows (386 / amd64)](/doc/install/windows.html)
* [Linux (386 / amd64)](/doc/install/linux.html)
* [Mac OS X (amd64)](/doc/install/osx.html)

## Examples

After [installing the dependencies](#dependencies), you can install *most of* the Azul3D packages by simply installing the examples using the [go tool](https://golang.org/cmd/go/):

```
go get -u -v azul3d.org/examples/...
```

And the example binaries will be installed into the $GOPATH/bin directory (all prefixed with `azul3d_` for convenience).

## From Scratch

If you don't want the examples, you can download just the Azul3D engine itself
by running:

```
go get -u -v azul3d.org/engine/gfx
```
