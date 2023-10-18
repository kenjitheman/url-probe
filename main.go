package main

import (
	"fmt"

	"github.com/kenjitheman/url-probe/core"
)

func main() {
    core.PrintUsage()
    args := core.ParseArgs()
    fmt.Println(args)

    // todo: all
}
