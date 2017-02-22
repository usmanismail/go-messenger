package main

import (
	"os"

	"github.com/codegangsta/cli"
	"go-messenger/go-auth/app"
	"go-messenger/go-auth/logger"
)

func main() {

	cliApp := cli.NewApp()
	cliApp.Name = "go-auth"
	cliApp.Author = "Usman Ismail"
	cliApp.Email = "usman@techtraits.com"
	cliApp.Usage = "A RESTful Authentication Service with a Database backend"
	cliApp.Commands = []cli.Command{getRunCommand()}
	cliApp.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "log-level, l",
			Usage: "The log level to use",
			Value: "Info",
		},
		cli.StringFlag{
			Name:  "log-type, t",
			Usage: "The log type to use. console or syslog",
			Value: "console",
		},
	}
	cliApp.Run(os.Args)

}

func cliApplicationAction(c *cli.Context) {
	logger.SetupLogging(c.GlobalString("log-level"), c.GlobalString("log-type"))
}
func getRunCommand() cli.Command {

	actionRun := func(c *cli.Context) {
		cliApplicationAction(c)
		if !c.IsSet("db-host") {
			cli.ShowCommandHelp(c, "run")
			return
		}

		goAuthApp := app.NewApplication(c.String("db-user"), c.String("db-password"),
			c.String("database"), c.String("db-host"), c.Int("db-port"), c.Int("port"))

		goAuthApp.Run()

	}

	cmdRun := cli.Command{
		Name:   "run",
		Usage:  "Run the authentication service",
		Action: actionRun,
	}

	cmdRun.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "db-host",
			Usage: "The Database Hostname",
		},
		cli.IntFlag{
			Name:  "db-port",
			Usage: "The Database port",
			Value: 3306,
		},
		cli.StringFlag{
			Name:  "db-user",
			Usage: "The Database Username",
			Value: "messenger",
		},
		cli.StringFlag{
			Name:  "db-password",
			Usage: "The Database Password",
			Value: "messenger",
		},
		cli.StringFlag{
			Name:  "database",
			Usage: "The Database name",
			Value: "messenger",
		},
		cli.IntFlag{
			Name:  "port, p",
			Usage: "The port on which this app will serve requests",
			Value: 8080,
		},
	}

	return cmdRun
}
