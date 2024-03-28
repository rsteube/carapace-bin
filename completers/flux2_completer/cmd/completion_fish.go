package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var completion_fishCmd = &cobra.Command{
	Use:   "fish",
	Short: "Generates fish completion scripts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completion_fishCmd).Standalone()
	completionCmd.AddCommand(completion_fishCmd)
}
