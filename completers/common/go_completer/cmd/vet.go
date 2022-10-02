package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var vetCmd = &cobra.Command{
	Use:   "vet",
	Short: "report likely mistakes in packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vetCmd).Standalone()

	vetCmd.Flags().BoolS("n", "n", false, "print commands that would be executed")
	vetCmd.Flags().BoolS("x", "x", false, "print commands as they are executed")
	rootCmd.AddCommand(vetCmd)
}
