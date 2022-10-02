package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var attachCmd = &cobra.Command{
	Use:   "attach",
	Short: "Attach to running process and begin debugging.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(attachCmd).Standalone()
	attachCmd.Flags().Bool("continue", false, "Continue the debugged process on start.")
	rootCmd.AddCommand(attachCmd)

	carapace.Gen(attachCmd).PositionalCompletion(
		ps.ActionProcessIds(),
		carapace.ActionFiles(),
	)
}
