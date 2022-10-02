package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a new versioned artifact from source",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().Bool("push", false, "Push the artifact to the configured registry.")

	addGlobalOptions(buildCmd)
	addOperationOptions(buildCmd)

	rootCmd.AddCommand(buildCmd)
}
