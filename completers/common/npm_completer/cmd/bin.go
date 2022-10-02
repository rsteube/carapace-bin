package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var binCmd = &cobra.Command{
	Use:   "bin",
	Short: "Display npm bin folder",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(binCmd).Standalone()
	binCmd.Flags().BoolP("global", "g", false, "operate in global mode")

	rootCmd.AddCommand(binCmd)
}
