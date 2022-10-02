package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var selfUpdateCmd = &cobra.Command{
	Use:   "self-update",
	Short: "Update the restic binary",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(selfUpdateCmd).Standalone()
	selfUpdateCmd.Flags().String("output", "", "Save the downloaded file as `filename` (default: running binary itself)")
	rootCmd.AddCommand(selfUpdateCmd)

	carapace.Gen(selfUpdateCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
	})
}
