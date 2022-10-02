package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var fonts_listCmd = &cobra.Command{
	Use:   "list",
	Short: "list verbs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fonts_listCmd).Standalone()

	fontsCmd.AddCommand(fonts_listCmd)
}
