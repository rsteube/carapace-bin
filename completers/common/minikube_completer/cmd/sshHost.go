package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/minikube_completer/cmd/action"
	"github.com/spf13/cobra"
)

var sshHostCmd = &cobra.Command{
	Use:   "ssh-host",
	Short: "Retrieve the ssh host key of the specified node",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	sshHostCmd.Flags().Bool("append-known", false, "Add host key to SSH known_hosts file")
	sshHostCmd.Flags().StringP("node", "n", "", "The node to ssh into. Defaults to the primary control plane.")
	rootCmd.AddCommand(sshHostCmd)

	carapace.Gen(sshHostCmd).FlagCompletion(carapace.ActionMap{
		"node": action.ActionNodes(),
	})
}
