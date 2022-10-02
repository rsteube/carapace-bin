package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var codespace_sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "SSH into a codespace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codespace_sshCmd).Standalone()
	codespace_sshCmd.Flags().StringP("codespace", "c", "", "Name of the codespace")
	codespace_sshCmd.Flags().Bool("config", false, "Write OpenSSH configuration to stdout")
	codespace_sshCmd.Flags().BoolP("debug", "d", false, "Log debug data to a file")
	codespace_sshCmd.Flags().String("debug-file", "", "Path of the file log to")
	codespace_sshCmd.Flags().String("profile", "", "Name of the SSH profile to use")
	codespace_sshCmd.Flags().Int("server-port", 0, "SSH server port number (0 => pick unused)")
	codespace_sshCmd.Flags().Bool("stdio", false, "Proxy sshd connection to stdio")
	codespaceCmd.AddCommand(codespace_sshCmd)

	carapace.Gen(codespace_sshCmd).FlagCompletion(carapace.ActionMap{
		"codespace":  action.ActionCodespaces(),
		"debug-file": carapace.ActionFiles(),
		"profile":    carapace.ActionValues("TODO"), // TODO complete ssh profiles
	})

	// TODO ssh flags when os.Args contains `--`
}
