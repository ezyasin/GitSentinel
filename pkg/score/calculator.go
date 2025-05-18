package score

import (
	"time"

	"github.com/ezyasin/git-sentinel/pkg/github"
)

type Score struct {
	TotalScore      float64
	RiskScore       float64
	ReputationScore float64
}

func CalculateScore(repo *github.RepositoryInfo) Score {
	// Basic scoring factors
	activityScore := calculateActivityScore(repo)
	communityScore := calculateCommunityScore(repo)
	maintenanceScore := calculateMaintenanceScore(repo)

	// Combine scores with weights
	riskScore := (activityScore * 0.4) + (maintenanceScore * 0.6)
	reputationScore := (activityScore * 0.3) + (communityScore * 0.4) + (maintenanceScore * 0.3)
	totalScore := (riskScore * 0.4) + (reputationScore * 0.6)

	return Score{
		TotalScore:      totalScore,
		RiskScore:       riskScore,
		ReputationScore: reputationScore,
	}
}

func calculateActivityScore(repo *github.RepositoryInfo) float64 {
	// Check last commit date
	lastCommit, err := time.Parse("2006-01-02", repo.LastCommit)
	if err != nil {
		return 0.0
	}

	daysSinceLastCommit := time.Since(lastCommit).Hours() / 24
	if daysSinceLastCommit > 365 {
		return 0.2
	} else if daysSinceLastCommit > 180 {
		return 0.4
	} else if daysSinceLastCommit > 90 {
		return 0.6
	} else if daysSinceLastCommit > 30 {
		return 0.8
	}
	return 1.0
}

func calculateCommunityScore(repo *github.RepositoryInfo) float64 {
	score := 0.0

	// Stars contribution
	if repo.Stars > 1000 {
		score += 0.4
	} else if repo.Stars > 100 {
		score += 0.3
	} else if repo.Stars > 10 {
		score += 0.2
	} else {
		score += 0.1
	}

	// Forks contribution
	if repo.Forks > 100 {
		score += 0.3
	} else if repo.Forks > 10 {
		score += 0.2
	} else {
		score += 0.1
	}

	// Follower count contribution
	if repo.FollowerCount > 1000 {
		score += 0.3
	} else if repo.FollowerCount > 100 {
		score += 0.2
	} else {
		score += 0.1
	}

	return score
}

func calculateMaintenanceScore(repo *github.RepositoryInfo) float64 {
	score := 0.0

	// README presence
	if repo.HasReadme {
		score += 0.3
	}

	// Issue management
	totalIssues := repo.OpenIssues + repo.OpenPRs
	if totalIssues == 0 {
		score += 0.4
	} else if totalIssues < 10 {
		score += 0.3
	} else if totalIssues < 50 {
		score += 0.2
	} else {
		score += 0.1
	}

	// Account age contribution
	accountAge, err := time.Parse("2006-01-02", repo.AccountAge)
	if err == nil {
		years := time.Since(accountAge).Hours() / (24 * 365)
		if years > 5 {
			score += 0.3
		} else if years > 2 {
			score += 0.2
		} else {
			score += 0.1
		}
	}

	return score
}
