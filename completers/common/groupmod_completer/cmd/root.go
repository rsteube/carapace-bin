package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "groupmod",
	Short: "modify a group definition on the system",
	Long:  "https://linux.die.net/man/8/groupmod",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("gid", "g", "", "change the group ID to GID")
	rootCmd.Flags().BoolP("help", "h", false, "display this help message and exit")
	rootCmd.Flags().StringP("new-name", "n", "", "change the name to NEW_GROUP")
	rootCmd.Flags().BoolP("non-unique", "o", false, "allow to use a duplicate (non-unique) GID")
	rootCmd.Flags().StringP("password", "p", "", "change the password to this (encrypted)")
	rootCmd.Flags().StringP("prefix", "P", "", "prefix directory where are located the /etc/* files")
	rootCmd.Flags().StringP("root", "R", "", "directory to chroot into")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"prefix": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		os.ActionGroups(),
	)
}
