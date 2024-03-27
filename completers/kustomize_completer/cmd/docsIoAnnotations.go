package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var docsIoAnnotationsCmd = &cobra.Command{
	Use:   "docs-io-annotations",
	Short: "[Alpha] Documentation for annotations used by io.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docsIoAnnotationsCmd).Standalone()
	rootCmd.AddCommand(docsIoAnnotationsCmd)
}
