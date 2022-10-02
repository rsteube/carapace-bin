package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var unlinkWindowCmd = &cobra.Command{
	Use:   "unlink-window",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unlinkWindowCmd).Standalone()

	unlinkWindowCmd.Flags().BoolS("k", "k", false, "TODO description")
	unlinkWindowCmd.Flags().StringS("t", "t", "", "target-window")
	rootCmd.AddCommand(unlinkWindowCmd)
}
