package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "Print information about the Pulumi environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aboutCmd).Standalone()
	aboutCmd.PersistentFlags().BoolP("json", "j", false, "Emit output as JSON")
	aboutCmd.PersistentFlags().BoolP("transitive", "t", false, "Include transitive dependencies")
	rootCmd.AddCommand(aboutCmd)
}
