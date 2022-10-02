package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pub_outdatedCmd = &cobra.Command{
	Use:   "outdated",
	Short: "Analyze your dependencies to find which ones can be upgraded",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_outdatedCmd).Standalone()

	pub_outdatedCmd.Flags().Bool("color", false, "Color the output.")
	pub_outdatedCmd.Flags().Bool("dependency-overrides", false, "Show resolutions with `dependency_overrides`.")
	pub_outdatedCmd.Flags().Bool("dev-dependencies", false, "Take dev dependencies into account.")
	pub_outdatedCmd.Flags().StringP("directory", "C", "", "Run this in the directory <dir>.")
	pub_outdatedCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pub_outdatedCmd.Flags().Bool("json", false, "Output the results using a json format.")
	pub_outdatedCmd.Flags().String("mode", "", "Highlight versions with PROPERTY.")
	pub_outdatedCmd.Flags().Bool("no-color", false, "Do not color the output.")
	pub_outdatedCmd.Flags().Bool("no-dependency-overrides", false, "Do not show resolutions with `dependency_overrides`.")
	pub_outdatedCmd.Flags().Bool("no-dev-dependencies", false, "Do not take dev dependencies into account.")
	pub_outdatedCmd.Flags().Bool("no-show-all", false, "Do not include dependencies that are already")
	pub_outdatedCmd.Flags().Bool("no-transitive", false, "Do not show transitive dependencies.")
	pub_outdatedCmd.Flags().Bool("show-all", false, "Include dependencies that are already")
	pub_outdatedCmd.Flags().Bool("transitive", false, "Show transitive dependencies.")
	pubCmd.AddCommand(pub_outdatedCmd)

	carapace.Gen(pub_outdatedCmd).FlagCompletion(carapace.ActionMap{
		"directory": carapace.ActionDirectories(),
		"mode":      carapace.ActionValues("outdated", "null-safety"),
	})
}
