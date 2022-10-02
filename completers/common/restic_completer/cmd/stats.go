package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/restic_completer/cmd/action"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Scan the repository and show basic statistics",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statsCmd).Standalone()
	statsCmd.Flags().StringArrayP("host", "H", []string{}, "only consider snapshots with the given `host` (can be specified multiple times)")
	statsCmd.Flags().String("mode", "restore-size", "counting mode: restore-size (default), files-by-contents, blobs-per-file or raw-data")
	statsCmd.Flags().StringArray("path", []string{}, "only consider snapshots which include this (absolute) `path` (can be specified multiple times)")
	statsCmd.Flags().StringSlice("tag", []string{}, "only consider snapshots which include this `taglist` in the format `tag[,tag,...]` (can be specified multiple times)")
	rootCmd.AddCommand(statsCmd)

	carapace.Gen(statsCmd).FlagCompletion(carapace.ActionMap{
		"host": action.ActionSnapshotHosts(statsCmd),
		"mode": carapace.ActionValues("restore-size", "files-by-contents", "blobs-per-file", "raw-data"),
		"path": carapace.ActionFiles(),
		"tag": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionSnapshotTags(statsCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	carapace.Gen(statsCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionSnapshotIDs(statsCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
	)
}
