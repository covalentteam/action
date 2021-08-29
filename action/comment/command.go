package comment

import (
	"os"
	"strconv"

	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	core "github.com/sethvargo/go-githubactions"
)

var (
	missingInputGithubToken              = errors.New("Missing 'token' input")
	missingRepositoryOwnerInput          = errors.New("Missing 'owner' input")
	missingRepositoryNameInput           = errors.New("Missing 'repository' input")
	missingPullRequestIdInput            = errors.New("Missing 'pull_request_id' input")
	missingEnvironmentVarGithubSha       = errors.New("Missing 'GITHUB_SHA' environment vars")
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
	CommitId      string
	Token         string
}

func NewPullRequestReviewCommentFromGithub() (*PullRequestReviewComment, error) {
	owner := core.GetInput("owner")
	if owner == "" {
		return nil, missingRepositoryOwnerInput
	}

	repository := core.GetInput("repository")
	if repository == "" {
		return nil, missingRepositoryNameInput
	}

	token := core.GetInput("token")
	if token == "" {
		return nil, missingInputGithubToken
	}

	prID := core.GetInput("pull_request_id")
	if prID == "" {
		return nil, missingPullRequestIdInput
	}

	pullRequestId, err := strconv.Atoi(prID)
	if err != nil {
		return nil, unprocessablePullRequestID
	}

	commitID, ok := os.LookupEnv("GITHUB_SHA")
	if !ok {
		return nil, missingEnvironmentVarGithubSha
	}

	reviewComment := &PullRequestReviewComment{
		Owner:         github.String(owner),
		Repo:          github.String(repository),
		PullRequestID: github.Int(pullRequestId),
		CommitId:      commitID,
		Token:         token,
	}

	return reviewComment, nil
}

func NewPullRequestReviewCommentFromEnvironment() (*PullRequestReviewComment, error) {
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

	commitID, ok := os.LookupEnv("GITHUB_SHA")
	if !ok {
		return nil, missingEnvironmentVarGithubSha
	}

	reviewComment := &PullRequestReviewComment{
		PullRequestID: github.Int(pullRequestID),
		Owner:         github.String(owner),
		Repo:          github.String(repo),
		CommitId:      commitID,
		Token:         githubToken,
	}

	return reviewComment, nil
}
