package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var reproCmd = &cobra.Command{
	Use:   "repro",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reproCmd).Standalone()
	reproCmd.Flags().String("sdk-path", "", "override the sdk path to use")
	reproCmd.Flags().Duration("timeout", 0, "timeout to use")
	rootCmd.AddCommand(reproCmd)

	carapace.Gen(reproCmd).FlagCompletion(carapace.ActionMap{
		"sdk-path": carapace.ActionDirectories(),
	})

	carapace.Gen(reproCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
