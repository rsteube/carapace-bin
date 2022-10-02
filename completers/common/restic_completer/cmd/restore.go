package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/restic_completer/cmd/action"
	"github.com/spf13/cobra"
)

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Extract the data from a snapshot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(restoreCmd).Standalone()
	restoreCmd.Flags().StringArrayP("exclude", "e", []string{}, "exclude a `pattern` (can be specified multiple times)")
	restoreCmd.Flags().StringArrayP("host", "H", []string{}, "only consider snapshots for this host when the snapshot ID is \"latest\" (can be specified multiple times)")
	restoreCmd.Flags().StringArray("iexclude", []string{}, "same as `--exclude` but ignores the casing of filenames")
	restoreCmd.Flags().StringArray("iinclude", []string{}, "same as `--include` but ignores the casing of filenames")
	restoreCmd.Flags().StringArrayP("include", "i", []string{}, "include a `pattern`, exclude everything else (can be specified multiple times)")
	restoreCmd.Flags().StringArray("path", []string{}, "only consider snapshots which include this (absolute) `path` for snapshot ID \"latest\"")
	restoreCmd.Flags().StringSlice("tag", []string{}, "only consider snapshots which include this `taglist` for snapshot ID \"latest\"")
	restoreCmd.Flags().StringP("target", "t", "", "directory to extract data to")
	restoreCmd.Flags().Bool("verify", false, "verify restored files content")
	rootCmd.AddCommand(restoreCmd)

	carapace.Gen(restoreCmd).FlagCompletion(carapace.ActionMap{
		"host": action.ActionSnapshotHosts(restoreCmd),
		"path": carapace.ActionFiles(),
		"tag": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionSnapshotTags(restoreCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"target": carapace.ActionDirectories(),
	})

	carapace.Gen(restoreCmd).PositionalCompletion(
		action.ActionSnapshotIDs(restoreCmd),
	)
}
