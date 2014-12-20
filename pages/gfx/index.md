+Title = Graphics: Version 2

Azul3D has an extensive set of graphics programming packages, source code for these Go packages can be found in the [gfx repository](https://github.com/azul3d/gfx).

- [Packages](#packages)
- [Core Packages](#core-packages)
- Version History
  - [Version 2](#version-2)
  - [Version 1.0.1](#version-1-0-1)
  - [Version 1](#version-1)

## Packages

Following are the graphics packages which produce the most significant real-world results. They are built on top of [the graphics core](#core-packages).

| Package                         | Description                                           |
|---------------------------------|-------------------------------------------------------|
| [gfx.v2/sprite](/gfx.v2/sprite) | Sprite implements 2D sprite rendering.                |
| [gfx.v2/ice](/gfx.v2/ice)       | Ice is a 3D model and scene format.                   |
| [gfx.v2/cull](/gfx.v2/cull)     | Cull implements efficient camera view culling.        |
| [gfx.v2/text](/gfx.v2/text)     | Text implements rendering of text strings.            |

TODO(slimsag): make these or remove them from here?

## Core Packages

The packages below form _the graphics core_, which is the bases for all graphics related development.

| Core Package                    | Description                                           |
|---------------------------------|-------------------------------------------------------|
| [gfx.v2](/gfx.v2)               | A Go interface to modern GPU rendering API's.         |
| [gfx.v2/window](/gfx.v2/window) | Cross-platform window and graphics device management. |
| [gfx.v2/clock](/gfx.v2/clock)   | Measures performance of frame-based applications.     |
| [gfx.v2/gl2](/gfx.v2/gl2)       | OpenGL 2 based graphics device.                       |
| [gfx.v2/gles2](/gfx.v2/gles2)   | OpenGL ES 2 based graphics device.                    |
| [gfx.v2/webgl](gfx.v2/webgl)    | GopherJS/WebGL graphics device.                       |

## Version 2

- `import "azul3d.org/gfx.v2"`
- [Announcement](/news/2014/gfx-v2-released.html)
- [Changes](/gfx/v2-changes.html)

## Version 1.0.1

- `import "azul3d.org/gfx.v1"`
- [Announcement](/news/2014/gfx-gfx-gl2-minor-update.html)
- Fixed a bug causing Transforms to be constantly recalculated (see [issue #16](https://github.com/azul3d/gfx/issues/16)).

## Version 1

- `import "azul3d.org/gfx.v1"`
- Initial implementation.

