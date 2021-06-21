// Copyright © 2021 Covalentteam
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"context"
	"net/http"
	"time"

	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

var (
	failedCreateComment       = errors.New("Failed to create comment on client")
	failedCommentIsNotCreated = errors.New("Failed to create comment on server")
)

func create(ctx context.Context, cmd *PullRequestReviewComment) error {
	staticToken := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: cmd.Token,
		},
	)

	now := time.Now()
	oauthC := oauth2.NewClient(ctx, staticToken)
	client := github.NewClient(oauthC)

	comment := &github.PullRequestComment{
		Body:      github.String("pa pa pa pa ...pa ...pooooow"),
		CreatedAt: &now,
	}

	_, r, err := client.PullRequests.CreateComment(ctx, *cmd.Owner, *cmd.Repo, int(*cmd.PullRequestID), comment)
	if err != nil {
		return errors.Wrap(err, failedCreateComment.Error())
	}

	if r.StatusCode != http.StatusCreated {
		return failedCommentIsNotCreated
	}

	return nil
}
