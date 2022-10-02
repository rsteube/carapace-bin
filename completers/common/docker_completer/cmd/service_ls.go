package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var service_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(service_lsCmd).Standalone()

	service_lsCmd.Flags().StringP("filter", "f", "", "Filter output based on conditions provided")
	service_lsCmd.Flags().String("format", "", "Pretty-print services using a Go template")
	service_lsCmd.Flags().BoolP("quiet", "q", false, "Only display IDs")
	serviceCmd.AddCommand(service_lsCmd)
}
