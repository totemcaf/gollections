# Gollections - A collections' library for Golang generics

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/mod/github.com/totemcaf/gollections)

Gollections is a small library with helper functions to handle common collections in a type safe manner.

It uses the new generics support in GO that allows to create generic functions let the compiler check the use
of the correct types.

# Requirements

# Installation

The recommended way to get started using the Gollections is by using Go modules to install the dependency in your project.

    go get github.com/totemcaf/gollections

When using a version of Go that does not support modules, the library can be installed using dep by running

    dep ensure -add "github.com/totemcaf/gollections"

# Usage

Instead of writing

```go
package main

import "fmt"

func SearchValue(as []string, toSearch string) int {
    for idx, a := range as {
        if a == toSearch {
            return idx
        }
    }
    return -1
}

func main() {
    strings := []string{"hi", "hello", "world"}

    fmt.Printf("%s, found at %d", "hello", SearchValue(strings, "hello"))
}
```

And then repeating for other type, you can now write:

```go
package main

import "fmt"

func SearchValue[A comparable](as []A, toSearch A) int {
    for idx, a := range as {
        if a == toSearch {
            return idx
        }
    }
    return -1
}

func main() {
    strings := []string{"hi", "hello", "world"}

    fmt.Printf("%s, found at %d", "hello", SearchValue(strings, "hello"))

    ints := []int{ 4, 14, 42 }
    fmt.Printf("%d, found at %d", 42, SearchValue(ints, 42))
}
```

And also new generic functions can be used:

```go
package main

import (
    "fmt"

    "github.com/totemcaf/gollections/lists"
)

func main() {
    strings := lists.Of("hi", "Hello", "world")

    fmt.Printf("%s, found at %d", "hello", lists.Index(strings, "hello"))
}
```

Composing the fluent style you can do, for example:

```go
package main

import (
    "fmt"

    "github.com/totemcaf/gollections/lists"
)

func main() {
    words := lists.Of("hi", "Hello", "world")

    wordSize := func(w string) int {return len(w)}
    isLong := func(n int) bool {return n > 4}
    
    longWordCount := lists.Count(lists.Map(words, wordSize), isLong) 

    fmt.Printf("found %d long words", longWordCount)
}
```
