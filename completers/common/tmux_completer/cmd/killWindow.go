package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var killWindowCmd = &cobra.Command{
	Use:     "kill-window",
	Aliases: []string{"killw"},
	Short:   "destroy a given window",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(killWindowCmd).Standalone()

	killWindowCmd.Flags().BoolS("a", "a", false, "kill all windows except the one specified by -t")
	killWindowCmd.Flags().StringS("t", "t", "", "specify target window")
	rootCmd.AddCommand(killWindowCmd)

	carapace.Gen(killWindowCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionWindows(),
	})
}
