package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var installMultipleCmd = &cobra.Command{
	Use:   "install-multiple",
	Short: "push multiple APKs to the device for a single package and install them",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installMultipleCmd).Standalone()

	installMultipleCmd.Flags().BoolS("d", "d", false, "TODO no flag description")
	installMultipleCmd.Flags().BoolS("g", "g", false, "grant all runtime permissions")
	installMultipleCmd.Flags().Bool("instant", false, "cause the app to be installed as an ephemeral install app")
	installMultipleCmd.Flags().BoolS("l", "l", false, "TODO no flag description")
	installMultipleCmd.Flags().BoolS("r", "r", false, "replace existing application")
	installMultipleCmd.Flags().BoolS("s", "s", false, "TODO no flag description")
	installMultipleCmd.Flags().BoolS("t", "t", false, "allow test packages")
	rootCmd.AddCommand(installMultipleCmd)

	carapace.Gen(installMultipleCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".apk"),
	)
}
