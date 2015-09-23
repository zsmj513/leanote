package main

import (
	"log"

	"github.com/codegangsta/cli"
	"github.com/humboldtux/leanote/cmd"
)

const AppVersion = "0.0.1"

func init() {
	log.SetFlags(0)
	log.SetPrefix("leanote> ")
}

func main() {
	app := buildApp()
	app.RunAndExitOnError()
}

func buildApp() *cli.App {
	app := cli.NewApp()
	app.Name = "leanote"
	app.Version = AppVersion
	app.Usage = "Leanote Cli Tool."
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "email, e",
			Value:  "",
			Usage:  "Api email",
			EnvVar: "LEANOTE_EMAIL",
		},
		cli.StringFlag{
			Name:   "password, pwd",
			Value:  "",
			Usage:  "Api password",
			EnvVar: "LEANOTE_PWD",
		},
		cli.StringFlag{
			Name:   "api, a",
			Value:  "",
			Usage:  "API Url",
			EnvVar: "LEANOTE_APIURL",
		},
		cli.BoolFlag{Name: "debug,d", Usage: "Turn on debug output."},
	}
	app.Commands = []cli.Command{
		cmd.NotebookCommand,
		cmd.NoteCommand,
		cmd.TagCommand,
		cmd.UserCommand,
		cmd.FilesCommand,
	}

	return app
}
