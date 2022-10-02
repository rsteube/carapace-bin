package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_attachCmd = &cobra.Command{
	Use:   "attach",
	Short: "Attach local standard input, output, and error streams to a running container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_attachCmd).Standalone()

	container_attachCmd.Flags().String("detach-keys", "", "Override the key sequence for detaching a container")
	container_attachCmd.Flags().Bool("no-stdin", false, "Do not attach STDIN")
	container_attachCmd.Flags().Bool("sig-proxy", false, "Proxy all received signals to the process (default true)")
	containerCmd.AddCommand(container_attachCmd)

	rootAlias(container_attachCmd, func(cmd *cobra.Command, isAlias bool) {
		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			"detach-keys": docker.ActionDetachKeys(),
		})

		carapace.Gen(cmd).PositionalCompletion(
			docker.ActionContainers(),
		)
	})
}
