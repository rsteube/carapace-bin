package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "usermod",
	Short: "modify a user account",
	Long:  "https://linux.die.net/man/8/usermod",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("add-subgids", "w", "", "add range of subordinate gids")
	rootCmd.Flags().StringP("add-subuids", "v", "", "add range of subordinate uids")
	rootCmd.Flags().BoolP("append", "a", false, "append the user to the supplemental GROUPS")
	rootCmd.Flags().BoolP("badnames", "b", false, "allow bad names")
	rootCmd.Flags().StringP("comment", "c", "", "new value of the GECOS field")
	rootCmd.Flags().StringP("del-subgids", "W", "", "remove range of subordinate gids")
	rootCmd.Flags().StringP("del-subuids", "V", "", "remove range of subordinate uids")
	rootCmd.Flags().StringP("expiredate", "e", "", "set account expiration date to EXPIRE_DATE")
	rootCmd.Flags().StringP("gid", "g", "", "force use GROUP as new primary group")
	rootCmd.Flags().StringP("groups", "G", "", "new list of supplementary GROUPS")
	rootCmd.Flags().BoolP("help", "h", false, "display this help message and exit")
	rootCmd.Flags().StringP("home", "d", "", "new home directory for the user account")
	rootCmd.Flags().StringP("inactive", "f", "", "set password inactive after expiration")
	rootCmd.Flags().BoolP("lock", "L", false, "lock the user account")
	rootCmd.Flags().StringP("login", "l", "", "new value of the login name")
	rootCmd.Flags().BoolP("move-home", "m", false, "move contents of the home directory to the")
	rootCmd.Flags().BoolP("non-unique", "o", false, "allow using duplicate (non-unique) UID")
	rootCmd.Flags().StringP("password", "p", "", "use encrypted password for the new password")
	rootCmd.Flags().StringP("prefix", "P", "", "prefix directory where are located the /etc/* files")
	rootCmd.Flags().StringP("root", "R", "", "directory to chroot into")
	rootCmd.Flags().StringP("shell", "s", "", "new login shell for the user account")
	rootCmd.Flags().StringP("uid", "u", "", "new UID for the user account")
	rootCmd.Flags().BoolP("unlock", "U", false, "unlock the user account")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"gid": os.ActionGroups(),
		"groups": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return os.ActionGroups().Invoke(c).Filter(c.Parts).ToA()
		}),
		"home":   carapace.ActionDirectories(),
		"prefix": carapace.ActionDirectories(),
		"root":   carapace.ActionDirectories(),
		"shell":  os.ActionShells(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		os.ActionUsers(),
	)
}
