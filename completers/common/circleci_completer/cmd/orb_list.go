package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var orb_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List orbs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(orb_listCmd).Standalone()
	orb_listCmd.PersistentFlags().BoolP("details", "d", false, "output all the commands, executors, and jobs, along with a tree of their parameters")
	orb_listCmd.PersistentFlags().Bool("json", false, "print output as json instead of human-readable")
	orb_listCmd.PersistentFlags().Bool("private", false, "exclusively list private orbs within a namespace")
	orb_listCmd.PersistentFlags().String("sort", "", "one of \"builds\"|\"projects\"|\"orgs\"")
	orb_listCmd.PersistentFlags().BoolP("uncertified", "u", false, "include uncertified orbs")
	orbCmd.AddCommand(orb_listCmd)

	carapace.Gen(orb_listCmd).FlagCompletion(carapace.ActionMap{
		"sort": carapace.ActionValues("builds", "projects", "orgs"),
	})
}
