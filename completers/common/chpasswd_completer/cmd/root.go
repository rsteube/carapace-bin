package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chpasswd",
	Short: "update passwords in batch mode",
	Long:  "https://linux.die.net/man/8/chpasswd",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("crypt-method", "c", "", "the crypt method")
	rootCmd.Flags().BoolP("encrypted", "e", false, "supplied passwords are encrypted")
	rootCmd.Flags().BoolP("help", "h", false, "display this help message and exit")
	rootCmd.Flags().BoolP("md5", "m", false, "encrypt the clear text password using the MD5 algorithm")
	rootCmd.Flags().StringP("root", "R", "", "directory to chroot into")
	rootCmd.Flags().StringP("sha-rounds", "s", "", "number of rounds for the SHA or BCRYPT crypt algorithms")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"crypt-method": carapace.ActionValues("DES", "MD5", "NONE", "SHA256", "SHA512"),
		"root":         carapace.ActionDirectories(),
	})
}
