package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Run a command in a running container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_execCmd).Standalone()

	container_execCmd.Flags().BoolP("detach", "d", false, "Detached mode: run command in the background")
	container_execCmd.Flags().String("detach-keys", "", "Override the key sequence for detaching a container")
	container_execCmd.Flags().StringP("env", "e", "", "Set environment variables")
	container_execCmd.Flags().BoolP("interactive", "i", false, "Keep STDIN open even if not attached")
	container_execCmd.Flags().Bool("privileged", false, "Give extended privileges to the command")
	container_execCmd.Flags().BoolP("tty", "t", false, "Allocate a pseudo-TTY")
	container_execCmd.Flags().StringP("user", "u", "", "Username or UID (format: <name|uid>[:<group|gid>])")
	container_execCmd.Flags().StringP("workdir", "w", "", "Working directory inside the container")
	containerCmd.AddCommand(container_execCmd)

	rootAlias(container_execCmd, func(cmd *cobra.Command, isAlias bool) {
		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			"detach-keys": docker.ActionDetachKeys(),
			"user":        os.ActionUserGroup(),
		})

		carapace.Gen(cmd).PositionalCompletion(
			docker.ActionContainers(),
		)
	})
}
