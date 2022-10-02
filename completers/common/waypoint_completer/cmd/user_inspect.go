package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var user_inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Show details about a single user",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(user_inspectCmd).Standalone()

	user_inspectCmd.Flags().String("username", "", "The user to lookup.")

	addGlobalOptions(user_inspectCmd)

	userCmd.AddCommand(user_inspectCmd)
}
