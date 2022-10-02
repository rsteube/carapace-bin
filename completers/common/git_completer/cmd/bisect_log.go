package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var bisect_logCmd = &cobra.Command{
	Use:   "log",
	Short: "show bisect log",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bisect_logCmd).Standalone()

	bisectCmd.AddCommand(bisect_logCmd)
}
