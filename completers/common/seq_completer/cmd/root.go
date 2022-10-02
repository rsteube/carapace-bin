package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "seq",
	Short: "print a sequence of numbers",
	Long:  "https://linux.die.net/man/1/seq",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("equal-width", "w", false, "equalize width by padding with leading zeroes")
	rootCmd.Flags().StringP("format", "f", "", "use printf style floating-point FORMAT")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().StringP("separator", "s", "", "use STRING to separate numbers (default: \n)")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
}
