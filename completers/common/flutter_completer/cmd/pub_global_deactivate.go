package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/flutter_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pub_global_deactivateCmd = &cobra.Command{
	Use:   "deactivate",
	Short: "Remove a previously activated package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_global_deactivateCmd).Standalone()

	pub_global_deactivateCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pub_globalCmd.AddCommand(pub_global_deactivateCmd)

	carapace.Gen(pub_global_deactivateCmd).PositionalCompletion(
		action.ActionActivePackages(),
	)
}
