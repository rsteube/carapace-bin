package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "sync a local build from $ANDROID_PRODUCT_OUT to the device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(syncCmd).Standalone()

	syncCmd.Flags().BoolS("Z", "Z", false, "disable compression")
	syncCmd.Flags().BoolS("l", "l", false, "list files that would be copied, but don't copy them")
	syncCmd.Flags().BoolS("n", "n", false, "dry run: push files to device without storing to the filesystem")
	syncCmd.Flags().StringS("z", "z", "", "enable compression with a specified algorithm (any, none, brotli)")
	rootCmd.AddCommand(syncCmd)

	carapace.Gen(syncCmd).FlagCompletion(carapace.ActionMap{
		"z": carapace.ActionValues("any", "none", "brotli"),
	})

	carapace.Gen(syncCmd).PositionalCompletion(
		carapace.ActionValues("all", "data", "odm", "oem", "product", "system", "system_ext", "vendor"),
	)
}
