package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "groupadd",
	Short: "create a new group",
	Long:  "https://linux.die.net/man/8/groupadd",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("force", "f", false, "exit successfully if the group already exists,")
	rootCmd.Flags().StringP("gid", "g", "", "use GID for the new group")
	rootCmd.Flags().BoolP("help", "h", false, "display this help message and exit")
	rootCmd.Flags().StringP("key", "K", "", "override /etc/login.defs defaults")
	rootCmd.Flags().BoolP("non-unique", "o", false, "allow to create groups with duplicate")
	rootCmd.Flags().StringP("password", "p", "", "use this encrypted password for the new group")
	rootCmd.Flags().StringP("prefix", "P", "", "directory prefix")
	rootCmd.Flags().StringP("root", "R", "", "directory to chroot into")
	rootCmd.Flags().BoolP("system", "r", false, "create a system account")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"prefix": carapace.ActionDirectories(),
		"root":   carapace.ActionDirectories(),
	})
}
