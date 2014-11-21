# Better Package Docs, Versioning
<p class="date">October 7, 2014</p>

We've made several improvements to package API documentation and versioning. We now link directly to [GoDoc](https://godoc.org) for package documentation and we're *now truly using* [Semantic Versioning](http://semver.org/).

No worries: *Everything we've changed is completely backwards compatible!*

# Documentation

If you visited our online package API documentation in the past, you would have noticed that we generate and host all of it ourselves.

If you did, you might have also noticed how completely sub-par it was in many different ways to [GoDoc](https://godoc.org) documentation. It lacked things like the ability to follow types by clicking them, keyboard shortcuts, etc.

In the past there were [circumstances](https://github.com/azul3d/issues/issues/26) that caused us to be unable to utilize *godoc.org* fully, but we've long since relieved those circumstances. *And today we've completely removed our sub-par API documentation and directly replaced it with godoc.org*!

If you visit the [packages page](/packages.html) you'll notice that it now directly links to [GoDoc](https://godoc.org) for every package's API documentation.

Going forward this will give us better package documentation and allow us to focus more time on [appropriate things](/doc/roadmap.html). Maybe we can even contribute to `godoc.org` / `gddo` directly and provide benefits for all Gophers!

# Semantic Versioning

After making some [bad statements](https://github.com/azul3d/issues/issues/25) and publishing a, *both invalid and needless*, survey -- we learned of something very important: *Semantic Versioning*.

Take note, buckle up, and do whatever you have to do to convince yourself to just read [the specification](http://semver.org). It really will tell you all that you need to know about versioning, and it's super short (in comparison to most standards).

Funny enough, what we were doing previously was actually written in the specification:

> This is not a new or revolutionary idea. In fact, you probably do something close to this already. The problem is that "close" isn't good enough.

And -- well -- *they are right*.

So what changed?

 * *All changes are backwards compatible*: existing code will continue to work.
 * Only major versions are allowed in import paths now (e.g. `v1` is allowed, `v1.1` is not).
 * Import the GLFW package at `native/glfw.v3` instead of `native/glfw.v3.1`.
 * Import in-development packages using `pkg.vN-dev` instead of `pkg.dev` (SemVer section 9).
 * Packages with no stable `v1` releases yet are imported under `v0` (SemVer section 4).

The [versioning page](/doc/versioning.html) has more information.

# semver.v1

What is it?

 * *Semantic Versioning for Go packages.*
 * [![GoDoc](http://godoc.org/azul3d.org/semver.v1?status.svg)](http://godoc.org/azul3d.org/semver.v1)
 * Like [gopkg.in](http://gopkg.in), but it runs in your own Go HTTP server.
 * Folder-based packages (e.g. `mydomain/my/pkg.v1` -> `github.com/myuser/my-pkg`).
 * Git tags and branches (e.g. `v1` -> `tag/branch v1.3.2`).
 * Development branches (e.g. `import "pkg.v2-dev"`).

We are using it because *our top priority is backwards compatability with existing code*, and we're certain *azul3d.org* isn't going away any time soon.

You probably shouldn't use it unless you need to though. [gopkg.in](http://gopkg.in) is a great service as-is and it will always be online -- your personal website may not be.

The new `semver.v1` package is well-documented and has plenty of tests, going forward it will work very well for us and perhaps a few others -- just give us a shout, we'd love to hear from you as always.

# Conclusion

By switching to GoDoc we get better API documentation for all of our Go packages. By following semantic versioning more to-the-dot in our future work we will be better off than ever before, and the new `azul3d.org/semver.v1` package will help us with this.

*We want to point focus towards developing clean Go packages with well-defined backwards compatability guarantees, so that upgrading is easier and old applications still build.*

Take note that the new `semver.v1` package is *not related to game development*, so we chose not to write an article purely for it's announcement.

