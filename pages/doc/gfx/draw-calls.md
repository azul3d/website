# gfx - Draw Calls

Draw calls occur any time that you *draw* an object to a canvas. Understanding *draw calls* is critical to diagnosing performance of 3D applications.

 * [What are they?](#what-are-they)
 * [Performance trade-off](#performance-trade-off)
 * [Example - Voxels](#example-voxels)
 * [Example - Sprites](#example-sprites)

# What are they?

Using the gfx package a draw call is done through the `gfx.Canvas` interface, like so:

`canvas.Draw(image.Rect(0,0,0,0), obj, cam)`

Where `obj` is a `gfx.Object` and `cam` is a `gfx.Camera`.

In literal OpenGL code the above draw operation would equate to either `glDrawArrays` or `glDrawElements` (depending on whether or not the mesh is indexed). Technically each `canvas.Draw` operations makes `N` draw calls, where `N` is the number of `gfx.Mesh` that the `gfx.Object` being drawn contains.

# Performance

The performance of draw calls depends primarily upon **the number of draw calls** and the **speed of the CPU**. Here is a simple example:

 * *Fast*: 1 draw call, 10,000 triangles/call.
 * *Also Fast*: 2 draw calls, 5,000 triangles/call.
 * *Slow*: 100 draw calls, 100 triangles/call.

*High numbers of draw-calls primarily max out the CPU -- not the GPU.*

For more information the performance of draw-calls, see the NVIDIA slides *["Batch, Batch, Batch:" What Does It Really Mean?](http://www.nvidia.com/docs/IO/8228/BatchBatchBatch.pdf)*.

# Performance trade-off

A naive thought would be *"Well, just draw the entire scene in one draw call then!"*. It's certainly possible to do -- but keep in mind that a single draw call is restricted to:

 * One single graphics state (one shader, one set of textures, one alpha mode, etc).
 * One draw order (triangles are drawn in the order that they appear in the mesh).

Thus you can see *the trade-off of dynamism* in place. If you want to use multiple shaders, have objects with different textures, etc then *you need several draw calls*.

# Example - Voxels

Voxel-based games often make sections of multiple blocks effectively called *chunks*. With our new knowledge of *draw calls* we can explain one of the things that *chunks* solved for these games:

 * Many draw calls are expensive.
 * Putting all of the blocks into one single giant mesh makes updates (read: *creating/destroying blocks*) slow.
 * The idea: break things up into medium sized *chunks*.
  * Small enough that dynamic updates are fast, the smaller the chunk the more fast dynamic updates are.
  * Large enough that we have *lesser draw calls*, the larger the chunk the faster rendering is.

# Example - Sprites

For another example let's take simple non-animated sprites. A sprite is a square card made up of two triangles (a quad) and a `gfx.Texture` (i.e. the sprite sheet image). Lets say we want to draw *500 sprites*, there are two cases:

 1. Each sprite is it's own `gfx.Object` (i.e. no *"batching"*):
  * 500 draw calls (more draw calls means more CPU will be used: *slower overall performance*).
  * Sprites can pick-and-choose any shader, texture, etc.
  * Sprites can be drawn in any order.
  * Each sprite can move, scale, rotate, etc independently (without merging all of the objects into a *"batch"*).
 2. Each sprite is merged into one single `gfx.Object`, called a *batch*.
  * 1 draw call (the CPU and GPU will be working in parralel: better overall performance).
  * Each sprite must use the same exact shader, texture, etc.
  * Sprites can only be drawn in one single order (their natural draw order).
  * Moving, scaling, rotating sprites required merging all of the objects into a *batch* again.

The best solution here is to *know your target, and test*. Does the sprite need to move constantly, and use it's own dedicated shader, textures, etc? Or will it be mostly static, could it share shaders and textures with many other sprites? The performance you acquire depends highly on how well you can answer these questions.

Note that with modern hardware, 500 draw calls is *very little* and you could get away with many more than that -- *it's just an example*.
