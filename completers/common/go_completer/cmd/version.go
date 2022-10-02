package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print Go version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()

	versionCmd.Flags().BoolS("m", "m", false, "print each executable's embedded module version information")
	versionCmd.Flags().BoolS("v", "v", false, "report unrecognized files")
	rootCmd.AddCommand(versionCmd)
}
