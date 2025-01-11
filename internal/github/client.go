package github

import (
	"context"
	"os"

	"github.com/google/go-github/v68/github"
	"golang.org/x/oauth2"
)

type Client struct {
	api *github.Client
}

// optional auth token.
func CreateClient() *Client {
	ctx := context.Background()
	token := os.Getenv("GITHUB_TOKEN")
	if token != "" {
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		tc := oauth2.NewClient(ctx, ts)
		return &Client{api: github.NewClient(tc)}
	}
	return &Client{api: github.NewClient(nil)}
}

func (c *Client) FetchUserProfile(username string) (*github.User, error) {
	user, _, err := c.api.Users.Get(context.Background(), username)
	return user, err
}

func (c *Client) FetchUserRepos(username string) ([]*github.Repository, error) {
	opt := &github.RepositoryListByUserOptions{
		Type: "owner",
		Sort: "pushed",
	}
	repos, _, err := c.api.Repositories.ListByUser(context.Background(), username, opt)
	return repos, err
}
