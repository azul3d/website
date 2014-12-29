+Title = Graphics: Version 2 Changes

- [Overview](#overview)
- [Moved Packages](#moved-packages)
- gfx
  - [Share By Communicating](#gfx-share-by-communicating)
  - [Mesh Appension](#gfx-mesh-appension)
  - [Renderer Has Been Renamed](#gfx-renderer-has-been-renamed)
  - [Headroom For Other Graphics Devices](#gfx-headroom-for-other-graphics-devices)
  - [Lines And Points](#gfx-lines-and-points)
  - [No Empty Rectangles](#gfx-no-empty-rectangles)
  - [State Sharing](#gfx-state-sharing)
  - [Other Changes](#gfx-other-changes)
- gfx/window
  - [Better Event Names](#gfxwindow-better-event-names)
  - [Multiple Windows](#gfxwindow-multiple-windows)
  - [Other Changes](#gfxwindow-other-changes)
- gfx/gl2
  - [Better Warning Feedback](#gfxgl2-better-warning-feedback)
  - [Other Changes](#gfxgl2-other-changes)
- [gfx/clock: Changes](#gfxclock-changes)
- [gfx/gfxutil: New Package](#gfxgfxutil-new-package)
- [Full Changes](#full-changes)

## Overview

Please see the [announcement](/news/2015/gfx-version-2.html).

## Moved Packages

As part of an on-going [consolidation effort](https://github.com/azul3d/issues/issues/33) co-dependant packages have moved into a single (versioned) repository.

| Before                     | After                      |
|----------------------------|----------------------------|
| `azul3d.org/gfx.v1`        | `azul3d.org/gfx.v2`        |
| `azul3d.org/gfx/window.v2` | `azul3d.org/gfx.v2/window` |
| `azul3d.org/gfx/gl2.v2`    | `azul3d.org/gfx.v2/gl2`    |
| `azul3d.org/clock.v1`      | `azul3d.org/gfx.v2/clock`  |

## gfx: Share By Communicating

Probably one of the biggest API differences with the _gfx_ package in version two is that there is _no more locking_. It was cumbersome and hard to understand. We believe it was one of the major blocks holding the _gfx_ package back, and it really hits home with the Go article: "_[Share Memory By Communicating](https://blog.golang.org/share-memory-by-communicating)"_.

This mess:

```
object.Lock()
object.Textures[0].Lock()
object.Textures[0].Source = foo
object.Textures[0].Unlock()
object.Unlock()
```

Is now just:

```
object.Textures[0].Source = foo
```

See [issue #66](https://github.com/azul3d/gfx/issues/66).

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

## gfx: Headroom For Other Graphics Devices

The `gfx` package intends to be a generic Go graphics API. It shouldn't strictly tie you to OpenGL factoids -- even if those are what you'll be using _90% of the time_. To make ourselves completely clear though:

- No non-OpenGL (Microsoft's DirectX, Apple's Metal, AMD's Mandle) graphics devices are planned today.

But, in case someone did want to, we've made some headroom so that adding support would be more clear:

The `DeviceInfo` structure has many OpenGL and GLSL fields that will not be used for other non-OpenGL graphics API's. To aid this we've split OpenGL and GLSL information into two seperate structs, which `gfx.DeviceInfo` now has a non-nil pointer to if it's available (See [issue #62](https://github.com/azul3d/gfx/issues/62)).

We've also split `gfx.Shader` to have a seperate `gfx.GLSLSources` structure (See [issue #63](https://github.com/azul3d/gfx/issues/63)).

## gfx: Lines And Points

Meshes were traditionally drawn as just _triangles_, now you may choose how they are drawn by setting `Mesh.Primitive` equal to one of `gfx.Triangles`, `gfx.Lines`, or `gfx.Points`.

See [issue #48](https://github.com/azul3d/gfx/issues/48).

## gfx: No Empty Rectangles

In the past the drawing and clearing operates implicitly considered an empty rectangle `image.Rect(0, 0, 0, 0)` to mean _"the whole area"_. This was almost certainly the wrong choice and we've remedied this by removing this _so-called_ feature.

See [issue #15](https://github.com/azul3d/gfx/issues/15).

## gfx: State Sharing

Each `gfx.Object` today has it's own `gfx.State` which is incredibly nice as it allows _objects_ to define their _graphics state_ fully. This is a radical shift from an OpenGL-like state-bound API. The downside to this was that multiple `gfx.Object` couldn't share _the same_ `gfx.State` -- this caused a memory overhead (important: _not a state change overhead_).

We've gone ahead and made `gfx.State` a pointer-based structure so that multiple `gfx.Object` can share the same state, much like how they share the same texture, shader, or mesh.

See [issue #83](https://github.com/azul3d/gfx/issues/83).

## gfx: Other Changes

Below is a list of other significant changes made to the _gfx_ package:

- `TexCoord` and `Color` are now valid types for use in the `Shader.Input` map and as data to `VertexAttrib` (see [issue #23](https://github.com/azul3d/gfx/issues/23)).
- Added a convenience `Mesh.Normals` slice for storing the normals of a mesh (see [issue #11](https://github.com/azul3d/gfx/issues/11)).
- The TexWrap mode `BorderColor` is not always present, e.g. in OpenGL ES 2 (see [issue #56](https://github.com/azul3d/gfx/issues/56)).
- Clarify: Some renderers, e.g. OpenGL ES, only support boolean occlusion queries (see [issue #57](https://github.com/azul3d/gfx/issues/57)).
- We now use `go generate` and the `stringer` command to generate `String` methods (see [issue #64](https://github.com/azul3d/gfx/issues/64)).

## gfx/window: Better Event Names

Three new EventMask group masks are defined, each one selects all cusor/mouse/keyboard events:

```
CursorEvents = CursorMovedEvents|CursorEnterEvents|CursorExitEvents
MouseEvents = MouseButtonEvents|MouseScrolledEvents
KeyboardEvents = KeyboardButtonEvents|KeyboardTypedEvents
```

The keyboard and mouse event names are more clear and now reflect each-other directly:

| Before                       | After                                                       |
|------------------------------|-------------------------------------------------------------|
| `window.MouseEvents`         | `window.MouseButtonEvents` (proper)                         |
| `mouse.Event`                | `mouse.ButtonEvent` (proper)                                |
| `window.KeyboardStateEvents` | `window.KeyboardButtonEvents` (matches `MouseButtonEvents`) |
| `keyboard.StateEvent`        | `keyboard.ButtonEvent` (matches `mouse.ButtonEvent`)        |
| `keyboard.TypedEvent`        | `keyboard.Typed` (matches `mouse.Scrolled`)                 |

## gfx/window: Other Changes

- Support for multiple windows (see [issue #38](https://github.com/azul3d/gfx/issues/38)).
- Support for run-time fullscreen<->windowed mode switching (see [issue #34](https://github.com/azul3d/gfx/issues/34)).
- Exposed the main thread for clients that need it (see [issue #39](https://github.com/azul3d/gfx/issues/39)).
- Improved package documentation (see [issue #49](https://github.com/azul3d/gfx/pull/49)).
- Uses a 24/bpp framebuffer by default (see [issue #24](https://github.com/azul3d/gfx/issues/41)).
- The `gles2` build tag enables the use of the new OpenGL ES 2 renderer on desktops (see [issue #43](https://github.com/azul3d/gfx/issues/43)).
- Makes use of mouse.v2 and keyboard.v2 now (see [this page](/news/2015/keyboard-mouse-version-2.html)).

## gfx/gl2: Better Warning Feedback

For shader inputs of invalid types that will be ignored, developers will now receive a debug warning in the console (See [issue #77](https://github.com/azul3d/gfx/issues/77)):

```
Shader input "Foo" uses an invalid shader input type "float64", ignoring.
```

Previously, if you tried to draw an object with a invalid shader input (or an incomplete `gfx.Object`), you would receive one warning each time you try to draw the object:

```
Shader input "Foo" uses an invalid shader input type "float64", ignoring.
Shader input "Foo" uses an invalid shader input type "float64", ignoring.
Shader input "Foo" uses an invalid shader input type "float64", ignoring.
Shader input "Foo" uses an invalid shader input type "float64", ignoring.
...
```

Now you will just receive one for each individual object (See [issue #86](https://github.com/azul3d/gfx/issues/86)).

## gfx/gl2: Other Changes

- Uses one single OpenGL context instead of the previous two (see [issue #24](https://github.com/azul3d/gfx/issues/24)).
- Improved package documentation (See [issue #54](https://github.com/azul3d/gfx/issues/54) and [this commit](https://github.com/azul3d/gfx-gl2/commit/493f72dbb36547e394f2d4995ee7d74dbf7b86d4)).
- `gl2.Renderer` is now `gl2.Device` (See [issue #60](https://github.com/azul3d/gfx/issues/60)).
- `gl2.Device` is now an interface (See [issue #52](https://github.com/azul3d/gfx/issues/52)).
- `gl2.New` now takes option function parameters (See [issue #53](https://github.com/azul3d/gfx/issues/53)).
- Fix a caching failure of shader uniform locations (See [issue #58](https://github.com/azul3d/gfx/issues/58)).
- Assets are now (optionally) shared across multiple gl2 devices (See [issue #28](https://github.com/azul3d/gfx/issues/28)).

## gfx/gfxutil: New Package

The distance and state sorting utilities (`ByDist`, `ByState`, and `InsertionSort`) from the _gfx_ package have moved to this new package (see [issue #59](https://github.com/azul3d/gfx/issues/59)).

There is a new `gfxutil.OpenShader` utility function that can be used to simply open a GLSL vertex and fragment shader (see [issue #73](https://github.com/azul3d/gfx/issues/73)).

And to match the above, there is also a new `gfxutil.OpenTexture` utility function to open texture image files (see [issue #74](https://github.com/azul3d/gfx/issues/74)).

## gfx/clock: Changes

The _clock_ package has also had a major overhaul and thorough review of all it's code. This included making it even more accurate at performing stable frame rate limiting for graphics hardware not supporting vertical-sync (most laptop chipsets).

A few symbol names have been shortened, Specifically _"Average"_ to just _"Avg"_.

| Before                       | After                    |
|------------------------------|--------------------------|
| `AverageFrameRate`           | `AvgFrameRate`           |
| `AverageFrameRateSamples`    | `AvgSampleCount`         |
| `SetAverageFrameRateSamples` | `SetAvgSampleCount`      |

See more about the review in [issue #71](https://github.com/azul3d/gfx/issues/71).

## Full Changes

A full list of all the changes commited made between v1.0.1 and v2 is available [here](https://github.com/azul3d/gfx/compare/v1.0.1...v2).

