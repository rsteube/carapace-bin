package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "slides",
	Short: "Terminal based presentation tool",
	Long:  "https://maaslalani.com/slides/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "help for slides")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(".md"),
	)
}
