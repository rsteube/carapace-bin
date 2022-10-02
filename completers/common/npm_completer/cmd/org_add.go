package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var org_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Add a new developer to an org",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(org_setCmd).Standalone()

	orgCmd.AddCommand(org_setCmd)

	carapace.Gen(org_setCmd).PositionalCompletion(
		carapace.ActionValues(),
		carapace.ActionValues(),
		carapace.ActionValues("developer", "admin", "owner"),
	)
}
