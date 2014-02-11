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
				producer.Run(cf.SingleUri("cloudamqp-n/a"))
			},
		},
		{
			Name: "worker",
			ShortName: "w",
			Usage: "Run the Worker process",
			Action: func(c *cli.Context) {
				worker.Run(cf.SingleUri("cloudamqp-n/a"))
			},
		},
		{
			Name: "service",
			ShortName: "s",
			Usage: "Print the bound service URI for the provided service type",
			Action: func(c *cli.Context) {
				serviceType := c.Args().First()
				fmt.Println(cf.SingleUri(serviceType))
			},
		},
	}

	app.Run(os.Args)
}