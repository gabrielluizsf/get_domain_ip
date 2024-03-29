package cli

import (
	"fmt"
	"net"

	"github.com/urfave/cli"
)

// Generate will return the command line application ready to run
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "command line application"
	app.Usage = "search IP's and server names on the internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "internet address IP search",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "servers",
			Usage:  "search the name of servers on the internet",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
		LogError(err);

	for _, ip := range ips {
		fmt.Println(ip);
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host");

	servers, err := net.LookupNS(host)
		LogError(err);

	for _, serve := range servers {
		fmt.Println(serve);
	}
}
