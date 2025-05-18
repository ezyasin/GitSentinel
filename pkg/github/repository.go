package github

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	gh "github.com/google/go-github/v58/github"
)

type RepositoryInfo struct {
	Owner         string
	Name          string
	Stars         int
	Forks         int
	OpenIssues    int
	OpenPRs       int
	LastCommit    string
	HasReadme     bool
	AccountAge    string
	FollowerCount int
	RepoCount     int
}

func ParseRepositoryURL(repoURL string) (owner, repo string, err error) {
	parsedURL, err := url.Parse(repoURL)
	if err != nil {
		return "", "", fmt.Errorf("invalid URL: %v", err)
	}

	parts := strings.Split(strings.Trim(parsedURL.Path, "/"), "/")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid repository URL format")
	}

	return parts[0], parts[1], nil
}

func GetRepositoryInfo(ctx context.Context, client *gh.Client, owner, repo string) (*RepositoryInfo, error) {
	// Get repository information
	repository, _, err := client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		return nil, fmt.Errorf("failed to get repository: %v", err)
	}

	// Get owner information
	user, _, err := client.Users.Get(ctx, owner)
	if err != nil {
		return nil, fmt.Errorf("failed to get user information: %v", err)
	}

	// Get repository contents to check for README
	_, _, _, err = client.Repositories.GetContents(ctx, owner, repo, "README.md", nil)
	hasReadme := err == nil

	// Get last commit
	commits, _, err := client.Repositories.ListCommits(ctx, owner, repo, &gh.CommitsListOptions{
		ListOptions: gh.ListOptions{PerPage: 1},
	})
	lastCommit := ""
	if err == nil && len(commits) > 0 {
		lastCommit = commits[0].Commit.Committer.Date.Format("2006-01-02")
	}

	return &RepositoryInfo{
		Owner:         owner,
		Name:          repo,
		Stars:         repository.GetStargazersCount(),
		Forks:         repository.GetForksCount(),
		OpenIssues:    repository.GetOpenIssuesCount(),
		OpenPRs:       repository.GetOpenIssuesCount(),
		LastCommit:    lastCommit,
		HasReadme:     hasReadme,
		AccountAge:    user.GetCreatedAt().Format("2006-01-02"),
		FollowerCount: user.GetFollowers(),
		RepoCount:     user.GetPublicRepos(),
	}, nil
}
