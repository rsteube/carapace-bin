package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var fontsCmd = &cobra.Command{
	Use:   "fonts",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fontsCmd).Standalone()

	rootCmd.AddCommand(fontsCmd)
}
