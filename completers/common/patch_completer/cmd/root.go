package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "patch",
	Short: "appy a diff file to an original",
	Long:  "https://linux.die.net/man/1/patch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("backup", "b", false, "Back up the original contents of each file.")
	rootCmd.Flags().Bool("backup-if-mismatch", false, "Back up if the patch does not match exactly.")
	rootCmd.Flags().StringP("basename-prefix", "Y", "", "Prepend PREFIX to backup file basenames.")
	rootCmd.Flags().BoolP("batch", "t", false, "Ask no questions; skip bad-Prereq patches; assume reversed.")
	rootCmd.Flags().Bool("binary", false, "Read and write data in binary mode.")
	rootCmd.Flags().BoolP("context", "c", false, "Interpret the patch as a context difference.")
	rootCmd.Flags().StringP("directory", "d", "", "Change the working directory to DIR first.")
	rootCmd.Flags().Bool("dry-run", false, "Do not actually change any files; just print what would happen.")
	rootCmd.Flags().BoolP("ed", "e", false, "Interpret the patch as an ed script.")
	rootCmd.Flags().BoolP("force", "f", false, "Like -t, but ignore bad-Prereq patches, and assume unreversed.")
	rootCmd.Flags().BoolP("forward", "N", false, "Ignore patches that appear to be reversed or already applied.")
	rootCmd.Flags().StringP("fuzz", "F", "", "Set the fuzz factor to LINES for inexact matching.")
	rootCmd.Flags().StringP("get", "g", "", "Get files from RCS etc. if positive; ask if negative.")
	rootCmd.Flags().Bool("help", false, "Output this help.")
	rootCmd.Flags().StringP("ifdef", "D", "", "Make merged if-then-else output using NAME.")
	rootCmd.Flags().BoolP("ignore-whitespace", "l", false, "Ignore white space changes between patch and input.")
	rootCmd.Flags().StringP("input", "i", "", "Read patch from PATCHFILE instead of stdin.")
	rootCmd.Flags().Bool("merge", false, "Merge using conflict markers instead of creating reject files.")
	rootCmd.Flags().Bool("no-backup-if-mismatch", false, "Back up mismatches only if otherwise requested.")
	rootCmd.Flags().BoolP("normal", "n", false, "Interpret the patch as a normal difference.")
	rootCmd.Flags().StringP("output", "o", "", "Output patched files to FILE.")
	rootCmd.Flags().Bool("posix", false, "Conform to the POSIX standard.")
	rootCmd.Flags().StringP("prefix", "B", "", "Prepend PREFIX to backup file names.")
	rootCmd.Flags().Bool("quiet", false, "Work silently unless an error occurs.")
	rootCmd.Flags().String("quoting-style", "", "output file names using quoting style WORD.")
	rootCmd.Flags().String("read-only", "", "How to handle read-only input files")
	rootCmd.Flags().StringP("reject-file", "r", "", "Output rejects to FILE.")
	rootCmd.Flags().String("reject-format", "", "Create 'context' or 'unified' rejects.")
	rootCmd.Flags().BoolP("remove-empty-files", "E", false, "Remove output files that are empty after patching.")
	rootCmd.Flags().BoolP("reverse", "R", false, "Assume patches were created with old and new files swapped.")
	rootCmd.Flags().BoolP("set-time", "T", false, "Likewise, assuming local time.")
	rootCmd.Flags().BoolP("set-utc", "Z", false, "Set times of patched files, assuming diff uses UTC (GMT).")
	rootCmd.Flags().BoolP("silent", "s", false, "Work silently unless an error occurs.")
	rootCmd.Flags().StringP("strip", "p", "", "Strip NUM leading components from file names.")
	rootCmd.Flags().StringP("suffix", "z", "", "Append SUFFIX to backup file names.")
	rootCmd.Flags().BoolP("unified", "u", false, "Interpret the patch as a unified difference.")
	rootCmd.Flags().Bool("verbose", false, "Output extra information about the work being done.")
	rootCmd.Flags().BoolP("version", "v", false, "Output version info.")
	rootCmd.Flags().StringP("version-control", "V", "", "Use STYLE version control.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"directory":       carapace.ActionDirectories(),
		"input":           carapace.ActionFiles(),
		"output":          carapace.ActionFiles(),
		"read-only":       carapace.ActionValues("ignore", "warn", "fail"),
		"reject-file":     carapace.ActionFiles(),
		"reject-format":   carapace.ActionValues("context", "unified"),
		"version-control": carapace.ActionValues("simple", "numbered", "existing"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
