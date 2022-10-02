package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "termux-apt-repo",
	Short: "Create a repository with deb files",
	Long:  "https://github.com/termux/termux-apt-repo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.Flags().Bool("use-hard-links", false, "use hard links instead of copying deb files. Will not work")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionDirectories(),
		carapace.ActionDirectories(),
	)
}
