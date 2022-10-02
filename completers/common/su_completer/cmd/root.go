package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "su",
	Short: "run a command with substitute user and group ID",
	Long:  "https://linux.die.net/man/1/su",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("command", "c", "", "pass a single command")
	rootCmd.Flags().BoolP("fast", "f", false, "pass -f to the shell")
	rootCmd.Flags().StringP("group", "g", "", "specify the primary group")
	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().BoolP("login", "l", false, "make the shell a login shell")
	rootCmd.Flags().BoolS("m", "m", false, "do not reset environment variables")
	rootCmd.Flags().BoolP("preserve-environment", "p", false, "do not reset environment variables")
	rootCmd.Flags().BoolP("pty", "P", false, "create a new pseudo-terminal")
	rootCmd.Flags().String("session-command", "", "pass a single command and do not create a new session")
	rootCmd.Flags().StringP("shell", "s", "", "run given shell")
	rootCmd.Flags().StringP("supp-group", "G", "", "specify a supplemental group")
	rootCmd.Flags().BoolP("version", "V", false, "display version")
	rootCmd.Flags().StringP("whitelist-environment", "w", "", "don't reset specified variables")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"group":      os.ActionGroups(),
		"shell":      os.ActionShells(),
		"supp-group": os.ActionGroups(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		os.ActionUsers(),
	)
}
