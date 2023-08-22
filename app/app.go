package app

import "github.com/urfave/cli"

// Generate command line application
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Application"
	app.Usage = "Search for IPs and Names of servers on the Internet"

	return app
}