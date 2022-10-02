package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chsh",
	Short: "Change your login shell",
	Long:  "https://en.wikipedia.org/wiki/Chsh",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "u", false, "display this help")
	rootCmd.Flags().BoolP("list-shells", "l", false, "print list of shells and exit")
	rootCmd.Flags().StringP("shell", "s", "", "specify login shell")
	rootCmd.Flags().BoolP("version", "v", false, "display version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"shell": os.ActionShells(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		os.ActionUsers(),
	)
}
