package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List running compose projects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lsCmd).Standalone()
	lsCmd.Flags().BoolP("all", "a", false, "Show all stopped Compose projects")
	lsCmd.Flags().String("filter", "", "Filter output based on conditions provided.")
	lsCmd.Flags().String("format", "pretty", "Format the output. Values: [pretty | json].")
	lsCmd.Flags().BoolP("quiet", "q", false, "Only display IDs.")
	rootCmd.AddCommand(lsCmd)

	carapace.Gen(lsCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("pretty", "json"),
	})
}
