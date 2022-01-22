# go-flags-demo
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-1-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/go-easygen/go-flags-demo?status.svg)](http://godoc.org/github.com/go-easygen/go-flags-demo)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-easygen/go-flags-demo)](https://goreportcard.com/report/github.com/go-easygen/go-flags-demo)
[![Build Status](https://github.com/go-easygen/go-flags-demo/actions/workflows/go-release-build.yml/badge.svg?branch=master)](https://github.com/go-easygen/go-flags-demo/actions/workflows/go-release-build.yml)
[![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-B.svg)](http://godoc.org/github.com/go-easygen/wireframe)


redo - global option redo

Redo global option via automatic code-gen


## TOC
- [go-flags-demo - global option redo](#go-flags-demo---global-option-redo)
  - [Synopsis](#synopsis)
- [Usage](#usage)
- [Development History](#development-history)
  - [Demo](#demo)
  - [Update 3](#update-3)
  - [Update on 2022-01-22](#update-on-2022-01-22)
- [Download/install binaries](#downloadinstall-binaries)
  - [The binary executables](#the-binary-executables)
  - [Distro package](#distro-package)
  - [Debian package](#debian-package)
- [Install Source](#install-source)
- [Author](#author)

## go-flags-demo - global option redo

`go-flags-demo` will redo global option via automatic code-gen

```sh
$ go-flags-demo -V || true
redo - global option redo
Copyright (C) 2022, Myself <me@mine.org>

Redo global option via automatic code-gen

Built on 2022-01-22
Version 0.1.0
```

### Synopsis

Check out [`easygen` cli code-gen example for go-flags](https://github.com/go-easygen/easygen/issues/46) on how these Go code were ***automatically** generated*.


## Usage


```sh
$ go-flags-demo || true
Please specify one command of: build, install or publish

Usage:
  go-flags-demo [OPTIONS] <build | install | publish>

Application Options:
  -H, --host=    Host address (default: localhost) [$REDO_HOST]
  -p, --port=    Listening port (default: 80) [$REDO_PORT]
  -f, --force    Force start [$REDO_FORCE]
  -v, --verbose  Verbose mode (Multiple -v options increase the verbosity)
  -V, --version  Show program version and exit

Help Options:
  -h, --help     Show this help message

Available commands:
  build    Build the network application
  install  Install the network application
  publish  Publish the network application
```

## Development History

The following are the demos while the [`wireframed`](https://github.com/suntong/lang/tree/master/lang/Go/src/sys/go-flags/wireframed) was developed (thus having the executable name of `wireframed`):

### Demo

``` sh
$ wireframed
Please specify one command of: build, install or publish

Usage:
  wireframed [OPTIONS] <build | install | publish>

Application Options:
  -H, --host=    host address (default: localhost) [$REDO_HOST]
  -p, --port=    listening port (default: 80) [$REDO_PORT]
  -f, --force    force start [$REDO_FORCE]
  -v, --verbose  Verbose mode (Multiple -v options increase the verbosity)

Help Options:
  -h, --help     Show this help message

Available commands:
  build    Build the network application
  install  Install the network application
  publish  Publish the network application



$ wireframed publish
the required flag `-d, --dir' was not specified

Usage:
  wireframed [OPTIONS] publish [publish-OPTIONS] ID Num Rest...

Publish the built network application to central repo

Application Options:
  -H, --host=       host address (default: localhost) [$REDO_HOST]
  -p, --port=       listening port (default: 80) [$REDO_PORT]
  -f, --force       force start [$REDO_FORCE]
  -v, --verbose     Verbose mode (Multiple -v options increase the verbosity)

Help Options:
  -h, --help        Show this help message

[publish command options]
      -d, --dir=    publish dir
          --suffix= source file suffix (default: .go,.c,.s)
      -o, --out=    output filename



$ wireframed publish -d .
the required arguments `ID` and `Num` were not provided

Usage:
  wireframed [OPTIONS] publish [publish-OPTIONS] ID Num Rest...

Publish the built network application to central repo

Application Options:
  -H, --host=       host address (default: localhost) [$REDO_HOST]
  -p, --port=       listening port (default: 80) [$REDO_PORT]
  -f, --force       force start [$REDO_FORCE]
  -v, --verbose     Verbose mode (Multiple -v options increase the verbosity)

Help Options:
  -h, --help        Show this help message

[publish command options]
      -d, --dir=    publish dir
          --suffix= source file suffix (default: .go,.c,.s)
      -o, --out=    output filename



$ wireframed publish -d ../dist v1 123 abc def
Publish the network application
Copyright (C) 2022, Myself <me@mine.org>

Doing Publish, with []string{}
../dist .go,.c,.s  {v1 123 [abc def]}

```

The last two lines are output from:


``` go
	fmt.Printf("Doing Publish, with %#v\n", args)
	fmt.Println(x.Dir, x.Suffix, x.Out, x.Args)
```


``` sh
$ wireframed install --dir ../new abc def
Install the network application
Copyright (C) 2022, Myself <me@mine.org>

Doing Install, with []string{"abc", "def"}
../new .go,.c,.s
```

### Update 1

After [ouput global option](https://github.com/suntong/lang/commit/5ab110dfde2c2a481e14f8cb8afa4d4b3cb4bd23#diff-4040661b6ba228abb2484408a7a62935f7660d89545b11fc65567f54d36d7501),

``` sh
$ wireframed publish -d ../dist v1 123 abc def
Publish the network application
Copyright (C) 2022, Myself <me@mine.org>

Doing Publish, with {Host:localhost Port:80 Force:false Verbflg:0x4a8120 Verbose:0}, []
../dist .go,.c,.s  {v1 123 [abc def]}



$ REDO_HOST=myserver REDO_PORT=8080 wireframed publish -v -d ../dist v1 123 abc def
Publish the network application
Copyright (C) 2022, Myself <me@mine.org>

Doing Publish, with {Host:myserver Port:8080 Force:false Verbflg:0x4a8120 Verbose:1}, []
../dist .go,.c,.s  {v1 123 [abc def]}



$ REDO_HOST=myserver REDO_PORT=8080 wireframed publish -vv -H newserver --port 8888 --force -d ../dist v1 123 abc def
Publish the network application
Copyright (C) 2022, Myself <me@mine.org>

Doing Publish, with {Host:newserver Port:8888 Force:true Verbflg:0x4a8120 Verbose:2}, []
../dist .go,.c,.s  {v1 123 [abc def]}
```

### Update 2

After turning `--suffix` to use choices.

``` sh
$ wireframed publish
the required flag `-d, --dir' was not specified

Usage:
  wireframed [OPTIONS] publish [publish-OPTIONS] ID Num Rest...

Publish the built network application to central repo

Application Options:
  -H, --host=                  host address (default: localhost) [$REDO_HOST]
  -p, --port=                  listening port (default: 80) [$REDO_PORT]
  -f, --force                  force start [$REDO_FORCE]
  -v, --verbose                Verbose mode (Multiple -v options increase the verbosity)

Help Options:
  -h, --help                   Show this help message

[publish command options]
      -d, --dir=               publish dir
      -s, --suffix=[.go|.c|.h] source file suffix for publish
      -o, --out=               output filename

# see that the --suffix= now use choices



$ wireframed publish -d ../dist v1 123 abc def
Publish the network application
Copyright (C) 2022, Myself <me@mine.org>

Doing Publish, with {Host:localhost Port:80 Force:false Verbflg:0x4a8180 Verbose:0}, []
../dist []  {v1 123 [abc def]}

# the --suffix= has empty choices by default



$ wireframed publish -d ../dist -s .c -s .h v1 123 abc def
Publish the network application
Copyright (C) 2022, Myself <me@mine.org>

Doing Publish, with {Host:localhost Port:80 Force:false Verbflg:0x4a8180 Verbose:0}, []
../dist [.c .h]  {v1 123 [abc def]}

# the --suffix= has two choices now



$ wireframed publish -d ../dist -s .x v1 123 abc def
Invalid value `.x' for option `-s, --suffix'. Allowed values are: .go, .c or .h

Usage:
  wireframed . . .

# when the --suffix= has been provided with wrong choice

```

### Update 3

After switching to `go-easygen` which has `clis` support functions,

``` sh
$ wireframed install -v
Install the network application
Copyright (C) 2022, Myself <me@mine.org>

[redo::install] Doing Install, with {Host:localhost Port:80 Force:false Verbflg:0x4a92e0 Verbose:1}, []

./ .go,.c,.s
[redo::install] Warning: Install, Exec, Sample warning: Instance not found
```

- The `clis.Verbose(1,` will only output if the Verbose level is >=1. 
- Removing the `-v` option from the command line the "Doing Install" line will disappear.
- With `clis.WarnOn` & `clis.AbortOn` reporting warning or critical errors (in color) will be a breeze.

### Update on 2022-01-22

After update to new template, which uses `-V, --version` to show program version,

``` sh
$ wireframed
Please specify one command of: build, install or publish

Usage:
  wireframed [OPTIONS] <build | install | publish>

Application Options:
  -H, --host=    Host address (default: localhost) [$REDO_HOST]
  -p, --port=    Listening port (default: 80) [$REDO_PORT]
  -f, --force    Force start [$REDO_FORCE]
  -v, --verbose  Verbose mode (Multiple -v options increase the verbosity)
  -V, --version  Show program version and exit

Help Options:
  -h, --help     Show this help message

Available commands:
  build    Build the network application
  install  Install the network application
  publish  Publish the network application

$ wireframed -V
redo - global option redo
Copyright (C) 2022, Myself <me@mine.org>

Redo global option via automatic code-gen

Built on 2022-01-22
Version 0.1.0
```

## Download/install binaries

- The latest binary executables are available 
as the result of the Continuous-Integration (CI) process.
- I.e., they are built automatically right from the source code at every git release by [GitHub Actions](https://docs.github.com/en/actions).
- There are two ways to get/install such binary executables
  * Using the **binary executables** directly, or
  * Using **packages** for your distro

### The binary executables

- The latest binary executables are directly available under  
https://github.com/go-easygen/go-flags-demo/releases/latest 
- Pick & choose the one that suits your OS and its architecture. E.g., for Linux, it would be the `go-flags-demo_verxx_linux_amd64.tar.gz` file. 
- Available OS for binary executables are
  * Linux
  * Mac OS (darwin)
  * Windows
- If your OS and its architecture is not available in the download list, please let me know and I'll add it.
- The manual installation is just to unpack it and move/copy the binary executable to somewhere in `PATH`. For example,

``` sh
tar -xvf go-flags-demo_*_linux_amd64.tar.gz
sudo mv -v go-flags-demo_*_linux_amd64/go-flags-demo /usr/local/bin/
rmdir -v go-flags-demo_*_linux_amd64
```


### Distro package

- [Packages available for Linux distros](https://cloudsmith.io/~suntong/repos/repo/packages/) are
  * [Alpine Linux](https://cloudsmith.io/~suntong/repos/repo/setup/#formats-alpine)
  * [Debian](https://cloudsmith.io/~suntong/repos/repo/setup/#formats-deb)
  * [RedHat](https://cloudsmith.io/~suntong/repos/repo/setup/#formats-rpm)

The repo setup instruction url has been given above.
For example, for [Debian](https://cloudsmith.io/~suntong/repos/repo/setup/#formats-deb) --

### Debian package


```sh
curl -1sLf \
  'https://dl.cloudsmith.io/public/suntong/repo/setup.deb.sh' \
  | sudo -E bash

# That's it. You then can do your normal operations, like

sudo apt-get update
apt-cache policy go-flags-demo

sudo apt-get install -y go-flags-demo
```

## Install Source

To install the source code instead:

```
go get -v -u github.com/go-easygen/go-flags-demo
```

## Author

Tong SUN  
![suntong from cpan.org](https://img.shields.io/badge/suntong-%40cpan.org-lightgrey.svg "suntong from cpan.org")

_Powered by_ [**WireFrame**](https://github.com/go-easygen/wireframe)  
[![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-Y.svg)](http://godoc.org/github.com/go-easygen/wireframe)  
the _one-stop wire-framing solution_ for Go cli based projects, from _init_ to _deploy_.
