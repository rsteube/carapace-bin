package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/rifle_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rifle",
	Short: "ranger's file opener",
	Long:  "https://ranger.github.io/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("c", "c", "", "read config from specified file instead of default")
	rootCmd.Flags().StringS("f", "f", "", "use additional flags")
	rootCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.Flags().BoolS("l", "l", false, "list possible ways to open the files")
	rootCmd.Flags().StringS("p", "p", "", "pick a method to open the files")
	rootCmd.Flags().Bool("version", false, "show program's version number and exit")
	rootCmd.Flags().StringS("w", "w", "", "open the files with PROGRAM")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"c": carapace.ActionFiles(),
		"p": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return action.ActionMethods(c.Args[0])
			}
			return carapace.ActionValues()
		}),
		"w": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
