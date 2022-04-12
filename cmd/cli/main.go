package main

import (
	"fmt"
	"os"

	"github.com/mribica/bills/cli"
)

func main() {
	app := cli.App{}
	err := app.LoadProviders()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	ch := make(chan string)
	app.Run(ch)

	for range app.Providers {
		fmt.Print(<-ch)
	}
}
