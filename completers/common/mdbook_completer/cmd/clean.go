package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Deletes a built book",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanCmd).Standalone()

	cleanCmd.Flags().StringP("dest-dir", "d", "", "Output directory for the book")
	cleanCmd.Flags().BoolP("help", "h", false, "Prints help information")
	cleanCmd.Flags().BoolP("version", "V", false, "Prints version information")
	rootCmd.AddCommand(cleanCmd)

	carapace.Gen(cleanCmd).FlagCompletion(carapace.ActionMap{
		"dest-dir": carapace.ActionDirectories(),
	})

	carapace.Gen(cleanCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
