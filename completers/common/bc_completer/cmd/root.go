package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bc",
	Short: "An arbitrary precision calculator language",
	Long:  "https://linux.die.net/man/1/bc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "print this usage and exit")
	rootCmd.Flags().BoolP("interactive", "i", false, "force interactive mode")
	rootCmd.Flags().BoolP("mathlib", "l", false, "use the predefined math routines")
	rootCmd.Flags().BoolP("quiet", "q", false, "don't print initial banner")
	rootCmd.Flags().BoolP("standard", "s", false, "non-standard bc constructs are errors")
	rootCmd.Flags().BoolP("version", "v", false, "print version information and exit")
	rootCmd.Flags().BoolP("warn", "w", false, "warn about non-standard bc constructs")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
