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
