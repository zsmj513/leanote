package cmd

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/humboldtux/leanote/api"
	"github.com/humboldtux/leanote/config"
)

var NotebookCommand = cli.Command{
	Name:    "notebook",
	Aliases: []string{"nb"},
	Usage:   "Notebook commands.",
	Subcommands: []cli.Command{
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "Create a notebook.",
			Action:  notebookCreate,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "seq, s",
					Value: 0,
				},
				cli.StringFlag{
					Name:  "parent, p",
					Value: "",
				},
			},
		},
		{
			Name:    "update",
			Aliases: []string{"u"},
			Usage:   "Update a notebook.",
			Action:  notebookUpdate,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "seq, s",
					Value: 1000000,
				},
				cli.StringFlag{
					Name:  "parent, p",
					Value: "",
				},
				cli.StringFlag{
					Name:  "title, t",
					Value: "",
				},
			},
		},

		{
			Name:    "show",
			Aliases: []string{"sh"},
			Usage:   "Show notebooks.",
			Action:  notebookShow,
		},
		{
			Name:    "state",
			Aliases: []string{"st"},
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
			Usage:  "Notebooks sync state.",
			Action: notebookState,
		},
	},
}

func notebookCreate(ctx *cli.Context) {
	if len(ctx.Args()) == 0 {
		fmt.Fprintf(os.Stderr, "Failed to create notebook without a title\n")
	} else {
		title := ctx.Args()[0]
		pid := ctx.String("parent")
		seq := ctx.Int("seq")

		conf, _ := config.New()
		c := api.NewClient(conf)
		notebook, err := c.Notebooks.AddNotebook(title, pid, seq)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create notebook: %s\n", err)
		}
		fmt.Printf("Notebook %s created\n%+v\n", title, notebook)
	}
}

func notebookUpdate(ctx *cli.Context) {
	if len(ctx.Args()) == 0 {
		fmt.Fprintf(os.Stderr, "Give the id of the notebook to update\n")
	} else {
		conf, _ := config.New()
		c := api.NewClient(conf)

		nid := ctx.Args()[0]
		title := ctx.String("title")
		pid := ctx.String("parent")
		seq := ctx.Int("seq")

		notebook, err := c.Notebooks.UpdateNotebook(nid, title, pid, seq)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to update notebook: %s\n", err)
		}
		fmt.Printf("Notebook %s updated\n%+v\n", notebook)
	}
}

func notebookShow(ctx *cli.Context) {

	conf, _ := config.New()
	c := api.NewClient(conf)

	if len(ctx.Args()) > 0 {
		for _, notebookId := range ctx.Args() {
			notes, err := c.Notes.NotesFromNotebook(notebookId)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to list Notes for Notebook %d: %s\n", notebookId, err)
			}
			fmt.Printf("#########\nNotebook %s\n#############\n", notebookId)
			for _, note := range notes {
				fmt.Printf("%+v\n", note)
			}
			fmt.Println("#########")
		}
	}

	if len(ctx.Args()) == 0 {
		notebooks, err := c.Notebooks.GetNotebooks()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list Notebooks: %s\n", err)
		}
		for _, notebook := range notebooks {
			fmt.Printf("%+v\n", notebook)
		}
	}
}

func notebookState(ctx *cli.Context) {
	afterUsn := ctx.Int("after-usn")
	maxEntry := ctx.Int("max-entry")

	conf, _ := config.New()
	c := api.NewClient(conf)

	notebooks, err := c.Notebooks.GetSyncNotebooks(afterUsn, maxEntry)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get Notebooks sync state: %s\n", err)
	}
	for _, notebook := range notebooks {
		fmt.Printf("%+v\n", notebook)
	}
}
