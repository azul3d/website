+Title = Let's try a monorepo (or two)!
+Date  = March 5, 2016

To make our lives as developers easier, and to enable us to be more productive
in the limited time we have, we've decided to migrate Azul3D into two monorepos
instead of the monstrosity we had before.

- Engine: https://github.com/azul3d/engine
- Examples: https://github.com/azul3d/examples

If you're eager to learn more about what is going on and why, check out the [very detailed issue on the topic.](https://github.com/azul3d/engine/issues/1)

## Import Path Change

All import paths are now under `azul3d.org/engine/...`, for example:

| Before                | After                     |
|-----------------------|---------------------------|
| `azul3d.org/gfx.v1`   | `azul3d.org/engine/gfx`   |
| `azul3d.org/audio.v1` | `azul3d.org/engine/audio` |

Likewise the examples are now at `azul3d.org/examples/...` and not `azul3d.org/examples.v1/...`.

## Vendoring

- Users are now expected to vendor `azul3d.org/engine` packages using tools like godep, glide, govendor, etc. as they would with normal packages.
- Development speed will not be sacrificed in order to maintain backwards compatibility, so it's highly reccomended that you vendor Azul3D packages when
writing applications.

## Released Packages

The following packages have also been released as part of this whole migration:

| Package                 | In-development branch released     |
|-------------------------|------------------------------------|
| `azul3d.org/keyboard`   | `v2-unstable`                      |
| `azul3d.org/mouse`      | `v2-unstable`                      |
| `azul3d.org/native/al`  | `v1.1-unstable`                    |
| `azul3d.org/examples`   | `gfxv2`.                           |
| `azul3d.org/gfx`        | `v2-unstable`                      |
| `azul3d.org/audio`      | `v2-dev`                           |
| `azul3d.org/audio/wav`  | `v1.1-dev`                         |
| `azul3d.org/audio/flac` | `master`   (not previously public) |

And `azul3d.org/native/glfw.v5` is now imported at `github.com/go-gl/glfw/v3.1/glfw` (the official location for this repository -- as always).

## And last but not least.. an apology for the noise

You weren't the only one that got a bit spammed by all this progress, and we can
only hope you got less than the three pages we got!

![sorry](http://i.imgur.com/s2Mi79j.png)
