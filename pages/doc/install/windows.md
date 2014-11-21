# Windows Installation

What follows is a detailed step-by-step guide to install all Azul3D dependencies on a Windows (386/amd64) platform. For other platforms please see the [installation page](/doc/install).

* [Install Go](#install-go)
* [Install Git](#install-git)
* [Install MinGW](#install-mingw)
* [Install Azul3D](#install-azul3d)

# Install Go

Install the latest version of Go from [here](http://golang.org/doc/install).

Note: *Azul3D requires at least Go version 1.3*

Set the environment variable `GOPATH`, for example to `C:\Users\joe\Desktop\godev`.

# Install Git

Install git (use default installation options) from [here](http://git-scm.com/downloads).

Add git to the PATH environment variable:

* *32-bit:* `C:\Program Files\Git\bin`
* *64-bit:* `C:\Program Files (x86)\Git\bin`

# Install MinGW

We highly reccomend using TDM GCC -- installation is very simple, just download the installer and use the default options:

* [32-bit Installer](http://sourceforge.net/projects/tdm-gcc/files/TDM-GCC%20Installer/tdm-gcc-4.8.1-3.exe/download)
* [64-bit Installer](http://sourceforge.net/projects/tdm-gcc/files/TDM-GCC%20Installer/tdm64-gcc-4.8.1-3.exe/download)

Note: *Due to [an existing bug](https://github.com/go-gl/glfw3/issues/91) we reccomend only using GCC 4.8.1 or lower.*

# Install Azul3D

At this point, *you've successfully installed the dependencies*!

Follow the rest of the instructions on the [installation page](/doc/install).

