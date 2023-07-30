package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// set and get/print env from os
	os.Setenv("FOO", "BAR")
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0], " -> ", pair[1])
	}
}
