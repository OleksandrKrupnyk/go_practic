/*
 List 2.7
*/
package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "hello_cli"
	app.Usage = "Print Hello world"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "name ,n",
			Value: "World",
			Usage: "Who to say hello to.",
		},
	}
	app.Action = func(c *cli.Context) error {
		name := c.GlobalString("name")
		fmt.Printf("Hello %s\n", name)
		return nil
	}
	app.Run(os.Args)
}
