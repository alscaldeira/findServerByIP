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

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Find IP of servers on internet",
			Flags:  flags,
			Action: findIps,
		},
		{
			Name:   "server",
			Usage:  "Find servers by IP on internet",
			Flags:  flags,
			Action: findServers,
		},
	}

	return app
}

func findIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func findServers(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
