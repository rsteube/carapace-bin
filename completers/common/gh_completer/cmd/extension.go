package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var extensionCmd = &cobra.Command{
	Use:     "extension",
	Short:   "Manage gh extensions",
	Aliases: []string{"extensions", "ext"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(extensionCmd).Standalone()
	rootCmd.AddCommand(extensionCmd)
}
