package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Tests that a book's Rust code samples compile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()

	testCmd.Flags().StringP("dest-dir", "d", "", "Output directory for the book")
	testCmd.Flags().BoolP("help", "h", false, "Prints help information")
	testCmd.Flags().StringP("library-path", "L", "", "A comma-separated list of directories to add to ")
	testCmd.Flags().BoolP("version", "V", false, "Prints version information")
	rootCmd.AddCommand(testCmd)

	carapace.Gen(testCmd).FlagCompletion(carapace.ActionMap{
		"dest-dir":     carapace.ActionDirectories(),
		"library-path": carapace.ActionDirectories(),
	})

	carapace.Gen(testCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
