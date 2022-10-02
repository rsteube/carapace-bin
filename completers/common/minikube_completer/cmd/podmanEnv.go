package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var podmanEnvCmd = &cobra.Command{
	Use:   "podman-env",
	Short: "Configure environment to use minikube's Podman service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	podmanEnvCmd.Flags().String("shell", "", "Force environment to be configured for a specified shell: [fish, cmd, powershell, tcsh, bash, zsh], default is auto-detect")
	podmanEnvCmd.Flags().BoolP("unset", "u", false, "Unset variables instead of setting them")
	rootCmd.AddCommand(podmanEnvCmd)

	carapace.Gen(podmanEnvCmd).FlagCompletion(carapace.ActionMap{
		"shell": carapace.ActionValues("fish", "cmd", "powershell", "tcsh", "bash", "zsh", "auto-detect"),
	})
}
