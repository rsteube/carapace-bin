package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/restic_completer/cmd/action"
	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Modify tags on snapshots",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tagCmd).Standalone()
	tagCmd.Flags().StringSlice("add", []string{}, "`tags` which will be added to the existing tags in the format `tag[,tag,...]` (can be given multiple times)")
	tagCmd.Flags().StringArrayP("host", "H", []string{}, "only consider snapshots for this `host`, when no snapshot ID is given (can be specified multiple times)")
	tagCmd.Flags().StringArray("path", []string{}, "only consider snapshots which include this (absolute) `path`, when no snapshot-ID is given")
	tagCmd.Flags().StringSlice("remove", []string{}, "`tags` which will be removed from the existing tags in the format `tag[,tag,...]` (can be given multiple times)")
	tagCmd.Flags().StringSlice("set", []string{}, "`tags` which will replace the existing tags in the format `tag[,tag,...]` (can be given multiple times)")
	tagCmd.Flags().StringSlice("tag", []string{}, "only consider snapshots which include this `taglist`, when no snapshot-ID is given")
	rootCmd.AddCommand(tagCmd)

	carapace.Gen(tagCmd).FlagCompletion(carapace.ActionMap{
		"add": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionSnapshotTags(tagCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"host": action.ActionSnapshotHosts(tagCmd),
		"remove": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionSnapshotTags(tagCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"set": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionSnapshotTags(tagCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"tag": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionSnapshotTags(tagCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	carapace.Gen(tagCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionSnapshotIDs(tagCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
	)
}
