package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List configs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_lsCmd).Standalone()

	config_lsCmd.Flags().StringP("filter", "f", "", "Filter output based on conditions provided")
	config_lsCmd.Flags().String("format", "", "Pretty-print configs using a Go template")
	config_lsCmd.Flags().BoolP("quiet", "q", false, "Only display IDs")
	configCmd.AddCommand(config_lsCmd)
}
