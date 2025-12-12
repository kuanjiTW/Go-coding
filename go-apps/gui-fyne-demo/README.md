
Introduction
============

The files are adapted from [fyne_demo](https://github.com/fyne-io/demo).


How to Run
==========


Windows
-------

A compiler is required to install a dependency of Fyne Demo.
Assume [WSL2](https://learn.microsoft.com/en-us/windows/wsl/install)
has already installed on Windows.

Run the following commands to install required system packages.

```
$ sudo apt update
$ sudo apt install golang gcc libgl1-mesa-dev xorg-dev libxcursor-dev libxinerama-dev libxrandr-dev libxi-dev
```

Then run the following commands to install Fyne Demo
and execute it.

```
$ go install fyne.io/fyne/v2/cmd/fyne_demo@latest
$ ~/go/bin/fyne_demo
```


Linux/macOS
-----------

To run the demo application, execute the following commands.

```
$ go install fyne.io/fyne/v2/cmd/fyne_demo@latest 
$ fyne_demo
```

`fyne_demo` may fail in macOS. In this case, install the Fyne Demo
Mac app with the following commands.

```
$ go install fyne.io/tools/cmd/fyne@latest
$ fyne install fyne.io/demo
```

Then a Fyne Demo app is ready for execution.
