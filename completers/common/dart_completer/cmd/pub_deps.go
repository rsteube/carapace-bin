package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pub_depsCmd = &cobra.Command{
	Use:   "deps",
	Short: "Print package dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_depsCmd).Standalone()

	pub_depsCmd.Flags().Bool("dev", false, "Include dev dependencies.")
	pub_depsCmd.Flags().StringP("directory", "C", "", "Run this in the directory<dir>")
	pub_depsCmd.Flags().Bool("executables", false, "List all available executables.")
	pub_depsCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pub_depsCmd.Flags().Bool("json", false, "Output dependency information in a json format.")
	pub_depsCmd.Flags().Bool("no-dev", false, "Do not include dev dependencies.")
	pub_depsCmd.Flags().StringP("style", "s", "", "How output should be displayed.")

	pubCmd.AddCommand(pub_depsCmd)

	carapace.Gen(pub_depsCmd).FlagCompletion(carapace.ActionMap{
		"style": carapace.ActionValues("compact", "tree", "list"),
	})
}
