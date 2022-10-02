package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pub_uploaderCmd = &cobra.Command{
	Use:   "uploader",
	Short: "Manage uploaders for a package on pub.dartlang.org",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_uploaderCmd).Standalone()

	pub_uploaderCmd.Flags().StringP("directory", "C", "", "Run this in the directory<dir>.")
	pub_uploaderCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pub_uploaderCmd.Flags().Bool("package", false, "The package whose uploaders will be modified.")
	pubCmd.AddCommand(pub_uploaderCmd)

	carapace.Gen(pub_uploaderCmd).FlagCompletion(carapace.ActionMap{
		"directory": carapace.ActionDirectories(),
	})

	carapace.Gen(pub_uploaderCmd).PositionalCompletion(
		carapace.ActionValues("add", "remove"),
	)
}
