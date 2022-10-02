package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var user_modifyCmd = &cobra.Command{
	Use:   "modify",
	Short: "Modify details about a user",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(user_modifyCmd).Standalone()

	user_modifyCmd.Flags().String("display-name", "", "The display name for a user.")
	user_modifyCmd.Flags().String("new-username", "", "Set a new username for this user.")
	user_modifyCmd.Flags().String("username", "", "The user to modify.")

	addGlobalOptions(user_modifyCmd)

	userCmd.AddCommand(user_modifyCmd)
}
