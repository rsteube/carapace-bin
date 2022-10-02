package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var fmtCmd = &cobra.Command{
	Use:   "fmt",
	Short: "Rewrite waypoint.hcl configuration to a canonical format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fmtCmd).Standalone()

	fmtCmd.Flags().Bool("write", false, "Overwrite the input file.")
	addGlobalOptions(fmtCmd)
	rootCmd.AddCommand(fmtCmd)
}
