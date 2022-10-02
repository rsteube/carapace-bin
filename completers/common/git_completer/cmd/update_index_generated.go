package cmd

import (
	"github.com/spf13/cobra"
)

var update_indexCmd = &cobra.Command{
	Use:   "update-index",
	Short: "Register file contents in the working tree to the index",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	update_indexCmd.Flags().Bool("add", false, "do not ignore new files")
	update_indexCmd.Flags().BoolP("again", "g", false, "only update entries that differ from HEAD")
	update_indexCmd.Flags().Bool("assume-unchanged", false, "mark files as \"not changing\"")
	update_indexCmd.Flags().String("cacheinfo", "", "add the specified entry to the index")
	update_indexCmd.Flags().String("chmod", "", "override the executable bit of the listed files")
	update_indexCmd.Flags().Bool("clear-resolve-undo", false, "(for porcelains) forget saved unresolved conflicts")
	update_indexCmd.Flags().Bool("force-remove", false, "remove named paths even if present in worktree")
	update_indexCmd.Flags().Bool("force-untracked-cache", false, "enable untracked cache without testing the filesystem")
	update_indexCmd.Flags().Bool("force-write-index", false, "write out the index even if is not flagged as changed")
	update_indexCmd.Flags().Bool("fsmonitor", false, "enable or disable file system monitor")
	update_indexCmd.Flags().Bool("fsmonitor-valid", false, "mark files as fsmonitor valid")
	update_indexCmd.Flags().Bool("ignore-missing", false, "ignore files missing from worktree")
	update_indexCmd.Flags().Bool("ignore-skip-worktree-entries", false, "do not touch index-only entries")
	update_indexCmd.Flags().Bool("ignore-submodules", false, "refresh: ignore submodules")
	update_indexCmd.Flags().Bool("index-info", false, "add entries from standard input to the index")
	update_indexCmd.Flags().String("index-version", "", "write index in this format")
	update_indexCmd.Flags().Bool("info-only", false, "add to index only; do not add content to object database")
	update_indexCmd.Flags().Bool("no-assume-unchanged", false, "clear assumed-unchanged bit")
	update_indexCmd.Flags().Bool("no-fsmonitor-valid", false, "clear fsmonitor valid bit")
	update_indexCmd.Flags().Bool("no-skip-worktree", false, "clear skip-worktree bit")
	update_indexCmd.Flags().BoolS("q", "q", false, "continue refresh even when index needs update")
	update_indexCmd.Flags().Bool("really-refresh", false, "like --refresh, but ignore assume-unchanged setting")
	update_indexCmd.Flags().Bool("refresh", false, "refresh stat information")
	update_indexCmd.Flags().Bool("remove", false, "notice files missing from worktree")
	update_indexCmd.Flags().Bool("replace", false, "let files replace directories and vice-versa")
	update_indexCmd.Flags().Bool("skip-worktree", false, "mark files as \"index-only\"")
	update_indexCmd.Flags().Bool("split-index", false, "enable or disable split index")
	update_indexCmd.Flags().Bool("stdin", false, "read list of paths to be updated from standard input")
	update_indexCmd.Flags().Bool("test-untracked-cache", false, "test if the filesystem supports untracked cache")
	update_indexCmd.Flags().Bool("unmerged", false, "refresh even if index contains unmerged entries")
	update_indexCmd.Flags().Bool("unresolve", false, "repopulate stages #2 and #3 for the listed paths")
	update_indexCmd.Flags().Bool("untracked-cache", false, "enable/disable untracked cache")
	update_indexCmd.Flags().Bool("verbose", false, "report actions to standard output")
	update_indexCmd.Flags().BoolS("z", "z", false, "with --stdin: input lines are terminated by null bytes")
	rootCmd.AddCommand(update_indexCmd)
}
