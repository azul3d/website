# gfx/window: version 2
<p class="date">September 10, 2014</p>

Azul3D has for a very long time had a [gfx/window](packages.html#gfx-window) package.

[Version 2](/gfx/window.v2) of the package was released in unison with recent [Mac OS X support](/news/2014/mac-osx-support.html) and is outlined here.

# Past

In the past, the [gfx/window](packages.html#gfx-window) package had actually just served as a *basic implementation* of how one **might** create a window using *Chippy* and initialize the OpenGL 2 renderer ([gfx/gl2]([gfx/window](packages.html#gfx-gl2)).

**The package clocked in at just one file, and 156 lines of code:** the idea was that the package would be so simple that -- if you *really* wanted it to do something else -- *you would just modify your own copy of it.*

Then as [we began using GLFW](/news/2014/yin-and-yang-using-glow-and-glfw.html), the package had to be completely rewritten (since before it was specifically for *Chippy*, our old window management package).

This left us asking questions like *"What should the package do? And better yet, not do?"*, we concluded:

* *"Keep it simple.."* - more complex applications can use window management API's directly.
* *"I just want a window to open so I can draw graphics to it."* - only provide what *most people use*.
* *"It should work everywhere!"* - this means mobile devices later on, too.

# Present

GLFW doesn't work on mobile devices -- and it's not entirely apparent if you would even want it to. Mobile devices offer their own superior API's which are often cleaner to use than respective desktop ones.

But this puts us in a very tricky situation: we get events through the GLFW API, and then through an Android or iOS API, so how do they interact? One *possible solution* is just to say that mobile and desktop applications are different, and therefore require different API's -- but this is becoming less true each day.

Version 2 of the package takes this all of this into consideration, and plans ahead for the future. Lets take a look at what the package gives us today:

**A complete application:**

```
import(
	"azul3d.org/gfx/window.v2"
	"azul3d.org/gfx.v1"
)

func gfxLoop(w window.Window, r gfx.Renderer) {
    // Initialization here.
    for {
        // Render loop here.
    }
}

func main() {
    window.Run(gfxLoop, nil)
}
```

In the above program `window.Run` opens a window and runs the `gfxLoop`. The second (and optional) parameter to `window.Run` is a set of properties to set on the window *before* it's been opened (or `nil`):

```
func main() {
	props := window.NewProps()
	props.SetTitle("My App {FPS}")
    window.Run(gfxLoop, props)
}
```

If the title has `{FPS}` in it -- the window title will automatically have it's `{FPS}` portion updated with the applications *frames per second*, neat!

You can request a window's properties be changed post-creation (and from any goroutine, too!):

```
props := w.Props()
props.SetSize(640, 480)
w.Request(props)
```

Events can be received from a `window.Window` via channels very easily:

```
// Create an events channel with sufficient buffer space.
events := make(chan window.Event, 256)

// Notify our channel anytime the cursor moves.
myWindow.Notify(events, window.CursorMovedEvents)

// Print as the cursor moves.
for event := range events{
	ev := event.(*window.CursorMoved)

	// Print the cursor position:
	fmt.Println(ev.X, ev.Y)
}
```

Keyboard and mouse state is tracked for you as well:

```
import(
	"azul3d.org/mouse.v1"
	"azul3d.org/keyboard.v1"
)

...

if myWindow.Mouse().Down(mouse.Left) {
	...
}
if myWindow.Keyboard().Down(keyboard.ArrowUp) {
	...
}
```

# Future

The [gfx/window](packages.html#gfx-window) package will add new event types for mobile devices (e.g. for touch events). You'll use the same API as you would for desktop applications, except you'll have touch events rather than mouse events, etc.

Support for fullscreen switching will be added (see [issue 2](https://github.com/azul3d/gfx-window/issues/2)), so that you toggle fullscreen *after a window has been opened*:

```
props := myWindow.Props()
props.SetFullscreen(!props.Fullscreen()) // Toggle fullscreen
myWindow.Request(props)
```

The same exact concept will also be applied to choosing settings like MSAA values (see [issue 4](https://github.com/azul3d/gfx-window/issues/4)). And perhaps we will add support for Joysticks (and onscreen ones on mobile devices?) as well (see [issue 5](https://github.com/azul3d/gfx-window/issues/5))!

There will *always* be exceptionally complex cases where using an abstraction layer such as this is not acceptable. For these cases *you can always use GLFW, or any other window management API directly (all you need is an OpenGL context)*.

We think that today -- the [gfx/window](packages.html#gfx-window) package provides an *easy, simple, and minimalistic* way to get an Azul3D application off the ground. And in the future, we hope that the package will help bridge cross-device application development in a simplistic mannor.

