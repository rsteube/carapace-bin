package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gh-dash",
	Short: "A beautiful CLI dashboard for GitHub",
	Long:  "https://github.com/dlvhdr/gh-dash",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("debug", "debug", false, "passing this flag will allow writing debug output to debug.log")

}
