package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sudoedit",
	Short: "edit files as another user",
	Long:  "https://linux.die.net/man/8/sudoedit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("askpass", "A", false, "use a helper program for password prompting")
	rootCmd.Flags().BoolP("bell", "B", false, "ring bell when prompting")
	rootCmd.Flags().StringP("chdir", "D", "", "change the working directory before running command")
	rootCmd.Flags().StringP("chroot", "R", "", "change the root directory before running command")
	rootCmd.Flags().StringP("close-from", "C", "", "close all file descriptors >= num")
	rootCmd.Flags().StringP("group", "g", "", "run command as the specified group name or ID")
	rootCmd.Flags().Bool("help", false, "display help message and exit")
	rootCmd.Flags().StringP("host", "h", "", "run command on host (if supported by plugin)")
	rootCmd.Flags().BoolP("non-interactive", "n", false, "non-interactive mode, no prompts are used")
	rootCmd.Flags().StringP("prompt", "p", "", "use the specified password prompt")
	rootCmd.Flags().BoolP("reset-timestamp", "k", false, "invalidate timestamp file")
	rootCmd.Flags().BoolP("stdin", "S", false, "read password from standard input")
	rootCmd.Flags().StringP("user", "u", "", "run command (or edit file) as specified user name or ID")
	rootCmd.Flags().BoolP("version", "V", false, "display version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"chdir":  carapace.ActionDirectories(),
		"chroot": carapace.ActionDirectories(),
		"group":  os.ActionGroups(),
		"host":   net.ActionHosts(),
		"user":   os.ActionUsers(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
