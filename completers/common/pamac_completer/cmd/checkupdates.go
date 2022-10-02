package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var checkupdatesCmd = &cobra.Command{
	Use:   "checkupdates",
	Short: "Safely check for updates",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkupdatesCmd).Standalone()

	checkupdatesCmd.Flags().BoolP("aur", "a", false, "also check updates in AUR")
	checkupdatesCmd.Flags().String("builddir", "", "build directory (use with --devel)")
	checkupdatesCmd.Flags().Bool("devel", false, "also check development packages updates (use with --aur)")
	checkupdatesCmd.Flags().Bool("no-aur", false, "do not check updates in AUR")
	checkupdatesCmd.Flags().Bool("no-devel", false, "do not check development packages updates")
	checkupdatesCmd.Flags().BoolP("quiet", "q", false, "only print one line per update")
	rootCmd.AddCommand(checkupdatesCmd)

	carapace.Gen(checkupdatesCmd).FlagCompletion(carapace.ActionMap{
		"builddir": carapace.ActionDirectories(),
	})
}
