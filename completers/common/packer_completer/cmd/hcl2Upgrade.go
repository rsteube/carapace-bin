package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var hcl2UpgradeCmd = &cobra.Command{
	Use:   "hcl2-upgrade",
	Short: "Rewrites HCL2 config files to canonical format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hcl2UpgradeCmd).Standalone()

	hcl2UpgradeCmd.Flags().StringS("output-file", "output-file", "", "Set output file name.")
	hcl2UpgradeCmd.Flags().BoolS("with-annotations", "with-annotations", false, "Add helper annotation comments to the file to help new HCL2 users understand the template format.")
	rootCmd.AddCommand(hcl2UpgradeCmd)

	carapace.Gen(hcl2UpgradeCmd).FlagCompletion(carapace.ActionMap{
		"output-file": carapace.ActionFiles(),
	})

	carapace.Gen(hcl2UpgradeCmd).PositionalCompletion(
		carapace.ActionFiles(".json", ".hcl"),
	)
}
