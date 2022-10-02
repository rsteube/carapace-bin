package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var bundle_createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a bundle",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bundle_createCmd).Standalone()

	bundle_createCmd.Flags().Bool("all-progress", false, "show progress meter during object writing phase")
	bundle_createCmd.Flags().Bool("all-progress-implied", false, "similar to --all-progress when progress meter is shown")
	bundle_createCmd.Flags().Bool("progress", false, "show progress meter")
	bundle_createCmd.Flags().BoolP("quiet", "q", false, "do not show progress meter")
	bundle_createCmd.Flags().String("version", "", "specify bundle format version")
	bundleCmd.AddCommand(bundle_createCmd)

	carapace.Gen(bundle_createCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)

	carapace.Gen(bundle_createCmd).PositionalAnyCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
