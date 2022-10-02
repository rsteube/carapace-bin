package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var orb_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new orb.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(orb_initCmd).Standalone()
	orb_initCmd.PersistentFlags().Bool("private", false, "initialize a private orb")
	orbCmd.AddCommand(orb_initCmd)
}
