package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var citoolCmd = &cobra.Command{
	Use:   "citool",
	Short: "Graphical alternative to git-commit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(citoolCmd).Standalone()
	rootCmd.AddCommand(citoolCmd)
}
