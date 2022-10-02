package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var adduserCmd = &cobra.Command{
	Use:   "adduser",
	Short: "Add a registry user account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(adduserCmd).Standalone()
	adduserCmd.Flags().String("scope", "", "associate with scope")

	rootCmd.AddCommand(adduserCmd)
}
