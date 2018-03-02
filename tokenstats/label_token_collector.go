package tokenstats

import (
	"context"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/urfave/cli"
	"golang.org/x/oauth2"
)

type LabelTokenCollector struct {
	client *github.Client
	cliCtx *cli.Context
}

func (c *LabelTokenCollector) Collect(ctx context.Context) TokenStats {
	issues, _, err := c.client.Issues.ListByRepo(ctx, c.cliCtx.String("owner"), c.cliCtx.String("repository"), &github.IssueListByRepoOptions{})

	if err != nil {
		panic(err)
	}

	return extractStats(issues)
}

func NewLabelTokenCollector(ctx context.Context, cliCtx *cli.Context) *LabelTokenCollector {
	token := cliCtx.String("token")
	httpClient := getHttpClient(token, ctx)

	return &LabelTokenCollector{
		client: github.NewClient(httpClient),
		cliCtx: cliCtx,
	}
}

func getHttpClient(token string, ctx context.Context) *http.Client {
	var oauthClient *http.Client

	if len(token) > 0 {
		tokenSource := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		oauthClient = oauth2.NewClient(ctx, tokenSource)
	}

	return oauthClient
}

func extractStats(issues []*github.Issue) TokenStats {
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
