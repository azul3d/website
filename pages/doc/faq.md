# FAQ

Compiled here is a list of frequently asked questions about Azul3D. If you have questions not listed here please ask in the [Mailing List](https://groups.google.com/forum/#!forum/azul3d).

* [What is Azul3D?](#what-is-azul3d)
* [Why Go?](#why-go)
* [What does it provide?](#what-does-it-provide)
* [Is there a roadmap?](#is-there-a-roadmap)
* [Are there any examples?](#are-there-any-examples)
* [How do I install it?](#how-do-i-install-it)
* [How does updating work?](#how-does-updating-work)
* [How does versioning work?](#how-does-versioning-work)
* [Why OpenGL 2?](#why-opengl-2)
* [What about the garbage collector?](#what-about-the-garbage-collector)
* [Does it have a level editor?](#does-it-have-a-level-editor)
* [Can it be scripted?](#can-it-be-scripted)
* [Where can I report bugs?](#where-can-i-report-bugs)
* [Where can I request features?](#where-can-i-request-features)
* [How can I get in touch?](#how-can-i-get-in-touch)
* [What types of applications?](#what-types-of-applications)
* [What platforms are supported?](#what-platforms-are-supported)


# What is Azul3D?

Azul3D is a 3D game engine written entirely from the ground up in the Go programming language. It is suitable for most types of 2D and 3D applications, it can even be used for non-game interactive applications. Unlike other modern game engines like Unity, JMonkey, etc, Azul3D does not attempt to provide anything more than programming packages suitable for game developers. There is no custom IDE, nor level editor.

It is a game engine for programmers, by programmers. It is minimalistic, but also extensable compared to most other engines.

It tries to provide the things that most game developers will use daily -- and it tries to provide them well.

# Why Go?

Go is one of the very few languages in existence today that provides concurrency at a language level, additionally, Go is compiled into native binaries and has the potential to reach C/C++ speeds. Only some of the most advanced open source game engines today provide good support for multi-threaded rendering and, because of Go, Azul3D is able to provide that feature with ease and simplicity.

There are many other areas where Go could excel at game development -- Go has syntax familiar to those coming from scripting languages and ideally games written in Go would also be 'scripted' in Go. You can also imagine each computer player in a simulated game being controlled by it's own goroutine -- something that would not be possible with the threading model present in other languages.


# What does it provide?

Azul3D provides a set of packages for developing games and other interactive 2D or 3D applications in Go, a list of those packages and their documentation can be found [here](/packages.html).

# Is there a roadmap?

Yes, please see the [roadmap page](/doc/roadmap.html).

# Are there any examples?

Yes! You can download and compile them by following the instructions found on the <a href="/doc/install">installation page.
Alternatively, [view the source on GitHub](https://github.com).

# How do I install it?

Please see the [installation page](/doc/install).

# How does updating work?

Please see the [updating page](/doc/updating.html).

# How does versioning work?

Please see the [versioning page](/doc/versioning.html).

# Why OpenGL 2?

A less known fact about OpenGL is that it is *extremely forward compatible*. *In fact all most everything you can do in OpenGL 3/4 is also readily available in OpenGL 2 via standardized extensions*. Although Azul3D only speaks to your GPU through the OpenGL 2 API; It uses the most modern OpenGL functions available.

This is actually benificial in many ways, one being that Azul3D can target older hardware that *mostly* supports OpenGL 3 but doesn't advertise it. When Azul3D does begin using an OpenGL 3 renderer the change will be completely seamless thanks to our [gfx package](/packages.html#gfx), meaning you won't have to change a thing.

**We get all the same features, the same performance, and we are able to access a wider range of hardware.**

# What about the garbage collector?

Go is by specification a *garbage collected programming language*, this means that developers do not need to manually control memory but also means they must be aware of how the garbage collector works in order to receive exceptional performance.

The default implementation of Go, uses a *stop-the-world garbage collector*. This means application code is paused while memory is being garbage collected. If this happens often, your graphical application could appear to be slow or laggy.

But Go is different from other garbage collected languages because it gives you much more control over garbage collection and memory management in general. In truth garbage collection does not hinder game development as much as people generally think because of the following reasons:

* OpenGL uses double buffering, when a frame is submitted to OpenGL it doesn't actually show up until the buffer swapping occurs.
* Azul3D uses a multi-threaded renderer, providing yet-another level of buffering (application logic and the renderer operate independently).
* GC pause times are often small, and much of the time negligible. If a developer pays special care to not make lots of garbage, all the better.
* Go 1.3 introduced [sync.Pool](http://www.golang.org/pkg/sync/#Pool), a garbage collector aware free list which allows developers to alleviate GC pressure by re-using old already allocated objects.
* The goal of the Go 1.5 (June 2015) Garbage Collector (GC) is to reduce GC latency, making Go acceptable for implementing a broad spectrum of systems requiring low response times. For further details please refer to the following design document: [Go 1.4+ Garbage Collection (GC) Plan and Roadmap.](http://golang.org/s/go14gc).

# Does it have a level editor?

Not right now.

Azul3D is mainly for programmers and doesn't yet offer any GUI-based tools. With that being said, there may be importers in the future for 3D modelers which will make importing scenes from those 3D modelers (such as Blender) into Azul3D trivial.

# Can it be scripted?

One of the key points of making games in Go is the easy syntax. You probably don't actually need a scripting language at all. With that being said, you could probably use some scripting language that can interact with Go to script games, yes.

# Where can I report bugs?

We appreciate that bugs specific to a single package be reported on that package's GitHub repository.

For bugs related to this website, etc, please file them in the [generic issue tracker](https://github.com/azul3d/issues/issues).

# Where can I request features?

We appreciate that feature requests specific to a single package be requested on that package's GitHub issue tracker.

For feature requests not related to any package, or related to this website, etc, please request them by creating an issue for them in the [generic issue tracker](https://github.com/azul3d/issues/issues).

# How can I get in touch?

Please make a post on the [Mailing List](https://groups.google.com/forum/#!forum/azul3d).

# What types of applications?

Any 2D or 3D graphical application, but mainly games.

# What platforms are supported?

Please see the [platform support](/doc/platform-support.html) page.

