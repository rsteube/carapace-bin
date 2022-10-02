package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "unexpand",
	Short: "convert spaces to tabs",
	Long:  "https://linux.die.net/man/1/unexpand",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "convert all blanks, instead of just initial blanks")
	rootCmd.Flags().Bool("first-only", false, "convert only leading sequences of blanks (overrides -a)")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().StringP("tabs", "t", "", "have tabs N characters apart instead of 8 (enables -a) or use comma separated list of tab positions")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
