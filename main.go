package main

import (
	"fmt"
	"os"

	"github.com/stellaraf/go-as14525/internal/cmd"
)

func main() {
	cli := cmd.New()
	if err := cli.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
