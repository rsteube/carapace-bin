package cmd

import (
	"github.com/spf13/cobra"
)

var completion_bashCmd = &cobra.Command{
	Use:   "bash",
	Short: "bash completion.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	completionCmd.AddCommand(completion_bashCmd)
}
