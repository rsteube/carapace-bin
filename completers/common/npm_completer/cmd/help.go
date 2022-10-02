package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Get help on npm",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(helpCmd).Standalone()
	helpCmd.Flags().String("viewer", "", "viewer to use")

	rootCmd.AddCommand(helpCmd)

}
