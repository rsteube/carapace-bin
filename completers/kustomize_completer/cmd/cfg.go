package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cfgCmd = &cobra.Command{
	Use:   "cfg",
	Short: "Commands for reading and writing configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cfgCmd).Standalone()
	rootCmd.AddCommand(cfgCmd)
}
