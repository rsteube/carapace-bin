package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var bundle_listHeadsCmd = &cobra.Command{
	Use:   "list-heads",
	Short: "list the references defined in the bundle",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bundle_listHeadsCmd).Standalone()

	bundleCmd.AddCommand(bundle_listHeadsCmd)

	carapace.Gen(bundle_listHeadsCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)

	carapace.Gen(bundle_listHeadsCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return git.ActionBundleHeads(c.Args[0])
		}),
	)
}
