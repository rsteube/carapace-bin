package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "groupdel",
	Short: "delete a group",
	Long:  "https://linux.die.net/man/8/groupdel",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("force", "f", false, "delete group even if it is the primary group of a user")
	rootCmd.Flags().BoolP("help", "h", false, "display this help message and exit")
	rootCmd.Flags().StringP("prefix", "P", "", "prefix directory where are located the /etc/* files")
	rootCmd.Flags().StringP("root", "R", "", "directory to chroot into")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"prefix": carapace.ActionDirectories(),
		"root":   carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		os.ActionGroups(),
	)
}
