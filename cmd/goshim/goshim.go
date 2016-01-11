package main

import (
	"os"

	"github.com/Songmu/goshim"
)

func main() {
	os.Exit(goshim.Run(os.Args[1:]))
}
