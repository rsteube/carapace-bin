package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var gen_autocompleteCmd = &cobra.Command{
	Use:   "autocomplete",
	Short: "Generate shell autocompletion script for Hugo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	gen_autocompleteCmd.PersistentFlags().StringP("completionfile", "f", "", "autocompletion file, defaults to stdout")
	gen_autocompleteCmd.PersistentFlags().StringP("type", "t", "bash", "autocompletion type (bash, zsh, fish, or powershell)")
	genCmd.AddCommand(gen_autocompleteCmd)

	carapace.Gen(gen_autocompleteCmd).FlagCompletion(carapace.ActionMap{
		"completionfile": carapace.ActionFiles(),
		"type":           carapace.ActionValues("bash", "zsh", "fish", "powershell"),
	})
}
