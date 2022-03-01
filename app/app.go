package app

import "github.com/urfave/cli"

// Generate will prepare the application to be executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Find by IP"
	app.Usage = "Find IP and names of servers in internet"
	return app
}
