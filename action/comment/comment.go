package comment

import (
	"context"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

var (
	ErrorCreateComment     = errors.New("failed to create comment on client")
	ErrorNotCreatedComment = errors.New("failed to create comment on server")
)

func Do(ctx context.Context, cmd *Comment) error {
	comment := &github.IssueComment{
		Body: github.String("Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
	}

	static := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: cmd.Token,
	})

	hclient := oauth2.NewClient(ctx, static)
	gclient := github.NewClient(hclient)

	_, r, err := gclient.Issues.CreateComment(ctx, cmd.Organization, cmd.Repository, cmd.PullRequestID, comment)
	if err != nil {
		return errors.Wrap(err, ErrorCreateComment.Error())
	}

	if r.StatusCode != http.StatusCreated {
		return ErrorNotCreatedComment
	}

	return nil
}
