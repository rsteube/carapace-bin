package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var certsCmd = &cobra.Command{
	Use:   "certs",
	Short: "Commands related to handling kubernetes certificates",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(certsCmd).Standalone()
	rootCmd.AddCommand(certsCmd)
}
