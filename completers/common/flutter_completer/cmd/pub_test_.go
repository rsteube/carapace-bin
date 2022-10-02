package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pub_testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run the test package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_testCmd).Standalone()

	pub_testCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pubCmd.AddCommand(pub_testCmd)
}
