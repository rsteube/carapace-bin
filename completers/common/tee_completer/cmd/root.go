package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tee",
	Short: "read from standard input and write to standard output and files",
	Long:  "https://linux.die.net/man/1/tee",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("append", "a", false, "append to the given FILEs, do not overwrite")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("ignore-interrupts", "i", false, "ignore interrupt signals")
	rootCmd.Flags().String("output-error", "", "set behavior on write error.")
	rootCmd.Flags().BoolS("p", "p", false, "diagnose errors writing to non pipes")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"output-error": carapace.ActionValuesDescribed(
			"warn", "diagnose errors writing to any output",
			"warn-nopipe", "diagnose errors writing to any output not a pipe",
			"exit", "exit on error writing to any output",
			"exit-nopipe", "exit on error writing to any output not a pipe",
		),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
