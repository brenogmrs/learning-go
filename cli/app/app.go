package app

import "github.com/urfave/cli"

// Generate() will return the cli ready to be used
func Generate() *cli.App {
	app := cli.NewApp()

	app.Name = "CLI"
	app.Usage = "Fetch IPS and Server names"

	return app
}
