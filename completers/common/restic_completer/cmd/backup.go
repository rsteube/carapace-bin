package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/restic_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/time"
	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Create a new backup of files and/or directories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupCmd).Standalone()
	backupCmd.Flags().BoolP("dry-run", "n", false, "do not upload or write any data, just show what would be done")
	backupCmd.Flags().StringArrayP("exclude", "e", []string{}, "exclude a `pattern` (can be specified multiple times)")
	backupCmd.Flags().Bool("exclude-caches", false, "excludes cache directories that are marked with a CACHEDIR.TAG file. See https://bford.info/cachedir/ for the Cache Directory Tagging Standard")
	backupCmd.Flags().StringArray("exclude-file", []string{}, "read exclude patterns from a `file` (can be specified multiple times)")
	backupCmd.Flags().StringArray("exclude-if-present", []string{}, "takes `filename[:header]`, exclude contents of directories containing filename (except filename itself) if header of that file is as provided (can be specified multiple times)")
	backupCmd.Flags().String("exclude-larger-than", "", "max `size` of the files to be backed up (allowed suffixes: k/K, m/M, g/G, t/T)")
	backupCmd.Flags().StringArray("files-from", []string{}, "read the files to backup from `file` (can be combined with file args; can be specified multiple times)")
	backupCmd.Flags().StringArray("files-from-raw", []string{}, "read the files to backup from `file` (can be combined with file args; can be specified multiple times)")
	backupCmd.Flags().StringArray("files-from-verbatim", []string{}, "read the files to backup from `file` (can be combined with file args; can be specified multiple times)")
	backupCmd.Flags().BoolP("force", "f", false, "force re-reading the target files/directories (overrides the \"parent\" flag)")
	backupCmd.Flags().StringP("host", "H", "", "set the `hostname` for the snapshot manually. To prevent an expensive rescan use the \"parent\" flag")
	backupCmd.Flags().String("hostname", "", "set the `hostname` for the snapshot manually")
	backupCmd.Flags().StringArray("iexclude", []string{}, "same as --exclude `pattern` but ignores the casing of filenames")
	backupCmd.Flags().StringArray("iexclude-file", []string{}, "same as --exclude-file but ignores casing of `file`names in patterns")
	backupCmd.Flags().Bool("ignore-ctime", false, "ignore ctime changes when checking for modified files")
	backupCmd.Flags().Bool("ignore-inode", false, "ignore inode number changes when checking for modified files")
	backupCmd.Flags().BoolP("one-file-system", "x", false, "exclude other file systems, don't cross filesystem boundaries and subvolumes")
	backupCmd.Flags().String("parent", "", "use this parent `snapshot` (default: last snapshot in the repo that has the same target files/directories)")
	backupCmd.Flags().Bool("stdin", false, "read backup from stdin")
	backupCmd.Flags().String("stdin-filename", "stdin", "`filename` to use when reading from stdin")
	backupCmd.Flags().StringSlice("tag", []string{}, "add `tags` for the new snapshot in the format `tag[,tag,...]` (can be specified multiple times)")
	backupCmd.Flags().String("time", "", "`time` of the backup (ex. '2012-11-01 22:08:41') (default: now)")
	backupCmd.Flags().Bool("with-atime", false, "store the atime for all files and directories")
	rootCmd.AddCommand(backupCmd)

	carapace.Gen(backupCmd).FlagCompletion(carapace.ActionMap{
		"exclude-file":        carapace.ActionFiles(),
		"files-from":          carapace.ActionFiles(),
		"files-from-raw":      carapace.ActionFiles(),
		"files-from-verbatim": carapace.ActionFiles(),
		"host":                action.ActionSnapshotHosts(backupCmd),
		"parent":              action.ActionSnapshotIDs(backupCmd),
		"tag": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionSnapshotTags(backupCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"time": time.ActionDateTime(time.DateTimeOpts{}),
	})

	carapace.Gen(backupCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
