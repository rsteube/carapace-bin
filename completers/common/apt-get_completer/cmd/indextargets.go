package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var indextargetsCmd = &cobra.Command{
	Use:   "indextargets",
	Short: "Displays by default a deb822 formatted listing of information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(indextargetsCmd).Standalone()

	rootCmd.AddCommand(indextargetsCmd)
}
