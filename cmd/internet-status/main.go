package main

import (
	"fmt"
	"log"
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
	status, err := internetstatus.Full()
	if err != nil {
		log.Fatalf("Failed to check internet status: %v", err)
	}

	fmt.Println(status)
	return nil
}
