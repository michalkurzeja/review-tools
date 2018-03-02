package tokenstats

import (
	"github.com/urfave/cli"

	"context"
)

func NewLabelStats() *cli.Command {
	return &cli.Command{
		Name:     "label-stats",
		Aliases:  []string{"ls"},
		Category: "Tokens",
		Usage:    "Loads and presents the number of label occurrences in the reviews",
		Action: func(c *cli.Context) error {
			return action(c)
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "token, t",
				EnvVar: "GITHUB_TOKEN",
			},
			cli.StringFlag{
				Name:   "owner, o",
				EnvVar: "REPO_OWNER",
			},
			cli.StringFlag{
				Name:   "repository, r",
				EnvVar: "REPO_NAME",
			},
		},
	}
}

func action(cliCtx *cli.Context) error {
	ctx := context.Background()
	collector := NewLabelTokenCollector(ctx, cliCtx)

	stats := collector.Collect(ctx)

	NewCliOutput().Output(stats)

	return nil
}
