package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "groupmems",
	Short: "administer members of a user's primary group",
	Long:  "https://linux.die.net/man/8/groupmems",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("add", "a", "", "add username to the members of the group")
	rootCmd.Flags().StringP("delete", "d", "", "remove username from the members of the group")
	rootCmd.Flags().StringP("group", "g", "", "change groupname instead of the user's group")
	rootCmd.Flags().BoolP("help", "h", false, "display this help message and exit")
	rootCmd.Flags().BoolP("list", "l", false, "list the members of the group")
	rootCmd.Flags().BoolP("purge", "p", false, "purge all members from the group")
	rootCmd.Flags().StringP("root", "R", "", "directory to chroot into")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"add":    os.ActionUsers(),
		"delete": os.ActionUsers(),
		"group":  os.ActionGroups(),
		"root":   carapace.ActionFiles(),
	})
}
