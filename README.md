# bmfzf

<img src="https://img.shields.io/badge/go-v1.14-blue.svg"/> [![GoDoc](https://godoc.org/github.com/po3rin/bmfzf?status.svg)](https://godoc.org/github.com/po3rin/bmfzf) ![Go Test](https://github.com/po3rin/bmfzf/workflows/Go%20Status/badge.svg) ![release](https://github.com/po3rin/bmfzf/workflows/release/badge.svg)

bmfzf lets you fuzzy search of Chrome Bookmarks.

<img src="./out.gif" width="640px">

## Install

#### ■ curl

you can install with curl.

```
$ curl -sf https://gobinaries.com/po3rin/bmfzf | sh
```

Go Binaries is an on-demand binary server, allowing non-Go users to quickly install tools
https://github.com/tj/gobinaries

#### ■ Release Page

https://github.com/po3rin/bmfzf/releases

#### ■ Go

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
$ open -a '/Applications/Google Chrome.app' $(bmfzf)
```

If you changed ```Bookmarks``` file location, please use -f option. 

About Bookmarks  
https://www.ubergizmo.com/how-to/find-google-chrome-bookmarks-computer/



## TODO

- [ ] Provide tools with Homebrew

