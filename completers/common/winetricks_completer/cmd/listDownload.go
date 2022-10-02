package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var listDownloadCmd = &cobra.Command{
	Use:   "list-download",
	Short: "list verbs which download automatically",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listDownloadCmd).Standalone()

	rootCmd.AddCommand(listDownloadCmd)
}
