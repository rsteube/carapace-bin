package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "compile and install packages and dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().BoolS("i", "i", false, "install the packages that are dependencies of the target")
	installCmd.Flags().StringS("o", "o", "", "set output file or directory")
	addBuildFlags(installCmd)
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"o": carapace.ActionFiles(),
	})
}
