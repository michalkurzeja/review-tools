package tokenstats

import (
	"context"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/urfave/cli"
	"golang.org/x/oauth2"
)

const pageSize = 100

type LabelTokenCollector struct {
	client *github.Client
	cliCtx *cli.Context
}

func (c *LabelTokenCollector) Collect(ctx context.Context) TokenStats {
	owner := c.cliCtx.String("owner")
	repo := c.cliCtx.String("repository")

	issues := c.collectIssues(ctx, owner, repo)

	return extractStats(issues)
}

func (c *LabelTokenCollector) collectIssues(ctx context.Context, owner string, repo string) []*github.Issue {
	var allIssues []*github.Issue
	page := 1

	for {
		issues, _, err := c.client.Issues.ListByRepo(ctx, owner, repo, &github.IssueListByRepoOptions{ListOptions: github.ListOptions{Page: page, PerPage: pageSize}})

		if err != nil {
			panic(err)
		}

		if len(issues) == 0 {
			break
		}

		allIssues = append(allIssues, issues...)
		page++
	}

	return allIssues
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
