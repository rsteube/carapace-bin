package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var linkWindowCmd = &cobra.Command{
	Use:   "link-window",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(linkWindowCmd).Standalone()

	linkWindowCmd.Flags().BoolS("a", "a", false, "TODO description")
	linkWindowCmd.Flags().BoolS("b", "b", false, "TODO description")
	linkWindowCmd.Flags().BoolS("d", "d", false, "TODO description")
	linkWindowCmd.Flags().BoolS("k", "k", false, "TODO description")
	linkWindowCmd.Flags().StringS("s", "s", "", "src-window")
	linkWindowCmd.Flags().StringS("t", "t", "", "dst-window")
	rootCmd.AddCommand(linkWindowCmd)
}
