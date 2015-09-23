package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/humboldtux/leanote/api"
	"github.com/humboldtux/leanote/config"
)

var TagCommand = cli.Command{
	Name:    "tag",
	Aliases: []string{"t"},
	Usage:   "Tag commands.",
	Subcommands: []cli.Command{
		{
			Name:    "create",
			Aliases: []string{"a"},
			Usage:   "Add a tag.",
			Action:  addTag,
		},
		{
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "Delete a tag.",
			Action:  deleteTag,
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "after-usn, u",
					Value: -1,
				},
				cli.IntFlag{
					Name:  "max-entry, m",
					Value: 10000,
				},
			},
			Usage:  "List tags.",
			Action: listTags,
		},
	},
}

func addTag(ctx *cli.Context) {
	if len(ctx.Args()) != 1 {
		log.Fatal("Please provide just one argument as tag name to add")
	} else {
		name := ctx.Args()[0]
		conf, _ := config.New()
		c := api.NewClient(conf)
		tag, err := c.Tags.AddTag(name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to add tag: %s\n", err)
		} else {
			fmt.Printf("%+v\n", tag)
		}
	}
}

func deleteTag(ctx *cli.Context) {
	if len(ctx.Args()) != 1 {
		log.Fatal("Please provide just one argument as tag name to delete")
	} else {
		name := ctx.Args()[0]
		conf, _ := config.New()
		c := api.NewClient(conf)
		_, err := c.Tags.DeleteTag(name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to delete tag: %s\n", err)
		} else {
			fmt.Printf("Tag %s deleted\n", name)
		}
	}
}

func listTags(ctx *cli.Context) {

	afterUsn := ctx.Int("after-usn")
	maxEntry := ctx.Int("max-entry")

	conf, _ := config.New()
	c := api.NewClient(conf)
	tags, err := c.Tags.GetSyncTags(afterUsn, maxEntry)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to list tags: %s\n", err)
	}
	for _, tag := range tags {
		fmt.Printf("%+v\n", tag)
	}
}
