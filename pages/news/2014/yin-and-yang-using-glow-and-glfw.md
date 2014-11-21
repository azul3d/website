# Yin and Yang: Using Glow and GLFW
<p class="date">August 30, 2014</p>

Azul3D's *glwrap* command and *[Glow](https://github.com/go-gl/glow)* are both OpenGL wrapper generators for Go, our *Chippy* library and *[GLFW 3](https://github.com/glfw/glfw)* are OpenGL window management libraries both usable from Go.

They are both yin and yang -- each had features and issues that the other did not.

After a word of cautious advice from *[Dmitri Shuralyov](https://github.com/azul3d/issues/issues/5#issuecomment-51740222)*, who said:

> "We have the opportunity to do better. Both go-gl repos and Azul3D repos are open source and anyone can contribute to make them better. Assuming two packages have an identical purpose, wouldn't it be nicer to have just one place to contribute to instead?"

And then, everything just clicked: *"Why aren't we doing that today?"*

# Glow and glwrap

glwrap, our old OpenGL wrapper generator did not generate code that was buildable on OS X. The fix for this would have been only a few lines of code -- but instead we've dedicated ourselves towards working with the rest of the Go community to make Glow -- an open source OpenGL wrapper generator for Go -- better.

Starting out, this just meant a few minor changes like [fixing typos](https://github.com/go-gl/glow/pull/38) and [resolving template files](https://github.com/go-gl/glow/pull/36) outside the source directory so you can run Glow from any directory. Nothing fancy, just minor changes.

Ultimately these minor changes led to us making some larger feature additions to Glow, like a [-restrict](https://github.com/go-gl/glow/pull/43) flag that allows you to specify a JSON file of OpenGL symbols that you want the generated code restricted to. *This lowers the file size and compiler time of the generated code significantly*

And then a [later pull request](https://github.com/go-gl/glow/pull/41) to clean up the package synopsis of generated code led to a nice amount of collaboration thus making the code a bit cleaner -- which is suprising because of how clean Glow's code is already!

The latest version of our OpenGL renderer ([azul3d.org/gfx/gl2.v2](http://azul3d.org/gfx/gl2.v2)) features the use of Glow-OpenGL wrappers internally -- this is amongst a few other [minor changes.](https://github.com/azul3d/gfx-gl2#version-2) to the renderer.

# GLFW3 and Chippy

Chippy -- an OpenGL window management package for Go, was actually the very first Azul3D package(!), and was created before GLFW 3 was even in development. Because of this it's not suprising how outdated it is. Although it generally shares the exact same features as GLFW 3 today -- it doesn't support OS X and has no Wayland client on Linux, and it's generally more buggy.

*Chippy is no longer supported*. The code will stay where it is so older applications will continue to work as-is, but why play catch-up trying to make Chippy as good as GLFW 3 *already is today*? We can devote our efforts to making GLFW 3 an even better library for everyone.

In fact, we've already made improvements to the Go bindings to GLFW 3 by [making them go-gettable](https://github.com/go-gl/glfw3/pull/83) without the installation of GLFW 3 beforehand -- as well as finding and identifying a few [deadlocks](https://github.com/go-gl/glfw3/issues/8#issuecomment-53814911).

A new version of the Azul3D window package (track progress [here](https://github.com/azul3d/issues/issues/20)) -- `gfx/window` ([azul3d.org/gfx/window.v2](http://azul3d.org/gfx/window.v2)) will operate as a small abstraction layer over GLFW providing some much-needed features like easy fullscreen switching.

In the future the `gfx/window` package will allow you to develop cross-device (desktop and mobile) applications with ease -- ignoring all of those tricky corner-cases.

# Conclusion

In conclusion, instead of independent projects each with different features (and likewise equally lacking), we now have one project for OpenGL wrapper generation in Go, and one project for OpenGL window creation in Go, and the both of them are more clean and robust than ever before because of it.

All of this brings the future of having a truly cross-platform game engine written in Go closer to us than ever before, *and we are extremely excited about it!*
