package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-sentinel",
	Short: "GitSentinel analyzes GitHub repositories for risk and reputation.",
	Long: `A command-line tool to assess the risk and reputation score of GitHub repositories by
	analyzing various factors including activity, community engagement, and maintenance.`,
	Version: "0.1.0",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
