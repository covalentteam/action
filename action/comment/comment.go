package comment

import (
	"context"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

var (
	failedCreateComment       = errors.New("failed to create comment on client")
	failedCommentIsNotCreated = errors.New("failed to create comment on server")
)

func Do(ctx context.Context, cmd *Comment) error {
	staticToken := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: cmd.Token,
		},
	)

	oauthC := oauth2.NewClient(ctx, staticToken)
	client := github.NewClient(oauthC)

	comment := &github.IssueComment{
		Body: github.String("Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
	}

	_, r, err := client.Issues.CreateComment(ctx, cmd.OrganizationName, cmd.RepositoryName, cmd.PullRequestID, comment)
	if err != nil {
		return errors.Wrap(err, failedCreateComment.Error())
	}

	if r.StatusCode != http.StatusCreated {
		return failedCommentIsNotCreated
	}

	return nil
}
