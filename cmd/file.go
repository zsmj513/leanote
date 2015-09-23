package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/humboldtux/leanote/api"
	"github.com/humboldtux/leanote/config"
)

var FilesCommand = cli.Command{
	Name:    "files",
	Aliases: []string{"f"},
	Usage:   "Files commands.",
	Subcommands: []cli.Command{
		{
			Name:    "image",
			Aliases: []string{"i"},
			Usage:   "Get image.",
			Action:  getImage,
		},
		{
			Name:    "attach",
			Aliases: []string{"a"},
			Usage:   "Get attach.",
			Action:  getAttach,
		},
		{
			Name:    "note",
			Aliases: []string{"n"},
			Usage:   "Get notes attach.",
			Action:  getAllAttachs,
		},
	},
}

func getAttach(ctx *cli.Context) {
	if len(ctx.Args()) != 1 {
		log.Fatal("Please provide just one argument as attach id")
	} else {
		id := ctx.Args()[0]
		conf, _ := config.New()
		c := api.NewClient(conf)
		err := c.Files.GetAttach(id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed get attach: %s\n", err)
		}
	}
}

func getAllAttachs(ctx *cli.Context) {
	if len(ctx.Args()) != 1 {
		log.Fatal("Please provide just one argument as note id")
	} else {
		id := ctx.Args()[0]
		conf, _ := config.New()
		c := api.NewClient(conf)
		err := c.Files.GetAllAttachs(id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed get all attachs: %s\n", err)
		}
	}
}

func getImage(ctx *cli.Context) {
	if len(ctx.Args()) != 1 {
		log.Fatal("Please provide just one argument as image id")
	} else {
		id := ctx.Args()[0]
		conf, _ := config.New()
		c := api.NewClient(conf)
		err := c.Files.GetAttach(id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed get image: %s\n", err)
		}
	}
}
