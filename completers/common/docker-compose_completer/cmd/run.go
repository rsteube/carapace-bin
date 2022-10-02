package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a one-off command on a service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()
	runCmd.Flags().BoolP("detach", "d", false, "Run container in background and print container ID")
	runCmd.Flags().String("entrypoint", "", "Override the entrypoint of the image")
	runCmd.Flags().StringArrayP("env", "e", []string{}, "Set environment variables")
	runCmd.Flags().StringArrayP("label", "l", []string{}, "Add or override a label")
	runCmd.Flags().String("name", "", " Assign a name to the container")
	runCmd.Flags().BoolP("no-TTY", "T", false, "Disable pseudo-noTty allocation. By default docker compose run allocates a TTY")
	runCmd.Flags().Bool("no-deps", false, "Don't start linked services.")
	runCmd.Flags().StringArrayP("publish", "p", []string{}, "Publish a container's port(s) to the host.")
	runCmd.Flags().Bool("quiet-pull", false, "Pull without printing progress information.")
	runCmd.Flags().Bool("rm", false, "Automatically remove the container when it exits")
	runCmd.Flags().Bool("service-ports", false, "Run command with the service's ports enabled and mapped to the host.")
	runCmd.Flags().Bool("use-aliases", false, "Use the service's network useAliases in the network(s) the container connects to.")
	runCmd.Flags().StringP("user", "u", "", "Run as specified username or uid")
	runCmd.Flags().StringArrayP("volume", "v", []string{}, "Bind mount a volume.")
	runCmd.Flags().StringP("workdir", "w", "", "Working directory inside the container")
	rootCmd.AddCommand(runCmd)

	// TODO flag completion
	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"volume": action.ActionVolumes(runCmd),
	})

	carapace.Gen(runCmd).PositionalCompletion(
		action.ActionServices(runCmd),
	)
}
