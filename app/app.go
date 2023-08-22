package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate command line application
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Application"
	app.Usage = "Search for IPs and Names of servers on the Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "github.com/Kenji-Sh",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Get IPs address from server",
			Flags: flags,
			Action: getIPs,
		},
		{
            Name: "server",
            Usage: "Get server name from internet",
            Flags: flags,
            Action: getServer,
        },
	}

	return app
}

func getIPs(c *cli.Context) {
	host := c.String("host")

	ips, error := net.LookupIP(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func getServer(c *cli.Context) {
    host := c.String("host")

    servers, error := net.LookupNS(host)
    if error != nil {
        log.Fatal(error)
    }

    for _, server := range servers {
        fmt.Println(server)
    }
}