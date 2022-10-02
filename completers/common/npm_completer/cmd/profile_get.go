package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var profile_getCmd = &cobra.Command{
	Use:   "get",
	Short: "get profile value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(profile_getCmd).Standalone()

	profileCmd.AddCommand(profile_getCmd)

	// TODO profile key completion
}
