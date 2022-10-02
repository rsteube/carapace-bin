package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var orb_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an orb in the specified namespace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(orb_createCmd).Standalone()
	orb_createCmd.Flags().Bool("integration-testing", false, "Enable test mode to bypass interactive UI.")
	orb_createCmd.Flags().Bool("no-prompt", false, "Disable prompt to bypass interactive UI.")
	orb_createCmd.PersistentFlags().Bool("private", false, "Specify that this orb is for private use within your org, unlisted from the public registry.")
	orbCmd.AddCommand(orb_createCmd)
}
