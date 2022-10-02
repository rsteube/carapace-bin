package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var killSessionCmd = &cobra.Command{
	Use:   "kill-session",
	Short: "destroy a given session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(killSessionCmd).Standalone()

	killSessionCmd.Flags().BoolS("C", "C", false, "clear alerts (bell, activity, silence) in all windows linked to the session")
	killSessionCmd.Flags().BoolS("a", "a", false, "kill all session except the one specified by -t")
	killSessionCmd.Flags().StringS("t", "t", "", "specify target session")

	rootCmd.AddCommand(killSessionCmd)

	carapace.Gen(killSessionCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionSessions(),
	})
}
