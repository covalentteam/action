package main

import (
	"context"
	"os"
	"strconv"
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

	var owner = githubactions.GetInput("owner")
	if owner == "" {
		githubactions.Errorf("Missing 'owner' parameter")
		os.Exit(1)
	}

	var repo = githubactions.GetInput("repo")
	if repo == "" {
		githubactions.Errorf("Missing 'repo' parameter")
		os.Exit(1)
	}

	var prID = githubactions.GetInput("number")
	if prID == "" {
		githubactions.Errorf("Missing 'number' parameter")
		os.Exit(1)
	}

	pr, err := strconv.Atoi(prID)
	if err != nil {
		githubactions.Errorf("Pull Request ID is not valid.")
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

	comment := &github.PullRequestComment{
		Body:      github.String("pa pa pa pa ...pa ...pooooow"),
		CreatedAt: &now,
	}

	_, r, err := client.PullRequests.CreateComment(ctx, owner, repo, pr, comment)
	if err != nil {
		githubactions.Errorf("Failed to create comment: \n", err.Error())
		os.Exit(1)
	}

	if r.StatusCode != 201 {
		githubactions.Errorf("Unexpected status code: \n", r.Status)
		os.Exit(1)
	}
}
