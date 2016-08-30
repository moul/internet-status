package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/moul/internet-status"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
	/*cli.BoolFlag{
		Name:   "wait, w",
		Usage:  "Blocks while periodically check for internet connectivity",
		EnvVar: "INTERNETSTATUS_WAIT",
	},*/
	}
	app.Action = run
	app.Run(os.Args)
}

func run(c *cli.Context) error {
	result := internetstatus.Full()
	log.Printf("Access(): %v", result.Access())
	log.Printf("IPv4Access(): %v", result.IPv4Access())
	log.Printf("IPv6Access(): %v", result.IPv6Access())

	jsonOutput, _ := json.MarshalIndent(result.Map(), "", "  ")
	log.Printf("Result: %s\n", jsonOutput)
	return nil
}
