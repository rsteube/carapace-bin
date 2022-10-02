package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Delete all the resources created",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(destroyCmd).Standalone()

	destroyCmd.Flags().Bool("auto-approve", false, "Confirm destroying all resources.")

	addGlobalOptions(destroyCmd)
	addOperationOptions(destroyCmd)

	rootCmd.AddCommand(destroyCmd)
}
