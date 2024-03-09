package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:     "run [flags] COMMAND",
	Short:   "Run a command in a Dagger session",
	GroupID: "exec",
	Aliases: []string{"r"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()
	runCmd.Flags().SetInterspersed(false)

	runCmd.Flags().String("cleanup-timeout", "", "max duration to wait between SIGTERM and SIGKILL on interrupt")
	runCmd.Flags().Bool("focus", false, "Only show output for focused commands.")
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
