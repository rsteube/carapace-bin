package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ssh-agent",
	Short: "OpenSSH authentication agent",
	Long:  "https://linux.die.net/man/1/ssh-agent",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("D", "D", false, "Foreground mode")
	rootCmd.Flags().StringS("E", "E", "", "Specifies the hash algorithm used when displaying key fingerprints")
	rootCmd.Flags().StringS("P", "P", "", "Specify a pattern-list of acceptable paths for PKCS#11 provider and FIDO authenticator")
	rootCmd.Flags().StringS("a", "a", "", "Bind the agent to the UNIX-domain socket bind_address")
	rootCmd.Flags().BoolS("c", "c", false, "Generate C-shell commands on stdout")
	rootCmd.Flags().BoolS("d", "d", false, "Debug mode")
	rootCmd.Flags().BoolS("k", "k", false, "Kill the current agent")
	rootCmd.Flags().BoolS("s", "s", false, "Generate Bourne shell commands on stdout")
	rootCmd.Flags().StringS("t", "t", "", "Set a default value for the maximum lifetime of identities added to the agent")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"E": carapace.ActionValues("md5", "sha256"),
		"a": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
