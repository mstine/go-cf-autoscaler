package main

import (
	"os"
	"github.com/codegangsta/cli"
	"github.com/mstine/go-cf-autoscaler/producer"
	"github.com/mstine/go-cf-autoscaler/worker"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-cf-autoscaler"

	app.Commands = []cli.Command {
		{
			Name: "producer",
			ShortName: "p",
			Usage: "Run the Producer process",
			Action: func(c *cli.Context) {
				producer.Run()
			},
		},
		{
			Name: "worker",
			ShortName: "w",
			Usage: "Run the Worker process",
			Action: func(c *cli.Context) {
				worker.Run()
			},
		},
	}

	app.Run(os.Args)
}