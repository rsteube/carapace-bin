package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Prints this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(helpCmd).Standalone()

	rootCmd.AddCommand(helpCmd)

	carapace.Gen(helpCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			vals := make([]string, 0)
			for _, c := range rootCmd.Commands() {
				if !c.Hidden {
					vals = append(vals, c.Use, c.Short)
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		}),
	)
}
