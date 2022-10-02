package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "push a single package to the device and install it",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().BoolS("d", "d", false, "TODO no flag description")
	installCmd.Flags().BoolS("g", "g", false, "grant all runtime permissions")
	installCmd.Flags().Bool("instant", false, "cause the app to be installed as an ephemeral install app")
	installCmd.Flags().BoolS("l", "l", false, "TODO no flag description")
	installCmd.Flags().BoolS("r", "r", false, "replace existing application")
	installCmd.Flags().BoolS("s", "s", false, "TODO no flag description")
	installCmd.Flags().BoolS("t", "t", false, "allow test packages")
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).PositionalCompletion(
		carapace.ActionFiles(".apk"),
	)
}
