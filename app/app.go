package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate will prepare the application to be executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Find by IP"
	app.Usage = "Find IP and names of servers in internet"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Find IP and names of servers in internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: buscarIps,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
