# enaml
**E**naml is **N**ot **A** **M**arkup **L**anguage

[![Build Status](https://travis-ci.org/surrsurus/enaml.svg?branch=master)](https://travis-ci.org/surrsurus/enaml) ![Golang Version](https://img.shields.io/badge/golang-v1.9-green.svg) [![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0) 

<!-- <img align="center" src="https://github.com/surrsurus/enaml/blob/master/media/logo.png" alt="enaml" width=250> -->

## Usage

1. Install go-1.9 on your OS of choice and make sure the `go` executable is in your PATH
2. Run `go build`
3. Run `./enaml file.enaml` to generate a rendered html version. See the `examples` directory for some files to convert.

## What is enaml then?

Enaml is actually a markup renderer, that uses a simplified markup language similar to GFM. See the examples under `examples` for more. It was made for a user to take efficient notes quickly and be able to easiy render them to HTML for later viewing. 

## Why golang?

The old version of this project was written in scala and it was a nightmare to deploy. Go makes getting an executable so much easier.

## Testing

Run the full test suite, benchmarks, and check code coverage by running either `win-test.bat` on windows or `linux-run.sh` on linux in the root directory.

## Credit

A big thank you to the [Github Markdown CSS](https://github.com/sindresorhus/github-markdown-css) project. The CSS is really nice.