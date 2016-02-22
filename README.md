# minietcd [![Travis badge](https://travis-ci.org/toqueteos/minietcd.svg?branch=master)](https://travis-ci.org/toqueteos/minietcd) [![GoDoc](http://godoc.org/github.com/toqueteos/minietcd?status.png)](http://godoc.org/github.com/toqueteos/minietcd)

Super small and _"dumb"_ read-only client for [coreos/etcd](https://github.com/coreos/etcd).

Intented only for v2 protocol as of right now.

## Rationale

One of the main points of Go **was** compiler speed.

Since Go v1.5 this is not true anymore.

Now, consider CoreOS team decision to move both etcd's client and server code
into the same repository, things get even worse.

minietcd's use case is simple: get keys from etcd, do it simple, do it fast (in
terms of compiling speed).

I know there's a lot of stuff this library doesn't cover but it's intentional.

## Installation

Use: `go get github.com/toqueteos/minietcd`

## Example

```go
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/toqueteos/minietcd"
)

const (
    Endpoint = "127.0.0.1:4001"
    Root     = "foo"
)

func main() {
    conn := minietcd.New()
    conn.SetLoggingOutput(os.Stderr) // optional, os.Stdout by default

    if err := conn.Dial(Endpoint); err != nil {
        log.Fatalf("failed to connect to endpoint %q, error %q\n", Endpoint, err)
    }

    keys, err := conn.Keys(Root)
    if err != nil {
        log.Fatalf("failed to get %q keys with error %q\n", Root, err)
    }

    for k, v := range keys {
        fmt.Println(k, v)
    }
}
```
