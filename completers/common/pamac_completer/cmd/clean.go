package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean packages cache or build files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanCmd).Standalone()

	cleanCmd.Flags().BoolP("build-files", "b", false, "remove all build files")
	cleanCmd.Flags().BoolP("dry-run", "d", false, "do not remove files, only find candidate packages")
	cleanCmd.Flags().StringP("keep", "k", "", "specify how many versions of each package are kept")
	cleanCmd.Flags().Bool("no-confirm", false, "bypass any and all confirmation messages")
	cleanCmd.Flags().BoolP("uninstalled", "u", false, "only target uninstalled packages")
	cleanCmd.Flags().BoolP("verbose", "v", false, "also display all files names")
	rootCmd.AddCommand(cleanCmd)

	carapace.Gen(cleanCmd).FlagCompletion(carapace.ActionMap{
		"keep": carapace.ActionValues("0", "1", "2", "3", "4", "5"),
	})
}
