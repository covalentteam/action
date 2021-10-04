package main

import (
	"context"
	"flag"

	"github.com/covalentteam/template/action/comment"
	"github.com/pkg/errors"
	core "github.com/sethvargo/go-githubactions"
)

type runsOn int

const (
	manually runsOn = iota + 1
	onActions
)

func main() {
	cmd, err := newPullRequestReviewCommentCommand()
	if err != nil {
		core.Fatalf("Failed to construct command: \n %s", err.Error())
	}

	if err := comment.Do(context.Background(), cmd); err != nil {
		core.Fatalf("Failed to create comment: \n %s", err.Error())
	}
}

func newPullRequestReviewCommentCommand() (*comment.Comment, error) {
	var runsOn runsOn
	var help bool

	flag.IntVar((*int)(&runsOn), "runs-on", int(onActions), "Runs on: \n[1] Manually \n[2] GitHub Actions")
	flag.BoolVar(&help, "help", help, "Display help command")
	flag.Parse()

	switch runsOn {
	case manually:
		return comment.New(
			comment.WithGithubCommitFromEnvironment,
			comment.WithGithubTokenFromEnvironment,
			comment.WithOrganizationNameFromEnvironment,
			comment.WithPullRequestIDFromEnvironment,
			comment.WithRepositoryNameFromEnvironment,
		)
	case onActions:
		return comment.New(
			comment.WithGithubCommitFromInput,
			comment.WithGithubTokenFromInput,
			comment.WithOrganizationNameFromInput,
			comment.WithPullRequestIDFromInput,
			comment.WithRepositoryNameFromInput,
		)
	default:
		return nil, errors.New("not supported runs-on option: %s")
	}
}
