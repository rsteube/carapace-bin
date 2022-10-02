package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watches a book's files and rebuilds it on changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(watchCmd).Standalone()

	watchCmd.Flags().StringP("dest-dir", "d", "", "Output directory for the book")
	watchCmd.Flags().BoolP("help", "h", false, "Prints help information")
	watchCmd.Flags().BoolP("open", "o", false, "Open the compiled book in a web browser")
	watchCmd.Flags().BoolP("version", "V", false, "Prints version information")
	rootCmd.AddCommand(watchCmd)

	carapace.Gen(watchCmd).FlagCompletion(carapace.ActionMap{
		"dest-dir": carapace.ActionDirectories(),
	})

	carapace.Gen(watchCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
