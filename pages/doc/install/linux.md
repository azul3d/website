# Linux Installation

What follows is a detailed step-by-step guide to install all Azul3D dependencies on a Linux (386/amd64) platform. For other platforms please see the [installation page](/doc/install).

* [Install Go](#install-go)
* [Dependencies](#dependencies)
* [Ubuntu 14.04](#ubuntu-1404)
* [Ubuntu 13.10](#ubuntu-1310)
* [Arch Linux](#arch-linux)
* [Install Azul3D](#install-azul3d)

# Install Go

Install the latest version of Go from [here](http://golang.org/doc/install).

Note: *Azul3D requires at least Go version 1.3*

Set the environment variable `GOPATH`, for example to `/home/joe/godev`.

# Dependencies

The following are general dependencies that Azul3D requires on Linux, single-line installers for common Linux distributions are provided below.

If your distribution is not listed here please help us by [creating an issue](http://github.com/azul3d/issues/issues).

* Go 1.3
* git (to clone the source code).
* gcc
* OpenGL development files
* libx11 development files
* libx11-xcb development files
* libxcb development files (including icccm, image, randr, render-util, and xcb-xkb parts)
* libfreetype, libbz2 and libzip static libraries and headers

# Ubuntu 14.04

Using apt-get on Ubuntu 14.04 you can install all of the dependencies by running:

```
sudo apt-get install build-essential git mesa-common-dev libx11-dev libx11-xcb-dev libxcb-icccm4-dev libxcb-image0-dev libxcb-randr0-dev libxcb-render-util0-dev libxcb-xkb-dev libfreetype6-dev libbz2-dev libxxf86vm-dev libgl1-mesa-dev
```

# Ubuntu 13.10

With Ubuntu 13.10, you need to add the xorg-edgers PPA for the libxkb-xcb-dev package:

```
sudo add-apt-repository ppa:xorg-edgers/ppa
sudo apt-get update
```

Then you can use apt-get to install everything:

```
sudo apt-get install build-essential git mesa-common-dev libx11-dev libx11-xcb-dev libxcb-icccm4-dev libxcb-image0-dev libxcb-randr0-dev libxcb-render-util0-dev libxcb-xkb-dev libfreetype6-dev libbz2-dev
```

# Arch Linux

Using pacman on Arch Linux you can install all of the dependencies by running:

```
pacman -Sy base-devel git mesa libx11 libxcb xcb-util-wm xcb-util-image libxrandr xcb-util-renderutil libxkbcommon-x11 freetype2 bzip2
```

# Install Azul3D

At this point, *you've successfully installed the dependencies*!

Follow the rest of the instructions on the [installation page](/doc/install).
