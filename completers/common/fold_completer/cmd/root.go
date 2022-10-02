package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fold",
	Short: "wrap each input line to fit in specified width",
	Long:  "https://linux.die.net/man/1/fold",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("bytes", "b", false, "count bytes rather than columns")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("spaces", "s", false, "break at spaces")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
	rootCmd.Flags().StringP("width", "w", "", "use WIDTH columns instead of 80")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
