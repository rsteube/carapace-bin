package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/minikube_completer/cmd/action"
	"github.com/spf13/cobra"
)

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Log into the minikube environment (for debugging)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	sshCmd.Flags().Bool("native-ssh", true, "Use native Golang SSH client (default true). Set to 'false' to use the command line 'ssh' command when accessing the docker machine. Useful for the machine drivers when they will not start with 'Waiting for SSH'.")
	sshCmd.Flags().StringP("node", "n", "", "The node to ssh into. Defaults to the primary control plane.")
	rootCmd.AddCommand(sshCmd)

	carapace.Gen(sshCmd).FlagCompletion(carapace.ActionMap{
		"node": action.ActionNodes(),
	})
}
