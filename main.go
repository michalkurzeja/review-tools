package main

import (
	"log"
	"os"

	"review-tools/tokenstats"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "review-tools"
	app.Usage = "A set of tools for managing code reviews"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		*tokenstats.NewLabelStats(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
