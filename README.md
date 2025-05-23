# go-isatty

[![Godoc Reference](https://godoc.org/github.com/matt-FFFFFF/go-isatty?status.svg)](http://godoc.org/github.com/matt-FFFFFF/go-isatty)
[![Codecov](https://codecov.io/gh/matt-FFFFFF/go-isatty/branch/master/graph/badge.svg)](https://codecov.io/gh/matt-FFFFFF/go-isatty)
[![Coverage Status](https://coveralls.io/repos/github/matt-FFFFFF/go-isatty/badge.svg?branch=master)](https://coveralls.io/github/matt-FFFFFF/go-isatty?branch=master)
[![Go Report Card](https://goreportcard.com/badge/matt-FFFFFF/go-isatty)](https://goreportcard.com/report/matt-FFFFFF/go-isatty)

isatty for golang

## Usage

```go
package main

import (
 "fmt"
 "github.com/matt-FFFFFF/go-isatty"
 "os"
)

func main() {
 if isatty.IsTerminal(os.Stdout.Fd()) {
  fmt.Println("Is Terminal")
 } else if isatty.IsCygwinTerminal(os.Stdout.Fd()) {
  fmt.Println("Is Cygwin/MSYS2 Terminal")
 } else {
  fmt.Println("Is Not Terminal")
 }
}
```

## Installation

```
go get github.com/matt-FFFFFF/go-isatty
```

## Thanks

* This fork is based on the original and is used with thanks [mattn/go-isatty](https://github.com/mattn/go-isatty)
* k-takata: base idea for IsCygwinTerminal

    <https://github.com/k-takata/go-iscygpty>
