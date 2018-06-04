# archiveis

[![Documentation](https://godoc.org/github.com/jaytaylor/archive.is?status.svg)](https://godoc.org/github.com/jaytaylor/archive.is)
[![Build Status](https://travis-ci.org/jaytaylor/archive.is.svg?branch=master)](https://travis-ci.org/jaytaylor/archiveis)
[![Report Card](https://goreportcard.com/badge/github.com/jaytaylor/archive.is)](https://goreportcard.com/report/github.com/jaytaylor/archive.is)

### About

archive.is is a golang package for archiving web pages via [archive.is](https://archive.is).

Please be mindful and responsible and go easy on them, we want archive.is to last forever!

Created by [Jay Taylor](https://jaytaylor.com/).

Also see: [archive.org golang package](https://jaytaylor.com/archive.org)

### TODO

* Add timeout to `.Capture`.
* Consider unifying to single binary

### Requirements

* Go version 1.9 or newer

### Installation

```bash
go get jaytaylor.com/archive.is/...
```

### Usage

#### Command-line programs

##### `archive.is <url>`

Archive a fresh new copy of an HTML page

##### `archive.is-snapshots <url>`

Search for existing page snapshots

Search query examples:

* `microsoft.com` for snapshots from the host microsoft.com
* `*.microsoft.com` for snapshots from microsoft.com and all its subdomains (e.g. www.microsoft.com)
* `http://twitter.com/burgerking` for snapshots from exact url (search is case-sensitive)
* `http://twitter.com/burg*` for snapshots from urls starting with http://twitter.com/burg

#### Go package interfaces

##### Capture URL HTML Page Content

[capture.go](_examples/capture/capture.go):

```go
package main

import (
	"fmt"

	"github.com/jaytaylor/archive.is"
)

var captureURL = "https://jaytaylor.com/"

func main() {
	archiveURL, err := archiveis.Capture(captureURL)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully archived %v via archive.is: %v\n", captureURL, archiveURL)
}

// Output:
//
// Successfully archived https://jaytaylor.com/ via archive.is: https://archive.is/i2PiW
```

##### Search for Existing Snapshots

[search.go](_examples/search/search.go):

```go
package main

import (
    "fmt"
    "time"

    "github.com/jaytaylor/archive.is"
)

var searchURL = "https://jaytaylor.com/"

func main() {
    snapshots, err := archiveis.Search(searchURL, 10*time.Second)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%# v\n", snapshots)
}

// Output:
//
//
```

### Running the test suite

    go test ./...

#### License

Permissive MIT license, see the [LICENSE](LICENSE) file for more information.
