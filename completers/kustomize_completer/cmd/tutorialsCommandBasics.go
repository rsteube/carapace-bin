package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var tutorialsCommandBasicsCmd = &cobra.Command{
	Use:   "tutorials-command-basics",
	Short: "[Alpha] Tutorials for using basic config commands.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tutorialsCommandBasicsCmd).Standalone()
	rootCmd.AddCommand(tutorialsCommandBasicsCmd)
}
