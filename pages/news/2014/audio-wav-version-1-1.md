# audio/wav: version 1.1
<p class="date">October 18, 2014</p>

The [audio/wav](/packages.html#audio-wav) package is a *WAV encoder and decoder*. It allows you to decode many different types of WAV files into raw PCM audio data.

It operates through the [audio](/packages.html#audio) package, which is like the standard image package -- but for audio.

# Version 1.1

Version 1.1 of the [audio/wav](/audio/wav.v1) package is now available!

 * `import "azul3d.org/audio/wav.v1"`
 * [![GoDoc](https://godoc.org/azul3d.org/audio/wav.v1?status.svg)](https://godoc.org/azul3d.org/audio/wav.v1)

# Overview

 * Encoding support (was issue [#1](https://github.com/azul3d/audio-wav/issues/1)).
  * *You can decode and encode WAV audio in Go (thanks goes out to [Robin Eklind](https://github.com/mewmew) and [Henry Eklind](https://github.com/karlek)!).*
 * Massive (~100%) performance improvements (was issue [#8](https://github.com/azul3d/audio-wav/issues/8)).
  * *The decoder is much faster averaging around 5.5ms to decode 1s of 44100hz/2ch WAV audio (On a Pentium(R) Dual-Core CPU - T4500 @ 2.30GHz), which is on-par with the WAV encoder.*
 * Removed the unsafe dependency (was issue [#4](https://github.com/azul3d/audio-wav/issues/4)).
  * *The entire package is now pure Go.*
 * Tests and benchmarks have been added.
  * *Which lead to the identification of a now-fixed, [critical bug](https://github.com/azul3d/audio-wav/issues/12).*
 * [Full Changelog](https://github.com/azul3d/audio-wav/compare/v1...v1.1)

# Updating

To update you need to run the following (*note that the import path is the same*):

```
go get -u azul3d.org/audio/wav.v1
```

# Future Audio Work

 * More WAV performance enhancements (see issue [#8](https://github.com/azul3d/audio-wav/issues/8)).
 * A FLAC decoder ([currently in-development](https://github.com/azul3d/audio-flac)).
 * Idiomatic audio playback in Go, *with support for spatial 3D audio* (see issue [#2](https://github.com/azul3d/issues/issues/2)).
