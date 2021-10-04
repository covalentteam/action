package comment

import (
	"os"
	"strconv"

	"github.com/pkg/errors"
	core "github.com/sethvargo/go-githubactions"
)

var (
	missingInputGitHubToken                  = errors.New("missing input 'token'")
	missingInputGitHubCommit                 = errors.New("missing input 'commit'")
	missingInputGitHubPullRequestID          = errors.New("missing input 'pull-request-id'")
	missingInputGitHubRepositoryName         = errors.New("missing input 'repository-name'")
	missingInputGitHubOrganizationName       = errors.New("missing input 'organization-name'")
	missingEnvironmentGitHubCommit           = errors.New("missing environment variable 'GITHUB_SHA'")
	missingEnvironmentGitHubToken            = errors.New("missing environment variable 'GITHUB_TOKEN'")
	missingEnvironmentGitHubOrganizationName = errors.New("missing environment variable 'ORGANIZATION_NAME'")
	missingEnvironmentGitHubRepositoryName   = errors.New("missing environment variable 'REPOSITORY_NAME'")
	missingEnvironmentGitHubPullRequestID    = errors.New("missing environment variable 'PULL_REQUEST_ID'")
	unprocessablePullRequestID               = errors.New("unprocessable PullRequest id")
)

type Comment struct {
	Organization  string
	PullRequestID int
	Ref           string
	Repository    string
	Token         string
}

type Option func(comment *Comment) error

func New(options ...Option) (comment *Comment, err error) {
	for _, option := range options {
		if err = option(comment); err != nil {
			return nil, err
		}
	}

	return comment, nil
}

func WithGithubCommitFromInput(comment *Comment) error {
	if input := core.GetInput("commit"); input != "" {
		comment.Ref = input
	}
	return missingInputGitHubCommit
}

func WithGithubTokenFromInput(comment *Comment) error {
	if input := core.GetInput("token"); input != "" {
		comment.Token = input
	}
	return missingInputGitHubToken
}

func WithOrganizationNameFromInput(comment *Comment) error {
	if input := core.GetInput("organization-name"); input != "" {
		comment.Organization = input
	}
	return missingInputGitHubOrganizationName
}

func WithPullRequestIDFromInput(comment *Comment) error {
	input := core.GetInput("pull-request-id")
	if input == "" {
		return missingInputGitHubPullRequestID
	}

	pullRequestId, err := strconv.Atoi(input)
	if err != nil {
		return unprocessablePullRequestID
	}

	comment.PullRequestID = pullRequestId
	return nil
}

func WithRepositoryNameFromInput(comment *Comment) error {
	if input := core.GetInput("repository_name"); input != "" {
		comment.Repository = input
	}

	return missingInputGitHubRepositoryName
}

func WithGithubCommitFromEnvironment(comment *Comment) error {
	if value, ok := os.LookupEnv("GITHUB_SHA"); ok {
		comment.Ref = value
	}
	return missingEnvironmentGitHubCommit
}

func WithGithubTokenFromEnvironment(comment *Comment) error {
	if value, ok := os.LookupEnv("GITHUB_TOKEN"); ok {
		comment.Token = value
	}
	return missingEnvironmentGitHubToken
}

func WithOrganizationNameFromEnvironment(comment *Comment) error {
	if value, ok := os.LookupEnv("ORGANIZATION_NAME"); ok {
		comment.Organization = value
	}
	return missingEnvironmentGitHubOrganizationName
}

func WithPullRequestIDFromEnvironment(comment *Comment) error {
	value, ok := os.LookupEnv("PULL_REQUEST_ID")
	if !ok {
		return missingEnvironmentGitHubPullRequestID
	}

	number, err := strconv.Atoi(value)
	if err != nil {
		return unprocessablePullRequestID
	}

	comment.PullRequestID = number
	return nil
}

func WithRepositoryNameFromEnvironment(comment *Comment) error {
	if value, ok := os.LookupEnv("REPOSITORY_NAME"); ok {
		comment.Repository = value
	}
	return missingEnvironmentGitHubRepositoryName
}
