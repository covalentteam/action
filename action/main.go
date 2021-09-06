package main

import (
	"context"
	"flag"
	"os"

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
		core.Errorf("Failed to construct command: \n %s", err.Error())
		os.Exit(1)
	}

	if err := comment.Do(context.Background(), cmd); err != nil {
		core.Errorf("Failed to create comment: \n %s", err.Error())
		os.Exit(1)
	}
}

func newPullRequestReviewCommentCommand() (*comment.Comment, error) {
	var runsOn runsOn

	flag.IntVar((*int)(&runsOn), "runs-on", int(onActions), "Options: [1] manually [2] GitHub Actions")
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
