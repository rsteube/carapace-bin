package cmd

import (
	"github.com/spf13/cobra"
)

var ls_filesCmd = &cobra.Command{
	Use:   "ls-files",
	Short: "Show information about files in the index and the working tree",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	ls_filesCmd.Flags().String("abbrev", "", "use <n> digits to display SHA-1s")
	ls_filesCmd.Flags().BoolP("cached", "c", false, "show cached files in the output (default)")
	ls_filesCmd.Flags().Bool("debug", false, "show debugging data")
	ls_filesCmd.Flags().BoolP("deleted", "d", false, "show deleted files in the output")
	ls_filesCmd.Flags().Bool("directory", false, "show 'other' directories' names only")
	ls_filesCmd.Flags().Bool("empty-directory", false, "don't show empty directories")
	ls_filesCmd.Flags().Bool("eol", false, "show line endings of files")
	ls_filesCmd.Flags().Bool("error-unmatch", false, "if any <file> is not in the index, treat this as an error")
	ls_filesCmd.Flags().BoolP("exclude", "x", false, "<pattern>    skip files matching pattern")
	ls_filesCmd.Flags().BoolP("exclude-from", "X", false, "<file>    exclude patterns are read from <file>")
	ls_filesCmd.Flags().String("exclude-per-directory", "", "read additional per-directory exclude patterns in <file>")
	ls_filesCmd.Flags().Bool("exclude-standard", false, "add the standard git exclusions")
	ls_filesCmd.Flags().BoolS("f", "f", false, "use lowercase letters for 'fsmonitor clean' files")
	ls_filesCmd.Flags().Bool("full-name", false, "make the output relative to the project top directory")
	ls_filesCmd.Flags().BoolP("ignored", "i", false, "show ignored files in the output")
	ls_filesCmd.Flags().BoolP("killed", "k", false, "show files on the filesystem that need to be removed")
	ls_filesCmd.Flags().BoolP("modified", "m", false, "show modified files in the output")
	ls_filesCmd.Flags().BoolP("others", "o", false, "show other files in the output")
	ls_filesCmd.Flags().Bool("recurse-submodules", false, "recurse through submodules")
	ls_filesCmd.Flags().Bool("resolve-undo", false, "show resolve-undo information")
	ls_filesCmd.Flags().BoolP("stage", "s", false, "show staged contents' object name in the output")
	ls_filesCmd.Flags().BoolS("t", "t", false, "identify the file status with tags")
	ls_filesCmd.Flags().BoolP("unmerged", "u", false, "show unmerged files in the output")
	ls_filesCmd.Flags().BoolS("v", "v", false, "use lowercase letters for 'assume unchanged' files")
	ls_filesCmd.Flags().String("with-tree", "", "pretend that paths removed since <tree-ish> are still present")
	ls_filesCmd.Flags().BoolS("z", "z", false, "paths are separated with NUL character")
	rootCmd.AddCommand(ls_filesCmd)
}
