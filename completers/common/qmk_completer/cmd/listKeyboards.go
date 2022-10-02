package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var listKeyboardsCmd = &cobra.Command{
	Use:   "list-keyboards",
	Short: "List the keyboards currently defined within QMK",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listKeyboardsCmd).Standalone()

	listKeyboardsCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.AddCommand(listKeyboardsCmd)
}
