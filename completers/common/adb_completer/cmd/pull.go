package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "copy files/dirs from device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pullCmd).Standalone()

	pullCmd.Flags().BoolS("Z", "Z", false, "disable compression")
	pullCmd.Flags().BoolS("a", "a", false, "preserve file timestamp and mode")
	pullCmd.Flags().StringS("z", "z", "", "enable compression with a specified algorithm (any, none, brotli)")
	rootCmd.AddCommand(pullCmd)

	carapace.Gen(pullCmd).FlagCompletion(carapace.ActionMap{
		"z": carapace.ActionValues("any", "none", "brotli"),
	})
}
