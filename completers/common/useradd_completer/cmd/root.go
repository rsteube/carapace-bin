package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "useradd",
	Short: "create a new user or update default new user information",
	Long:  "https://linux.die.net/man/8/useradd",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("badnames", false, "do not check for bad names")
	rootCmd.Flags().StringP("base-dir", "b", "", "base directory for the home directory of the")
	rootCmd.Flags().Bool("btrfs-subvolume-home", false, "use BTRFS subvolume for home directory")
	rootCmd.Flags().StringP("comment", "c", "", "GECOS field of the new account")
	rootCmd.Flags().BoolP("create-home", "m", false, "create the user's home directory")
	rootCmd.Flags().BoolP("defaults", "D", false, "print or change default useradd configuration")
	rootCmd.Flags().StringP("expiredate", "e", "", "expiration date of the new account")
	rootCmd.Flags().StringP("gid", "g", "", "name or ID of the primary group of the new")
	rootCmd.Flags().StringP("groups", "G", "", "list of supplementary groups of the new")
	rootCmd.Flags().BoolP("help", "h", false, "display this help message and exit")
	rootCmd.Flags().StringP("home-dir", "d", "", "home directory of the new account")
	rootCmd.Flags().StringP("inactive", "f", "", "password inactivity period of the new account")
	rootCmd.Flags().StringP("key", "K", "", "override /etc/login.defs defaults")
	rootCmd.Flags().BoolP("no-create-home", "M", false, "do not create the user's home directory")
	rootCmd.Flags().BoolP("no-log-init", "l", false, "do not add the user to the lastlog and")
	rootCmd.Flags().BoolP("no-user-group", "N", false, "do not create a group with the same name as")
	rootCmd.Flags().BoolP("non-unique", "o", false, "allow to create users with duplicate")
	rootCmd.Flags().StringP("password", "p", "", "encrypted password of the new account")
	rootCmd.Flags().StringP("prefix", "P", "", "prefix directory where are located the /etc/* files")
	rootCmd.Flags().StringP("root", "R", "", "directory to chroot into")
	rootCmd.Flags().StringP("shell", "s", "", "login shell of the new account")
	rootCmd.Flags().StringP("skel", "k", "", "use this alternative skeleton directory")
	rootCmd.Flags().BoolP("system", "r", false, "create a system account")
	rootCmd.Flags().StringP("uid", "u", "", "user ID of the new account")
	rootCmd.Flags().BoolP("user-group", "U", false, "create a group with the same name as the user")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"base-dir": carapace.ActionDirectories(),
		"gid":      os.ActionGroups(),
		"groups": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return os.ActionGroups().Invoke(c).Filter(c.Parts).ToA()
		}),
		"home-dir": carapace.ActionDirectories(),
		"prefix":   carapace.ActionDirectories(),
		"root":     carapace.ActionDirectories(),
		"shell":    os.ActionShells(),
		"skel":     carapace.ActionDirectories(),
	})
}
