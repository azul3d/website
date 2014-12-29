+Title = Graphics: Version 2

Azul3D has an extensive set of graphics programming packages, source code for these Go packages can be found in the [gfx repository](https://github.com/azul3d/gfx).

- [Core Packages](#core-packages)
- Version History
  - [Version 2](#version-2)
  - [Version 1.0.1](#version-1-0-1)
  - [Version 1](#version-1)

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
- [Announcement](/news/2015/gfx-version-2.html)
- [Changes](/gfx/v2-changes.html)

## Version 1.0.1

- `import "azul3d.org/gfx.v1"`
- [Announcement](/news/2014/gfx-gfx-gl2-minor-update.html)
- Fixed a bug causing Transforms to be constantly recalculated (see [issue #16](https://github.com/azul3d/gfx/issues/16)).

## Version 1

- `import "azul3d.org/gfx.v1"`
- Initial implementation.

