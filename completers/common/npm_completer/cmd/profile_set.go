package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var profile_setCmd = &cobra.Command{
	Use:   "set",
	Short: "set profile value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(profile_setCmd).Standalone()

	profileCmd.AddCommand(profile_setCmd)

	// TODO profile key completion
}
