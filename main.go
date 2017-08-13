package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var (
	awsRegion     = "us-east-1"
	awsProfile    string
	awsFilterName string
	awsFilterTags string
	verbose       bool
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "region",
			Usage:       "AWS Region default `us-east-1`",
			Value:       "us-east-1",
			Destination: &awsRegion,
		},
		cli.StringFlag{
			Name:        "profile",
			Usage:       "AWS Config Profile to use `default`",
			Value:       "default",
			Destination: &awsProfile,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "ec2",
			Aliases: []string{"e"},
			Usage:   "AWS EC2 list",
			//Action:  ec2,
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Aliases: []string{"ls"},
					Usage:   "List ec2 instances",
					Action:  ec2List,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:        "filter, f",
							Usage:       "Filter based on EC2 tag:Name eg: `'web'`",
							Destination: &awsFilterName,
						},
						cli.StringFlag{
							Name:        "tag, t",
							Usage:       "Filter based on EC2 Tags eg: `env=prod[,cost_center=web]`",
							Destination: &awsFilterTags,
						},
					},
				},
			},
		},
		{
			Name:  "elb",
			Usage: "AWS Elastic Loadbalancer",
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Aliases: []string{"ls"},
					Usage:   "List elastic loadbalancers",
					Action:  elbList,
				},
			},
		},
		{
			Name:  "rds",
			Usage: "AWS RDS",
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Aliases: []string{"ls"},
					Usage:   "List RDS Instances",
					Action:  rdsList,
				},
			},
		},
		{
			Name:    "template",
			Aliases: []string{"t"},
			Usage:   "options for task templates",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add a new template",
					Action: func(c *cli.Context) error {
						fmt.Println("new task template: ", c.Args().First())
						return nil
					},
				},
				{
					Name:  "remove",
					Usage: "remove an existing template",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
