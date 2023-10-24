package cmd

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/stellaraf/go-as14525/bgp"
	"github.com/urfave/cli/v2"
)

var numberPattern = regexp.MustCompile(`\d{0,3}`)

func cmdBGPCommunityInverse() *cli.Command {
	return &cli.Command{
		Name:  "inverse",
		Usage: "Create an inverse match community based on a 3 digit ID",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "base",
				Value: "14525:51",
				Usage: "Base community and community prefix",
			},
		},
		Action: func(ctx *cli.Context) error {
			id := ctx.Args().First()
			base := ctx.String("base")
			valid := numberPattern.MatchString(id)
			if !valid {
				return fmt.Errorf("%s must be no more than 3 numbers", id)
			}
			num, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			result := bgp.InverseNumberMatch(num)
			fmt.Printf("%s%s\n", base, result)
			return nil
		},
	}
}

func cmdBGPCommunity() *cli.Command {
	return &cli.Command{
		Name:        "community",
		Usage:       "BGP Community Tools",
		Subcommands: []*cli.Command{cmdBGPCommunityInverse()},
	}
}

func cmdBGP() *cli.Command {
	return &cli.Command{
		Name:        "bgp",
		Usage:       "BGP Tools",
		Subcommands: []*cli.Command{cmdBGPCommunity()},
	}
}

func New() *cli.App {
	app := &cli.App{
		Name:  "as14525",
		Usage: "AS14525 Toolkit",
		Action: func(*cli.Context) error {
			return nil
		},
		Commands: []*cli.Command{cmdBGP()},
	}
	return app
}
