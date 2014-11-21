# Updating

There are two different types of updating an Azul3D package. One is where you wish to update to a newer and better package API (if one is available), and the other is where you wish to receive bug fixes or new features for the latest available version of a package.

* [Updating Versions](#updating-versions)
* [Getting Bug Fixes And Features](#getting-bug-fixes-and-features)

# Updating Versions

Any time API-incompatible changes must be made to an existing package, a new [version](/doc/versioning.html) is created.

Once a new version of a package is available, previous versions receive no further updates. As such it is important to periodically update your applications to newer versions of Azul3D packages. For example if an application imported the following package:

```
import "azul3d.org/chippy.v1"
```

And a newer version is available, (the package's [documentation](/chippy.v1#versions) would mention if a newer version was available). You would want to change each import statement to the newer version:

```
import "azul3d.org/chippy.v2
```

This can be a very tedious task if you have many imports. To make it significantly easier you can use a tool like [govers](http://godoc.org/launchpad.net/govers). To do this you first need to install it using the Go tool:

```
go get launchpad.net/govers
```

And then in the root of the source tree you want to change, you would run:

```
govers azul3d.org/chippy.v2
```

Which would rewrite all of the `v1` import paths to become `v2`. It will also warn you if you are importing code that uses a previous version of the same package -- which is a big no no.

You may still need to do some manual fixing yourself -- though. Since it's a new package version it means that some API incompatible change has been made, so your code may need to be updated to reflect the new changes.

# Getting Bug Fixes And Features

When we don't have to make API-incompatible changes, we don't. Instead these changes are made to the latest released version of the package.

So for instance: version one of the chippy package (azul3d.org/chippy.v1), may still receive bug fixes and new features as time goes on. You can track these changes easily through the package's [GitHub repository](https://github.com/azul3d/chippy).

Downloading all of the recent changes to a specific package is very easy to do using the update flag with the *go get* command:

```
go get -u azul3d.org/chippy.v1
```

If you want to update all of the Azul3D packages that you have downloaded, it is also easy using *...* file path expansion, like so:

```
go get -u azul3d.org/...
```

Which would update every Azul3D package that you have installed.

