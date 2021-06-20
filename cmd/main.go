package main

import (
	"context"
	"os"
	"time"

	"github.com/google/go-github/github"
	"github.com/sethvargo/go-githubactions"
	"golang.org/x/oauth2"
)

func main() {
	var token = githubactions.GetInput("token")
	if token == "" {
		githubactions.Errorf("Missing 'token' parameter")
		os.Exit(1)
	}

	var staticToken = oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: token,
		},
	)

	now := time.Now()
	ctx := context.Background()

	oauthC := oauth2.NewClient(ctx, staticToken)
	client := github.NewClient(oauthC)

	comment := github.RepositoryComment{
		Body:      github.String("pa pa pa pa ...pa ...pooooow"),
		CreatedAt: &now,
	}

	_, r, err := client.Repositories.CreateComment(ctx, "owner", "repo", "sha", &comment)
	if err != nil {
		githubactions.Errorf("Failed to create comment", err.Error())
		os.Exit(1)
	}

	if r.StatusCode != 201 {
		githubactions.Errorf("Unexpected status code: ", r.Status)
		os.Exit(1)
	}
}
