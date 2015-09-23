package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/humboldtux/leanote/api"
	"github.com/humboldtux/leanote/config"
)

var NoteCommand = cli.Command{
	Name:    "note",
	Aliases: []string{"n"},
	Usage:   "Note commands.",
	Subcommands: []cli.Command{
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "Create a note.",
			Action:  noteCreate,
		},
		{
			Name:    "edit",
			Aliases: []string{"e"},
			Usage:   "Edit a note.",
			Action:  noteEdit,
		},
		{
			Name:    "show",
			Aliases: []string{"sh"},
			Usage:   "Show notes.",
			Action:  noteShow,
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
			Usage:  "Notes sync state.",
			Action: noteState,
		},
	},
}

func noteCreate(ctx *cli.Context) {
	log.Fatal("Create not yet implemented")
}

func noteEdit(ctx *cli.Context) {
	log.Fatal("Edit not yet implemented")
}

func noteShow(ctx *cli.Context) {
	conf, _ := config.New()
	c := api.NewClient(conf)
	if len(ctx.Args()) > 0 {
		for _, noteId := range ctx.Args() {
			note, err := c.Notes.GetNoteAndContent(noteId)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to show notes %d: %s\n", noteId, err)
			}
			fmt.Printf("%+v\n", note)
		}
	}

	if len(ctx.Args()) == 0 {
		notes, err := c.Notes.GetNotes()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list all Notes: %s\n", err)
		}
		for _, note := range notes {
			fmt.Printf("%+v\n", note)
		}
	}
}

func noteState(ctx *cli.Context) {
	afterUsn := ctx.Int("after-usn")
	maxEntry := ctx.Int("max-entry")
	conf, _ := config.New()
	c := api.NewClient(conf)

	notes, err := c.Notes.GetSyncNotes(afterUsn, maxEntry)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get notes sync state: %s\n", err)
	}
	for _, note := range notes {
		fmt.Printf("%+v\n", note)
	}
}
