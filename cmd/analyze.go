package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/ezyasin/git-sentinel/pkg/github"
	"github.com/ezyasin/git-sentinel/pkg/score"
	gh "github.com/google/go-github/v58/github"
	"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze [repository-url]",
	Short: "Analyze a GitHub repository",
	Long: `Analyze a GitHub repository to calculate its risk and reputation score.
	The repository URL should be in the format: https://github.com/owner/repo`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repoURL := args[0]

		// Parse repository URL
		owner, repo, err := github.ParseRepositoryURL(repoURL)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing repository URL: %v\n", err)
			os.Exit(1)
		}

		// Create GitHub client
		client := gh.NewClient(nil)

		// Get repository information
		ctx := context.Background()
		repoInfo, err := github.GetRepositoryInfo(ctx, client, owner, repo)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting repository information: %v\n", err)
			os.Exit(1)
		}

		// Calculate scores
		scores := score.CalculateScore(repoInfo)

		// Print results
		fmt.Printf("\nRepository Analysis Results for %s/%s\n", owner, repo)
		fmt.Println("=========================================")
		fmt.Printf("Total Score: %.2f/1.00\n", scores.TotalScore)
		fmt.Printf("Risk Score: %.2f/1.00\n", scores.RiskScore)
		fmt.Printf("Reputation Score: %.2f/1.00\n", scores.ReputationScore)
		fmt.Println("\nRepository Details:")
		fmt.Printf("Stars: %d\n", repoInfo.Stars)
		fmt.Printf("Forks: %d\n", repoInfo.Forks)
		fmt.Printf("Open Issues: %d\n", repoInfo.OpenIssues)
		fmt.Printf("Open Pull Requests: %d\n", repoInfo.OpenPRs)
		fmt.Printf("Last Commit: %s\n", repoInfo.LastCommit)
		fmt.Printf("Has README: %v\n", repoInfo.HasReadme)
		fmt.Printf("Repo Owner Account Age: %s\n", repoInfo.AccountAge)
		fmt.Printf("Repo Owner Follower Count: %d\n", repoInfo.FollowerCount)
		fmt.Printf("Repo Owner Repository Count: %d\n", repoInfo.RepoCount)
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
}
