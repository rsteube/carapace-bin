package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker-buildx_completer/cmd/action"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Set the current builder instance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(useCmd).Standalone()
	useCmd.Flags().Bool("default", false, "Set builder as default for current context")
	useCmd.Flags().Bool("global", false, "Builder persists context changes")
	rootCmd.AddCommand(useCmd)

	carapace.Gen(useCmd).PositionalCompletion(
		action.ActionBuilders(),
	)
}
