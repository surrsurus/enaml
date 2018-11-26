# enaml [![Build Status](https://travis-ci.org/surrsurus/enaml.svg?branch=master)](https://travis-ci.org/surrsurus/enaml) [![Go Report Card](https://goreportcard.com/badge/github.com/surrsurus/enaml)](https://goreportcard.com/report/github.com/surrsurus/enaml) ![Golang Version](https://img.shields.io/badge/golang-v1.9-green.svg) [![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0) 

**E**naml is **N**ot **A** **M**arkup **L**anguage

<!-- <img align="center" src="https://github.com/surrsurus/enaml/blob/master/media/logo.png" alt="enaml" width=250> -->

Enaml is a markup renderer, that uses a simplified markup language similar to GFM and can convert it to HTML files for easy viewing later. See the examples under `examples` for more.

Enaml was originally made for a user to take efficient notes quickly using this simplified markup syntax and be able to easiy render them en masse to HTML. 

## Getting Started

Here's how to set up enaml on your computer.

### Prerequisites

1. Install [go-1.9](https://golang.org/dl/) on your OS of choice and make sure the `go` executable is in your PATH.

2. Download the latest master, and extract the zip or tarball

### Building

#### Windows

3. Run 

```
$ go build
``` 

in the root directory of enaml.

4. You should now see an `enaml.exe` in the directory.

#### Linux/OSX

3. Run 

```
$ go build
``` 

in the root directory of enaml.

4. You should now see an `enaml` binary in the directory.

### Running the Tests

Run the full test suite, benchmarks, and check code coverage by running either `win-test.bat` on windows or `linux-run.sh` on linux in the root directory.

In addition, you could also run `go test` in the root enaml directory.

### Examples

The `examples` folder has many examples. The file `syntax by example.enaml` details the whole syntax for the enaml-interpretable markup language. By running 

```
$ enaml.exe "examples\syntax by example.enaml"
```

 on windows or 
 
 ```
 $ enaml "examples/syntax by example.enaml"
 ``` 
 
 on linux/osx, you will generate a rendered enaml file as html in the `examples` directory.

## License

<img align="center" src="https://licensebuttons.net/l/GPL/2.0/88x62.png" alt="GPL" width=100>

This code is released under the GNU GENERAL PUBLIC LICENSE. All works in this repository are meant to be utilized under this license. You are entitled to remix, remodify, and redistribute this program as you see fit, under the condition that all derivative works must use the GPL Version 3.

## Acknowledgments

A big thank you to the [Github Markdown CSS](https://github.com/sindresorhus/github-markdown-css) project. The CSS is really nice.
