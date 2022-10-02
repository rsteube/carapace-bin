package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "docker-scan",
	Short: "A tool to scan your images",
	Long:  "https://docs.docker.com/engine/scan/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("accept-license", false, "Accept using a third party scanning provider")
	rootCmd.Flags().Bool("dependency-tree", false, "Show dependency tree with scan results")
	rootCmd.Flags().Bool("exclude-base", false, "Exclude base image from vulnerability scanning (requires --file)")
	rootCmd.Flags().StringP("file", "f", "", "Dockerfile associated with image, provides more detailed results")
	rootCmd.Flags().Bool("group-issues", false, "Aggregate duplicated vulnerabilities and group them to a single one (requires --json)")
	rootCmd.Flags().Bool("json", false, "Output results in JSON format")
	rootCmd.Flags().Bool("login", false, "Authenticate to the scan provider using an optional token (with --token), or web base token if empty")
	rootCmd.Flags().Bool("reject-license", false, "Reject using a third party scanning provider")
	rootCmd.Flags().String("severity", "", "Only report vulnerabilities of provided level or higher (low|medium|high)")
	rootCmd.Flags().String("token", "", "Authentication token to login to the third party scanning provider")
	rootCmd.Flags().Bool("version", false, "Display version of the scan plugin")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"file":     carapace.ActionFiles(),
		"severity": carapace.ActionValues("low", "medium", "high"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		docker.ActionRepositoryTags(),
	)
}
