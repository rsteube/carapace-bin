package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var attachSessionCmd = &cobra.Command{
	Use:   "attach-session",
	Short: "attach or switch to a session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(attachSessionCmd).Standalone()

	attachSessionCmd.Flags().BoolS("E", "E", false, "don't apply update-environment option")
	attachSessionCmd.Flags().StringS("c", "c", "", "specify working directory for the session")
	attachSessionCmd.Flags().BoolS("d", "d", false, "detach other clients attached to target session")
	attachSessionCmd.Flags().StringS("f", "f", "", "set client flags")
	attachSessionCmd.Flags().BoolS("r", "r", false, "put the client into read-only mode")
	attachSessionCmd.Flags().StringS("t", "t", "", "specify target session")
	attachSessionCmd.Flags().BoolS("x", "x", false, "with -d, send SIGHUP to the parent of the attached client")
	rootCmd.AddCommand(attachSessionCmd)

	carapace.Gen(attachSessionCmd).FlagCompletion(carapace.ActionMap{
		"c": carapace.ActionDirectories(),
		"f": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return tmux.ActionClientFlags().Invoke(c).Filter(c.Parts).ToA()
		}),
		"t": tmux.ActionSessions(),
	})
}
