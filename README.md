# archiveis

[![Documentation](https://godoc.org/github.com/jaytaylor/archive.is?status.svg)](https://godoc.org/github.com/jaytaylor/archive.is)
[![Build Status](https://travis-ci.org/jaytaylor/archive.is.svg?branch=master)](https://travis-ci.org/jaytaylor/archiveis)
[![Report Card](https://goreportcard.com/badge/github.com/jaytaylor/archive.is)](https://goreportcard.com/report/github.com/jaytaylor/archive.is)

### About

archive.is is a golang package for archiving web pages via [archive.is](https://archive.is).

Please be mindful and responsible and go easy on them, we want archive.is to last forever!

Created by [Jay Taylor](https://jaytaylor.com/).

### Requirements

* Go version 1.9 or newer

### Installation

```bash
go get github.com/jaytaylor.com/archive.is/...
```

### Usage

```go
package main

import (
	"fmt"

	"github.com/jaytaylor/archive.is"
)

var inputURL = "https://jaytaylor.com/"

func main() {
	archiveURL, err := archiveis.Capture(inputURL)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully archived %v via archive.is: %v\n", inputURL, archiveURL)
}

// Output:
//
// Successfully archived https://jaytaylor.com/ via archive.is: https://archive.is/i2PiW
```

### Running the test suite

    go test ./...

#### License

Permissive MIT license, see the [LICENSE](LICENSE) file for more information.
