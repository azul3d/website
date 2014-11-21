# Important - Import Paths Have Changed
<p class="date">July 26, 2014</p>

During the [big move to GitHub](/news/2014/moving-from-google-code-to-github.html) we reconsidered the layout of the versioned import paths of Azul3D packages.

Since one of the goals in our move to GitHub was to split up Azul3D's packages into multiple smaller repositories, it makes sense that we would want to possibly make API-incompatible changes (by incremented a version number, as usual) of a specific package *independent* of other packages: something that we couldn't do with the old versioning scheme.

# What To Do

Updating to the new import path scheme is easy using a modified go-fix command:

```
go get azul3d.org/cmd/azulfix.v1
azulfix.v1 path/to/src/
```

Which will rewrite all `azul3d.org/v1` import paths to their new counterparts.

# Why?

Looking at the previous import paths we had paths that looked like this:</p>

```
azul3d.org/v1/pkg
azul3d.org/v1/pkg/subpkg
```

Because the version number is stored in the middle of the URL, it would not be possible to have packages with different versions because they are all under the same version number. Instead, by moving the version to the end, we can have import paths like this:

```
azul3d.org/pkg.v1
azul3d.org/pkg/subpkg.v1
```

Which allows us to increment the version numbers of specific packages independently of one-another. This is crucial going forward because it means we can break API-compatability of a single package, but leave other ones alone. In code however you will still reference the packages using just:

```
pkg.DoSomething()
subpkg.DoSomething()
```

All in all, upgrading to the new system is made easy by the *azulfix* tool and going forward this will allow for better code in many situations.

# Implementation

If any part of this article is confusing to you, or if you're interested in the implementation details, please see [this page](/doc/versioning.html).

