package cmd

import (
	"github.com/spf13/cobra"
)

var template_pull_stackCmd = &cobra.Command{
	Use:   "stack",
	Short: "Downloads templates specified in the function definition yaml file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	template_pull_stackCmd.Flags().Bool("debug", false, "Enable debug output")
	template_pull_stackCmd.Flags().Bool("overwrite", false, "Overwrite existing templates?")
	template_pullCmd.AddCommand(template_pull_stackCmd)
}
