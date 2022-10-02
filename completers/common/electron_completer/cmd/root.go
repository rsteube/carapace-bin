package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "electron",
	Short: "Build cross platform desktop apps with JavaScript, HTML, and CSS",
	Long:  "https://www.electronjs.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("abi", "a", false, "Print the Node ABI version.")
	rootCmd.Flags().BoolP("interactive", "i", false, "Open a REPL to the main process.")
	rootCmd.Flags().BoolP("require", "r", false, "Module to preload (option can be repeated).")
	rootCmd.Flags().BoolP("version", "v", false, "Print the version.")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
