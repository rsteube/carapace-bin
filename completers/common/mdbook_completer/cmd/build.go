package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Builds a book from its markdown files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().StringP("dest-dir", "d", "", "Output directory for the book")
	buildCmd.Flags().BoolP("help", "h", false, "Prints help information")
	buildCmd.Flags().BoolP("open", "o", false, "Opens the compiled book in a web browser")
	buildCmd.Flags().BoolP("version", "V", false, "Prints version information")
	rootCmd.AddCommand(buildCmd)

	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"dest-dir": carapace.ActionDirectories(),
	})

	carapace.Gen(buildCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
