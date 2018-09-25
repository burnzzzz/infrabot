package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "infrabot"
	app.Usage = "easily tune your infrastructure"

	app.Commands = []cli.Command{

		{
			Name:    "namespace",
			Aliases: []string{"ns"},
			Usage:   "operations with namespaces",
			Subcommands: []cli.Command{
				{
					Name:    "create",
					Usage:   "add a new namespace",
					Aliases: []string{"cr"},
					Action: func(c *cli.Context) error {
						ns := Namespace{c.Args().First()}
						fmt.Println("Creating namespace: ", ns.name)
						ns.create()
						return nil
					},
				},
				{
					Name:    "remove",
					Usage:   "remove an existing namespace",
					Aliases: []string{"rm"},
					Action: func(c *cli.Context) error {
						ns := Namespace{c.Args().First()} //todo dry violated
						fmt.Println("Removing namespace: ", ns.name)
						ns.remove()
						return nil
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
