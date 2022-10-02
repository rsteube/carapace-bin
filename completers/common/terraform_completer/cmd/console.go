package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var consoleCmd = &cobra.Command{
	Use:   "console",
	Short: "Try Terraform expressions at an interactive command prompt",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(consoleCmd).Standalone()

	consoleCmd.Flags().StringS("state", "state", "", "Legacy option for the local backend only.")
	consoleCmd.Flags().StringArrayS("var", "var", []string{}, "Set a variable in the Terraform configuration.")
	consoleCmd.Flags().StringS("var-file", "var-file", "", "Set variables in the Terraform configuration from a file.")

	consoleCmd.Flag("state").NoOptDefVal = " "
	consoleCmd.Flag("var-file").NoOptDefVal = " "

	rootCmd.AddCommand(consoleCmd)

	carapace.Gen(consoleCmd).FlagCompletion(carapace.ActionMap{
		"state":    carapace.ActionFiles(),
		"var-file": carapace.ActionFiles(),
	})
}
