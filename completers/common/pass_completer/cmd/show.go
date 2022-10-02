package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pass_completer/cmd/action"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show existing password",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showCmd).Standalone()
	showCmd.Flags().StringP("clip", "c", "", "put on clipboard")
	rootCmd.AddCommand(showCmd)

	showCmd.Flag("clip").NoOptDefVal = " "

	carapace.Gen(showCmd).PositionalCompletion(action.ActionPassNames())
}
