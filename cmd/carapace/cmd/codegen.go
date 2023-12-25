package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var codegenCmd = &cobra.Command{
	Use:   "--codegen [spec]",
	Short: "generate code for spec file",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		scrape(args[0])
	},
}

func init() {
	carapace.Gen(codegenCmd).Standalone()

	carapace.Gen(codegenCmd).PositionalCompletion(
		carapace.ActionFiles(".yaml"),
	)
}
