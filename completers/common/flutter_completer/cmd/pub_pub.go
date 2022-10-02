package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pub_pubCmd = &cobra.Command{
	Use:   "pub",
	Short: "Pass the remaining arguments to Dart's \"pub\" tool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_pubCmd).Standalone()

	pub_pubCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pubCmd.AddCommand(pub_pubCmd)
}
