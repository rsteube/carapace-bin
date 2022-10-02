package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "circleci",
	Short: "Use CircleCI from the command line",
	Long:  "https://github.com/CircleCI-Public/circleci-cli",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.PersistentFlags().Bool("debug", false, "Enable debug logging.")
	rootCmd.PersistentFlags().String("endpoint", "graphql-unstable", "URI to your CircleCI GraphQL API endpoint")
	rootCmd.PersistentFlags().String("github-api", "https://api.github.com/", "Change the default endpoint to GitHub API for retrieving updates")
	rootCmd.Flags().BoolP("help", "h", false, "help for circleci")
	rootCmd.PersistentFlags().String("host", "https://circleci.com", "URL to your CircleCI host, also CIRCLECI_CLI_HOST")
	rootCmd.PersistentFlags().Bool("skip-update-check", false, "Skip the check for updates check run before every command.")
	rootCmd.PersistentFlags().String("token", "", "your token for using CircleCI, also CIRCLECI_CLI_TOKEN")
}
