package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates your application and its dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	updateCmd.Flags().Bool("allow-dirty", false, "Whether to allow updating when the repository contains modified or untracked files.")
	updateCmd.Flags().BoolP("create-commits", "C", false, "Create source control commits for updates and migrations.")
	updateCmd.Flags().Bool("force", false, "Ignore peer dependency version mismatches.")
	updateCmd.Flags().String("from", "", "Version from which to migrate from.")
	updateCmd.Flags().Bool("migrate-only", false, "Only perform a migration, do not update the installed version.")
	updateCmd.Flags().Bool("next", false, "Use the prerelease version, including beta and RCs.")
	updateCmd.Flags().String("packages", "", "The names of package(s) to update.")
	updateCmd.Flags().String("to", "", "Version up to which to apply migrations.")
	updateCmd.Flags().Bool("verbose", false, "Display additional details about internal operations during execution.")
	rootCmd.AddCommand(updateCmd)
}
