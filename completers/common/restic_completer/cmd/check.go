package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check the repository for errors",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkCmd).Standalone()
	checkCmd.Flags().Bool("check-unused", false, "find unused blobs")
	checkCmd.Flags().Bool("read-data", false, "read all data blobs")
	checkCmd.Flags().String("read-data-subset", "", "read a `subset` of data packs, specified as 'n/t' for specific subset or either 'x%' or 'x.y%' for random subset")
	checkCmd.Flags().Bool("with-cache", false, "use the cache")
	rootCmd.AddCommand(checkCmd)
}
