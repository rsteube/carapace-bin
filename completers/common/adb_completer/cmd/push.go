package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "copy local files/directories to device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pushCmd).Standalone()

	pushCmd.Flags().BoolS("Z", "Z", false, "disable compression")
	pushCmd.Flags().BoolS("n", "n", false, "dry run: push files to device without storing to the filesystem")
	pushCmd.Flags().Bool("sync", false, "only push files that are newer on the host than the device")
	pushCmd.Flags().StringS("z", "z", "", "enable compression with a specified algorithm (any, none, brotli)")
	rootCmd.AddCommand(pushCmd)

	carapace.Gen(pushCmd).FlagCompletion(carapace.ActionMap{
		"z": carapace.ActionValues("any", "none", "brotli"),
	})

	carapace.Gen(pushCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
