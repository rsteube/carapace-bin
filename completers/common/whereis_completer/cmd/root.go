package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "whereis",
	Short: "Locate the binary, source, and manual-page files for a command",
	Long:  "https://linux.die.net/man/1/whereis",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("B", "B", false, "define binaries lookup path")
	rootCmd.Flags().BoolS("M", "M", false, "define man and info lookup path")
	rootCmd.Flags().BoolS("S", "S", false, "define sources lookup path")
	rootCmd.Flags().BoolS("b", "b", false, "search only for binaries")
	rootCmd.Flags().BoolS("f", "f", false, "terminate <dirs> argument list")
	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().BoolS("l", "l", false, "output effective lookup paths")
	rootCmd.Flags().BoolS("m", "m", false, "search only for manuals and infos")
	rootCmd.Flags().BoolS("s", "s", false, "search only for sources")
	rootCmd.Flags().BoolS("u", "u", false, "search for unusual entries")
	rootCmd.Flags().BoolP("version", "V", false, "display version")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if !rootCmd.Flag("f").Changed && (rootCmd.Flag("B").Changed || rootCmd.Flag("M").Changed || rootCmd.Flag("S").Changed) {
				return carapace.ActionDirectories()
			}
			return os.ActionPathExecutables()
		}),
	)
}
