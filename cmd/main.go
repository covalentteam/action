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
	"flag"
	"os"

	"github.com/pkg/errors"
	"github.com/sethvargo/go-githubactions"
)

var (
	failedFromEnvironmentVars   = errors.New("Create command from environment vars")
	failedFromGithubActionInput = errors.New("Create command from Github Actions input")
)

func main() {
	cmd, err := newPullRequestReviewCommentCommand()
	if err != nil {
		githubactions.Errorf("Failed to construct command: \n", err.Error())
		os.Exit(1)
	}

	if err := create(context.Background(), cmd); err != nil {
		githubactions.Errorf("Failed to create comment: \n", err.Error())
		os.Exit(1)
	}
}

func newPullRequestReviewCommentCommand() (*PullRequestReviewComment, error) {
	onActions := false

	flag.BoolVar(&onActions, "on-actions", false, "Running action locally")
	flag.Parse()

	if onActions {
		cmd, err := newPullRequestReviewCommentFromGithub()
		if err != nil {
			return nil, errors.Wrap(err, failedFromGithubActionInput.Error())
		}

		return cmd, nil
	}

	cmd, err := newPullRequestReviewCommentFromEnvironment()
	if err != nil {
		return cmd, errors.Wrap(err, failedFromEnvironmentVars.Error())
	}

	return cmd, nil
}
