package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var dockerEnvCmd = &cobra.Command{
	Use:   "docker-env",
	Short: "Configure environment to use minikube's Docker daemon",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	dockerEnvCmd.Flags().Bool("no-proxy", false, "Add machine IP to NO_PROXY environment variable")
	dockerEnvCmd.Flags().String("shell", "", "Force environment to be configured for a specified shell: [fish, cmd, powershell, tcsh, bash, zsh], default is auto-detect")
	dockerEnvCmd.Flags().Bool("ssh-add", false, "Add SSH identity key to SSH authentication agent")
	dockerEnvCmd.Flags().Bool("ssh-host", false, "Use SSH connection instead of HTTPS (port 2376)")
	dockerEnvCmd.Flags().BoolP("unset", "u", false, "Unset variables instead of setting them")
	rootCmd.AddCommand(dockerEnvCmd)

	carapace.Gen(dockerEnvCmd).FlagCompletion(carapace.ActionMap{
		"shell": carapace.ActionValues("fish", "cmd", "powershell", "tcsh", "bash", "zsh", "auto-detect"),
	})
}
