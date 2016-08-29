package main

import (
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
	result := internetstatus.Full()
	log.Printf("Access(): %v", result.Access())
	log.Printf("IPv4Access(): %v", result.IPv4Access())
	log.Printf("IPv6Access(): %v", result.IPv6Access())
	log.Printf("Result: %v\n", result)
	return nil
}
