package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var listMachinesCmd = &cobra.Command{
	Use:   "list-machines",
	Short: "List local containers and host",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listMachinesCmd).Standalone()

	rootCmd.AddCommand(listMachinesCmd)
}
