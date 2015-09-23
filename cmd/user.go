package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/humboldtux/leanote/api"
	"github.com/humboldtux/leanote/config"
)

var UserCommand = cli.Command{
	Name:    "user",
	Aliases: []string{"u"},
	Usage:   "User commands.",
	Action:  info,
	Subcommands: []cli.Command{
		{
			Name:    "info",
			Aliases: []string{"i"},
			Usage:   "User info.",
			Action:  info,
		},
		{
			Name:    "state",
			Aliases: []string{"s"},
			Usage:   "Get user sync state.",
			Action:  getSyncState,
		},
		{
			Name:    "username",
			Aliases: []string{"u"},
			Usage:   "Update username.",
			Action:  updateUsername,
		},
		{
			Name:    "password",
			Aliases: []string{"p"},
			Usage:   "Update password.",
			Action:  updatePwd,
		},
		{
			Name:    "logo",
			Aliases: []string{"l"},
			Usage:   "Update logo.",
			Action:  updateLogo,
		},
		{
			Name:    "register",
			Aliases: []string{"r"},
			Usage:   "Register user.",
			Action:  register,
		},
	},
}

func info(ctx *cli.Context) {
	conf, _ := config.New()
	c := api.NewClient(conf)
	user, err := c.User.Info()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get user info: %s\n", err)
	}
	fmt.Printf("%+v\n", user)
	fmt.Printf("%+v\n", c.Account)
}

func getSyncState(ctx *cli.Context) {
	conf, _ := config.New()
	c := api.NewClient(conf)
	state, err := c.User.GetSyncState()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get user sync state: %s\n", err)
	}
	fmt.Printf("%+v\n", state)
}

func updatePwd(ctx *cli.Context) {
	if len(ctx.Args()) != 2 {
		log.Fatal("Please provide old and new passwd")
	} else {
		old := ctx.Args()[0]
		pwd := ctx.Args()[1]
		conf, _ := config.New()
		c := api.NewClient(conf)
		err := c.User.UpdatePwd(old, pwd)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to update passwd: %s\n", err)
		}
	}
}

func updateUsername(ctx *cli.Context) {
	if len(ctx.Args()) != 1 {
		log.Fatal("Please provide just one argument as filename")
	} else {
		name := ctx.Args()[0]
		conf, _ := config.New()
		c := api.NewClient(conf)
		err := c.User.UpdateUsername(name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to update username: %s\n", err)
		}
	}
}

func updateLogo(ctx *cli.Context) {
	if len(ctx.Args()) != 1 {
		log.Fatal("Please provide just one argument as new logo")
	} else {
		path := ctx.Args()[0]
		conf, _ := config.New()
		c := api.NewClient(conf)
		err := c.User.UpdateLogo(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to update logo: %s\n", err)
		}
	}
}

func register(ctx *cli.Context) {
	conf, _ := config.New()
	c := api.NewClient(conf)

	err := c.Auth.Register(conf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to register user: %s\n", err)
	}
	//Display user info
	info(ctx)
}
