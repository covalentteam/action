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
	"os"
	"strconv"

	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"github.com/sethvargo/go-githubactions"
)

var (
	missingInputGithubToken              = errors.New("Missing 'token' input")
	missingRepositoryOwnerInput          = errors.New("Missing 'owner' input")
	missingRepositoryNameInput           = errors.New("Missing 'repository' input")
	missingPullRequestIdInput            = errors.New("Missing 'pull_request_id' input")
	missingEnvironmentVarGithubToken     = errors.New("Missing 'GITHUB_TOKEN' environment vars")
	missingEnvironmentVarPullRequestId   = errors.New("Missing 'PULL_REQUEST_ID' environment vars")
	missingEnvironmentVarRepositoryName  = errors.New("Missing 'REPO_NAME' environment vars")
	missingEnvironmentVarRepositoryOwner = errors.New("Missing 'REPO_OWNER' environment vars")
	unprocessablePullRequestID           = errors.New("Unprocessable 'pull_request_id' value")
)

type PullRequestReviewComment struct {
	Owner         *string
	Repo          *string
	PullRequestID *int
	Token         string
}

func newPullRequestReviewCommentFromGithub() (*PullRequestReviewComment, error) {
	owner := githubactions.GetInput("owner")
	if owner == "" {
		return nil, missingRepositoryOwnerInput
	}

	repository := githubactions.GetInput("repository")
	if repository == "" {
		return nil, missingRepositoryNameInput
	}

	token := githubactions.GetInput("token")
	if token == "" {
		return nil, missingInputGithubToken
	}

	prID := githubactions.GetInput("pull_request_id")
	if prID == "" {
		return nil, missingPullRequestIdInput
	}

	pullRequestId, err := strconv.Atoi(prID)
	if err != nil {
		return nil, unprocessablePullRequestID
	}

	reviewComment := &PullRequestReviewComment{
		Owner:         github.String(owner),
		Repo:          github.String(repository),
		PullRequestID: github.Int(pullRequestId),
		Token:         token,
	}

	return reviewComment, nil
}

func newPullRequestReviewCommentFromEnvironment() (*PullRequestReviewComment, error) {
	owner, ok := os.LookupEnv("REPOSITORY_OWNER")
	if !ok {
		return nil, missingEnvironmentVarRepositoryOwner
	}

	repo, ok := os.LookupEnv("REPOSITORY_NAME")
	if !ok {
		return nil, missingEnvironmentVarRepositoryName
	}

	githubToken, ok := os.LookupEnv("GITHUB_TOKEN")
	if !ok {
		return nil, missingEnvironmentVarGithubToken
	}

	id, ok := os.LookupEnv("PULL_REQUEST_ID")
	if !ok {
		return nil, missingEnvironmentVarPullRequestId
	}

	pullRequestID, err := strconv.Atoi(id)
	if err != nil {
		return nil, unprocessablePullRequestID
	}

	reviewComment := &PullRequestReviewComment{
		PullRequestID: github.Int(pullRequestID),
		Owner:         github.String(owner),
		Repo:          github.String(repo),
		Token:         githubToken,
	}

	return reviewComment, nil
}
