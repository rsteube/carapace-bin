package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var duCmd = &cobra.Command{
	Use:   "du",
	Short: "Disk usage",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(duCmd).Standalone()
	duCmd.Flags().String("filter", "", "Provide filter values")
	duCmd.Flags().Bool("verbose", false, "Provide a more verbose output")
	rootCmd.AddCommand(duCmd)
}
