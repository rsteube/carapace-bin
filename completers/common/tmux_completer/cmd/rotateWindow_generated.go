package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rotateWindowCmd = &cobra.Command{
	Use:   "rotate-window",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rotateWindowCmd).Standalone()

	rotateWindowCmd.Flags().BoolS("D", "D", false, "TODO description")
	rotateWindowCmd.Flags().BoolS("U", "U", false, "TODO description")
	rotateWindowCmd.Flags().BoolS("Z", "Z", false, "TODO description")
	rotateWindowCmd.Flags().StringS("t", "t", "", "target-window")
	rootCmd.AddCommand(rotateWindowCmd)
}
