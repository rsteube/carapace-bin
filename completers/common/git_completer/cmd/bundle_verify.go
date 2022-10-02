package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var bundle_verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "check that a bundle file is valid",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bundle_verifyCmd).Standalone()

	bundle_verifyCmd.Flags().BoolP("quiet", "q", false, "do not show bundle details")
	bundleCmd.AddCommand(bundle_verifyCmd)

	carapace.Gen(bundle_verifyCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
