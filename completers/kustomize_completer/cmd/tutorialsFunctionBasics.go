package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var tutorialsFunctionBasicsCmd = &cobra.Command{
	Use:   "tutorials-function-basics",
	Short: "[Alpha] Tutorials for using functions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tutorialsFunctionBasicsCmd).Standalone()
	rootCmd.AddCommand(tutorialsFunctionBasicsCmd)
}
