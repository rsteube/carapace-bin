package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "i3-scrot",
	Short: "simple screenshot script",
	Long:  "https://gitlab.manjaro.org/packages/community/i3/i3-scrot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("desk", "d", false, "full screen")
	rootCmd.Flags().BoolP("help", "h", false, "display this information")
	rootCmd.Flags().BoolP("select", "s", false, "selection")
	rootCmd.Flags().BoolP("window", "w", false, "active window")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("desk").Changed || rootCmd.Flag("window").Changed {
				return carapace.ActionValues("1", "2", "3", "4", "5")
			} else {
				return carapace.ActionValues()
			}
		}),
	)
}
