package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var label_cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clones labels from one repository to another",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(label_cloneCmd).Standalone()
	label_cloneCmd.Flags().BoolP("force", "f", false, "Overwrite labels in the destination repository")
	labelCmd.AddCommand(label_cloneCmd)

	carapace.Gen(label_cloneCmd).PositionalCompletion(
		action.ActionOwnerRepositories(label_cloneCmd),
	)
}
