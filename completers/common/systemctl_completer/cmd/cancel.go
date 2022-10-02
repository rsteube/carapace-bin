package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "Cancel all, one, or more jobs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cancelCmd).Standalone()

	rootCmd.AddCommand(cancelCmd)

	// TODO job completion
}
