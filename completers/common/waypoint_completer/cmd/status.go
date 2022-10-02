package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "List and refresh application statuses",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statusCmd).Standalone()

	statusCmd.Flags().Bool("all-projects", false, "Output status about every project in a workspace.")
	statusCmd.Flags().Bool("json", false, "Output the status information as JSON.")
	statusCmd.Flags().Bool("refresh", false, "Refresh application status for the requested app or apps in a project.")
	statusCmd.Flags().Bool("verbose", false, "Display more details.")

	addGlobalOptions(statusCmd)

	rootCmd.AddCommand(statusCmd)
}
