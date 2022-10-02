package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sulogin",
	Short: "single-user login",
	Long:  "https://linux.die.net/man/8/sulogin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("force", "e", false, "examine password files directly if getpwnam(3) fails")
	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().BoolP("login-shell", "p", false, "start a login shell")
	rootCmd.Flags().StringP("timeout", "t", "", "max time to wait for a password (default: no limit)")
	rootCmd.Flags().BoolP("version", "V", false, "display version")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
