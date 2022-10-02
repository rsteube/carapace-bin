package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var fmtCmd = &cobra.Command{
	Use:   "fmt",
	Short: "gofmt (reformat) package sources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fmtCmd).Standalone()

	fmtCmd.Flags().BoolS("mod", "mod", false, "set module download mode")
	fmtCmd.Flags().BoolS("n", "n", false, "print commands that would be executed")
	fmtCmd.Flags().BoolS("x", "x", false, "print commands as they are executed")
	rootCmd.AddCommand(fmtCmd)

	carapace.Gen(fmtCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
