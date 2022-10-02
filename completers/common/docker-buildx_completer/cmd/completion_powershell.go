package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var completion_powershellCmd = &cobra.Command{
	Use:   "powershell",
	Short: "Generate the autocompletion script for powershell",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completion_powershellCmd).Standalone()
	completion_powershellCmd.Flags().Bool("no-descriptions", false, "disable completion descriptions")
	completionCmd.AddCommand(completion_powershellCmd)
}
