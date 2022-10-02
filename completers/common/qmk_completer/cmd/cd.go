package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cdCmd = &cobra.Command{
	Use:   "cd",
	Short: "Go to QMK Home",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cdCmd).Standalone()

	cdCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.AddCommand(cdCmd)
}
