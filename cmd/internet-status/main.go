package main

import (
	"fmt"
	"os"

	"github.com/moul/internet-status"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Action = run
	app.Run(os.Args)
}

func run(c *cli.Context) error {
	fmt.Println(internetstatus.Full())
	return nil
}
