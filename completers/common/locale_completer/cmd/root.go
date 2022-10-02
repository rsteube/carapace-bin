package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/localectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "locale",
	Short: "Get locale-specific information",
	Long:  "https://linux.die.net/man/1/locale",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all-locales", "a", false, "Write names of available locales")
	rootCmd.Flags().BoolP("category-name", "c", false, "Write names of selected categories")
	rootCmd.Flags().BoolP("charmaps", "m", false, "Write names of available charmaps")
	rootCmd.Flags().BoolP("help", "?", false, "Give this help list")
	rootCmd.Flags().BoolP("keyword-name", "k", false, "Write names of selected keywords")
	rootCmd.Flags().Bool("usage", false, "Give a short usage message")
	rootCmd.Flags().BoolP("verbose", "v", false, "Print more information")
	rootCmd.Flags().BoolP("version", "V", false, "Print program version")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("all-locales").Changed ||
				rootCmd.Flag("charmaps").Changed {
				return carapace.ActionValues()
			}
			return action.ActionLocales()
		}),
	)
}
