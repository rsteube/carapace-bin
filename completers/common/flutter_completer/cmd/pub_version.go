package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pub_versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print pub version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_versionCmd).Standalone()

	pub_versionCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pubCmd.AddCommand(pub_versionCmd)
}
