package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "describe all top-level provider sets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showCmd).Standalone()

	showCmd.Flags().StringS("tags", "tags", "", "append build tags to the default wirebuild")
	rootCmd.AddCommand(showCmd)
}
