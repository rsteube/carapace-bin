package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var listManualDownloadCmd = &cobra.Command{
	Use:   "list-manual-download",
	Short: "list verbs which download with some help from the user",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listManualDownloadCmd).Standalone()

	rootCmd.AddCommand(listManualDownloadCmd)
}
