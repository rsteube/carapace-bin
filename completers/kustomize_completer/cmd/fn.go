package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var fnCmd = &cobra.Command{
	Use:   "fn",
	Short: "Commands for running functions against configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fnCmd).Standalone()
	rootCmd.AddCommand(fnCmd)
}
