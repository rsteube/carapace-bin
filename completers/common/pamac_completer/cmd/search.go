package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for packages or files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().BoolP("aur", "a", false, "also search in AUR")
	searchCmd.Flags().BoolP("files", "f", false, "search for packages which own the given filenames")
	searchCmd.Flags().BoolP("installed", "i", false, "only search for installed packages")
	searchCmd.Flags().Bool("no-aur", false, "do not search in AUR")
	searchCmd.Flags().BoolP("quiet", "q", false, "only print names")
	searchCmd.Flags().BoolP("repos", "r", false, "only search for packages in repositories")
	rootCmd.AddCommand(searchCmd)
}
