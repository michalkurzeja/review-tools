package tokenstats

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/urfave/cli"
	"golang.org/x/oauth2"
)

type LabelTokenCollector struct {
	client *github.Client
	cliCtx *cli.Context
}

func NewLabelTokenCollector(ctx context.Context, cliCtx *cli.Context) *LabelTokenCollector {
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cliCtx.String("token")},
	)
	oauthClient := oauth2.NewClient(ctx, tokenSource)

	return &LabelTokenCollector{
		client: github.NewClient(oauthClient),
		cliCtx: cliCtx,
	}
}

func (c *LabelTokenCollector) Collect(ctx context.Context) TokenStats {
	issues, _, err := c.client.Issues.ListByRepo(ctx, c.cliCtx.String("owner"), c.cliCtx.String("repository"), &github.IssueListByRepoOptions{})

	if err != nil {
		panic(err)
	}

	stats := make(map[string]map[string]int)

	for _, issue := range issues {
		user := issue.User.GetLogin()
		labels := issue.Labels

		for _, label := range labels {
			_, exists := stats[user]

			if !exists {
				stats[user] = make(map[string]int)
			}

			stats[user][label.GetName()]++
		}
	}

	return *NewTokenStats(stats)
}
