+Title = Graphics: Version 2 Changes

- [Overview](#overview)
- [Moved Packages](#moved-packages)
- gfx: Changes
  - [Mesh Appension](#gfx-mesh-appension)
  - [Renderer Has Been Renamed](#gfx-renderer-has-been-renamed)
  - [Lines And Points](#gfx-lines-and-points)
  - [No Empty Rectangles](#gfx-no-empty-rectangles)
  - [Other Changes](#gfx-other-changes)
- [gfx/window: Changes](#gfxwindow-changes)
- [gfx/gl2: Changes](#gfxgl2-changes)
- [gfx/gfxsort: New Package](#gfxgfxsort-new-package)
- [Full Changelog](#full-changelog)

## Overview

TODO(slimsag): write an overview for v2.

## Moved Packages

As part of a [consolidation effort](https://github.com/azul3d/issues/issues/33) the co-dependant packages have moved into a single (versioned) repository.

| Before                     | After                      |
|----------------------------|----------------------------|
| `azul3d.org/gfx.v1`        | `azul3d.org/gfx.v2`        |
| `azul3d.org/gfx/window.v2` | `azul3d.org/gfx.v2/window` |
| `azul3d.org/gfx/gl2.v2`    | `azul3d.org/gfx.v2/gl2`    |
| `azul3d.org/clock.v1`      | `azul3d.org/gfx.v2/clock`  |

## gfx: Mesh Appension

Was [issue #21](https://github.com/azul3d/gfx/issues/21).

Multiple `gfx.Mesh` can now be appended together to form just a single mesh. This is done through the new `Mesh.Append` method.

Suprisingly, mesh appension is a rather cumbersome task and multiple meshes cannot always be safely appended together (i.e. if one has per-vertex colors and the other does not, do the colors get removed or new ones added?).

To remedy the above situation there is a new `MeshState` structure, and a new `Mesh.State` method which allows for one to query the existing state of a mesh (e.g. _"does it have per-vertex colors?"_). This allows for one to determine if two meshes can be appended together perfectly (i.e. they both have the same data).

## gfx: Renderer Has Been Renamed

The _Renderer_ interface has long been an incorrect name. Although _technically_ correct by it's definition, most people would assume that a _Renderer_ makes choices about how to present graphical objects on the screen. In reality a _Renderer_ is not so restrictive, and as such has been renamed to _Device_.

| Before                   | After               |
|--------------------------|---------------------|
| `gfx.Renderer`           | `gfx.Device`        |
| `gfx.GPUInfo`            | `gfx.DeviceInfo`    |
| `gfx.Renderer.GPUInfo()` | `gfx.Device.Info()` |

See [issue #60](https://github.com/azul3d/gfx/issues/60).

## gfx: Lines And Points

Meshes were traditionally drawn as just _triangles_, now you may choose how they are drawn by setting `Mesh.Primitive` equal to one of `gfx.Triangles`, `gfx.Lines`, or `gfx.Points`.

See [issue #48](https://github.com/azul3d/gfx/issues/48).

## gfx: No Empty Rectangles

In the past the drawing and clearing operates implicitly considered an empty rectangle `image.Rect(0, 0, 0, 0)` to mean _"the whole area"_. This was almost certainly the wrong choice and we've remedied this by removing this _so-called_ feature.

See [issue #15](https://github.com/azul3d/gfx/issues/15).

## gfx: Other Changes

Below is a list of other significant changes made to the _gfx_ package:

- `TexCoord` and `Color` are now valid types for use in the `Shader.Input` map and as data to `VertexAttrib` (see [#23](https://github.com/azul3d/gfx/issues/23)).
- Added a convenience `Mesh.Normals` slice for storing the normals of a mesh (see [issue #11](https://github.com/azul3d/gfx/issues/11)).
- The TexWrap mode `BorderColor` is not always present, e.g. in OpenGL ES 2 (see [issue #56](https://github.com/azul3d/gfx/issues/56)).
- Clarify: Some renderers, e.g. OpenGL ES, only support boolean occlusion queries (see [issue #57](https://github.com/azul3d/gfx/issues/57)).
- Split OpenGL and GLSL information out of DeviceInfo struct (See [issue #62](https://github.com/azul3d/gfx/issues/62)).
- A seperate structure is used for GLSL shader sources now (See [issue #63](https://github.com/azul3d/gfx/issues/63)).

## gfx/window: Changes

- Improved package documentation (see [#49](https://github.com/azul3d/gfx/pull/49)).
- Support for multiple windows (see [#38](https://github.com/azul3d/gfx/issues/38)).
- Exposed the main thread for clients that need it (see [#39](https://github.com/azul3d/gfx/issues/39)).
- Uses a 24/bpp framebuffer by default (see [#24](https://github.com/azul3d/gfx/issues/41)).
- The `gles2` build tag enables the use of the new OpenGL ES 2 renderer on desktops (see [#43](https://github.com/azul3d/gfx/issues/43)).

## gfx/gl2: Changes

- Uses one single OpenGL context instead of the previous two (see [#24](https://github.com/azul3d/gfx/issues/24)).
- Improved package documentation (See [issue #54](https://github.com/azul3d/gfx/issues/54) and [this commit](https://github.com/azul3d/gfx-gl2/commit/493f72dbb36547e394f2d4995ee7d74dbf7b86d4)).
- `gl2.Renderer` is now `gl2.Device` (See [#60](https://github.com/azul3d/gfx/issues/60)).
- `gl2.Device` is now an interface (See [#52](https://github.com/azul3d/gfx/issues/52)).
- `gl2.New` now takes option function parameters (See [#53](https://github.com/azul3d/gfx/issues/53)).
- Fix a caching failure of shader uniform locations (See [#58](https://github.com/azul3d/gfx/issues/58)).
- Assets are now (optionally) shared across multiple gl2 devices (See [#28](https://github.com/azul3d/gfx/issues/28)).

## gfx/gfxsort: New Package

- Sorting utilities from the `gfx` package moved here (See [#59](https://github.com/azul3d/gfx/issues/59)).
  - `ByDist` `ByState` and `InsertionSort`.

## Full Changes

A full list of all the changes commited made between v1.0.1 and v2 is available [here](https://github.com/azul3d/gfx/compare/v1.0.1...v2).

