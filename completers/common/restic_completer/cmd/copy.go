package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/restic_completer/cmd/action"
	"github.com/spf13/cobra"
)

var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy snapshots from one repository to another",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(copyCmd).Standalone()
	copyCmd.Flags().StringArrayP("host", "H", []string{}, "only consider snapshots for this `host`, when no snapshot ID is given (can be specified multiple times)")
	copyCmd.Flags().String("key-hint2", "", "key ID of key to try decrypting the destination repository first (default: $RESTIC_KEY_HINT2)")
	copyCmd.Flags().String("password-command2", "", "shell `command` to obtain the destination repository password from (default: $RESTIC_PASSWORD_COMMAND2)")
	copyCmd.Flags().String("password-file2", "", "`file` to read the destination repository password from (default: $RESTIC_PASSWORD_FILE2)")
	copyCmd.Flags().StringArray("path", []string{}, "only consider snapshots which include this (absolute) `path`, when no snapshot ID is given")
	copyCmd.Flags().String("repo2", "", "destination `repository` to copy snapshots to (default: $RESTIC_REPOSITORY2)")
	copyCmd.Flags().String("repository-file2", "", "`file` from which to read the destination repository location to copy snapshots to (default: $RESTIC_REPOSITORY_FILE2)")
	copyCmd.Flags().StringSlice("tag", []string{}, "only consider snapshots which include this `taglist`, when no snapshot ID is given")
	rootCmd.AddCommand(copyCmd)

	carapace.Gen(copyCmd).FlagCompletion(carapace.ActionMap{
		"host":              action.ActionSnapshotHosts(copyCmd),
		"password-command2": carapace.ActionFiles(),
		"password-file2":    carapace.ActionFiles(),
		"path":              carapace.ActionFiles(),
		"repo2":             action.ActionRepo(),
		"repository-file2":  carapace.ActionFiles(),
		"tag": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionSnapshotTags(copyCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	carapace.Gen(copyCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionSnapshotIDs(copyCmd).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
