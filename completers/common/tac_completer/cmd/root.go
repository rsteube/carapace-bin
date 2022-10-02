package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tac",
	Short: "concatenate and print files in reverse",
	Long:  "https://linux.die.net/man/1/tac",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("before", "b", false, "attach the separator before instead of after")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("regex", "r", false, "interpret the separator as a regular expression")
	rootCmd.Flags().StringP("separator", "s", "", "use STRING as the separator instead of newline")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
