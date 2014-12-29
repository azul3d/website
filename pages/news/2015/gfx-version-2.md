+Title = gfx: version 2
+Date  = January 1, 2015

To celebrate the beginning of a new year, we're releasing version 2 of the gfx packages. We're incredibly excited to see what this year brings for game development in Go and can't wait to see what people create.

What follows is _an overview_, for an in-depth list of the changes that gfx.v2 brings, please see [this page](/gfx/v2-changes.html).

## Overview

Version 2 of the gfx package brings support for OpenGL ES and WebGL / GopherJS.

_*This opens the possibility of Android and HTML5 applications made in Go using Azul3D.*_

It adds support for [line and point drawing](/gfx/v2-changes.html#gfx-lines-and-points), enables more performant [state sharing](/gfx/v2-changes.html#gfx-state-sharing) and brings everything together with [better event names](/gfx/v2-changes.html#gfxwindow-better-event-names) and support for [multiple windows](/gfx/v2-changes.html#gfxwindow-multiple-windows).

For a further detailed list of changes [see here](/gfx/v2-changes.html).

## Updating

Install gfx.v2 using _go get_:

```
go get azul3d.org/gfx.v2
```

As part of an on-going [consolidation effort](https://github.com/azul3d/issues/issues/33) co-dependant packages have moved into a single (versioned) repository. You will need to update your import paths as needed.

| Before                     | After                      |
|----------------------------|----------------------------|
| `azul3d.org/gfx.v1`        | `azul3d.org/gfx.v2`        |
| `azul3d.org/gfx/window.v2` | `azul3d.org/gfx.v2/window` |
| `azul3d.org/gfx/gl2.v2`    | `azul3d.org/gfx.v2/gl2`    |
| `azul3d.org/clock.v1`      | `azul3d.org/gfx.v2/clock`  |

As this is a _major version release_, it is API-incompatible with previous versions. You will need to review [the changes](/gfx/v2-changes.html) yourself and update your programs as needed.

If you encounter any issues updating or have questions, please [get in touch](/doc/community.html) -- we'd love to help!

## Conclusion

Users should expect to see better performance in applications if they take advantage of new abilities [like state sharing](/gfx/v2-changes.html#gfx-state-sharing), but overall roughly the same performance as in previous versions.

We believe that this release is a major step in the right direction of having a single uniform graphics API for Go. We're excited to see where we _Go_ in the future.
