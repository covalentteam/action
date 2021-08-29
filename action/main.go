package main

import (
	"context"
	"flag"
	"os"

	"github.com/covalentteam/template/action/comment"
	"github.com/pkg/errors"
	core "github.com/sethvargo/go-githubactions"
)

var (
	failedFromEnvironmentVars   = errors.New("Create command from environment vars")
	failedFromGithubActionInput = errors.New("Create command from Github Actions input")
)

func main() {
	cmd, err := newPullRequestReviewCommentCommand()
	if err != nil {
		core.Errorf("Failed to construct command: \n %s", err.Error())
		os.Exit(1)
	}

	if err := comment.Create(context.Background(), cmd); err != nil {
		core.Errorf("Failed to create comment: \n %s", err.Error())
		os.Exit(1)
	}
}

func newPullRequestReviewCommentCommand() (*comment.PullRequestReviewComment, error) {
	var onActions bool

	flag.BoolVar(&onActions, "on-actions", false, "Running action locally")
	flag.Parse()

	if onActions {
		cmd, err := comment.NewPullRequestReviewCommentFromGithub()
		if err != nil {
			return nil, errors.Wrap(err, failedFromGithubActionInput.Error())
		}

		return cmd, nil
	}

	cmd, err := comment.NewPullRequestReviewCommentFromEnvironment()
	if err != nil {
		return cmd, errors.Wrap(err, failedFromEnvironmentVars.Error())
	}

	return cmd, nil
}
