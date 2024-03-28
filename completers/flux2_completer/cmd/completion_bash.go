package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var completion_bashCmd = &cobra.Command{
	Use:   "bash",
	Short: "Generates bash completion scripts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completion_bashCmd).Standalone()
	completionCmd.AddCommand(completion_bashCmd)
}
