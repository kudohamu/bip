# bip

Wrapper of `pongo2.TemplateSet` for `go-bindata`

[![GoDoc](https://godoc.org/github.com/kudohamu/bip?status.svg)](https://godoc.org/github.com/kudohamu/bip)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Motivation

When use `go-bindata` and `pongo2` in combination, it is troublesome to build every time rewrite views (**even using Must**).
So `bip` behaves what read from files every executions when in development, and read from bindata when in production.

## Usage

```go
package main

import (
  ...
)

var tplSet *bip.TemplateSet
var indexTpl bip.Template

func init() {
  tplSet = bip.NewSet(asset.Asset) // asset is a package generated from bindata.
  indexTpl = bip.Must(tplSet.FromFile("index.html"))
}

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    if err := indexTpl.ExecuteWriter(bip.Context{}, w); err != nil {
      fmt.Fprintf(w, err.Error())
    }
  })
  http.ListenAndServe(":8080", nil)
}
```

if you use as production binary, you need to add `bip` to build tag.

```shell
$ go build -tags="bip ...other tags"
```

## Caution

`bip` depends on **master** of `pongo2`. **DO NOT** use v3.
