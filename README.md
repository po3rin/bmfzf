# bmfzf

<img src="https://img.shields.io/badge/go-v1.14-blue.svg"/> [![GoDoc](https://godoc.org/github.com/po3rin/bmfzf?status.svg)](https://godoc.org/github.com/po3rin/bmfzf) ![Go Test](https://github.com/po3rin/bmfzf/workflows/Go%20Status/badge.svg) ![release](https://github.com/po3rin/bmfzf/workflows/release/badge.svg)

bmfzf lets you fuzzy search of Chrome Bookmarks.

<img src="./out.gif" width="560px">

## Install

### Go

```bash
$ go get -u github.com/po3rin/bmfzf
```

## Usage

bmfzf cli return bookmark url.

```bash
# choose Google bookmark
$ bmfzf
https://www.google.co.jp
```

When combined with other tools, bmfzf is more useful! following command is fuzzy serch & open page in Chrome.

```bash
# MacOS example
open -a '/Applications/Google Chrome.app' $(chrbmfzf)
```

## TODO

- [ ] Provide tools with Homebrew

