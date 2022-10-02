package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pub_publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish the current package to pub.dartlang.org",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_publishCmd).Standalone()

	pub_publishCmd.Flags().StringP("directory", "C", "", "Run this in the directory<dir>.")
	pub_publishCmd.Flags().BoolP("dry-run", "n", false, "Validate but do not publish the package.")
	pub_publishCmd.Flags().BoolP("force", "f", false, "Publish without confirmation if there are no errors.")
	pub_publishCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pubCmd.AddCommand(pub_publishCmd)

	carapace.Gen(pub_publishCmd).FlagCompletion(carapace.ActionMap{
		"directory": carapace.ActionDirectories(),
	})
}
