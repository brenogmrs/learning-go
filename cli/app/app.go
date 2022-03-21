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

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Fetch IPs from internet addresses",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: fetchIps,
		},
	}

	return app
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
