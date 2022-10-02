package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var issue_board_viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View project issue board.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_board_viewCmd).Standalone()
	issue_boardCmd.AddCommand(issue_board_viewCmd)
}
