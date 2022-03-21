package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate() will return the cli ready to be used
func Generate() *cli.App {
	app := cli.NewApp()

	app.Name = "CLI"
	app.Usage = "Fetch IPS and Server names"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Fetch IPs from internet addresses",
			Flags:  flags,
			Action: fetchIps,
		},
		{
			Name:   "server",
			Usage:  "Fetch Servers name from internet addresses",
			Flags:  flags,
			Action: fetchServers,
		},
	}

	return app
}

func fetchServers(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)

	fmt.Printf("The server names for %s are the following: \n", host)

	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}

}

func fetchIps(c *cli.Context) {

	host := c.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The ips for %s are the following: \n", host)

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
