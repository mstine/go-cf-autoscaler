package main

import (
	"fmt"
	"os"
	"github.com/codegangsta/cli"
	"github.com/mstine/go-cf-autoscaler/producer"
	"github.com/mstine/go-cf-autoscaler/worker"
	"github.com/mstine/go-cf-autoscaler/cf"
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
				producer.Run(cf.SingleAmqpUri())
			},
		},
		{
			Name: "worker",
			ShortName: "w",
			Usage: "Run the Worker process",
			Action: func(c *cli.Context) {
				worker.Run(cf.SingleAmqpUri())
			},
		},
		{
			Name: "service",
			ShortName: "s",
			Usage: "Print the bound service URI",
			Action: func(c *cli.Context) {
				fmt.Println(cf.SingleAmqpUri())
			},
		},
	}

	app.Run(os.Args)
}