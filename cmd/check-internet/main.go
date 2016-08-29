package main

import (
	"fmt"
	"log"
	"os"

	"github.com/moul/check-internet"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Action = run
	app.Run(os.Args)
}

func run(c *cli.Context) error {
	status, err := checkinternet.Full()
	if err != nil {
		log.Fatalf("Failed to check internet status: %v", err)
	}

	fmt.Println(status)
	return nil
}
