package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "passwd",
	Short: "change user password",
	Long:  "https://linux.die.net/man/1/passwd",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "report password status on all accounts")
	rootCmd.Flags().BoolP("delete", "d", false, "delete the password for the named account")
	rootCmd.Flags().BoolP("expire", "e", false, "force expire the password for the named account")
	rootCmd.Flags().BoolP("help", "h", false, "display this help message and exit")
	rootCmd.Flags().StringP("inactive", "i", "", "set password inactive after expiration to INACTIVE")
	rootCmd.Flags().BoolP("keep-tokens", "k", false, "change password only if expired")
	rootCmd.Flags().BoolP("lock", "l", false, "lock the password of the named account")
	rootCmd.Flags().StringP("maxdays", "x", "", "set maximum number of days before password change to MAX_DAYS")
	rootCmd.Flags().StringP("mindays", "n", "", "set minimum number of days before password change to MIN_DAYS")
	rootCmd.Flags().BoolP("quiet", "q", false, "quiet mode")
	rootCmd.Flags().StringP("repository", "r", "", "change password in REPOSITORY repository")
	rootCmd.Flags().StringP("root", "R", "", "directory to chroot into")
	rootCmd.Flags().BoolP("status", "S", false, "report password status on the named account")
	rootCmd.Flags().BoolP("unlock", "u", false, "unlock the password of the named account")
	rootCmd.Flags().StringP("warndays", "w", "", "set expiration warning days to WARN_DAYS")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"root": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		os.ActionUsers(),
	)
}
