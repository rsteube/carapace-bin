package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "expand",
	Short: "convert tabs to spaces",
	Long:  "https://linux.die.net/man/1/expand",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("initial", "i", false, "do not convert tabs after non blanks")
	rootCmd.Flags().StringP("tabs", "t", "", "have tabs N characters apart, not 8. use comma separated list of tab positions")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
