package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zip",
	Short: "package and compress (archive) files",
	Long:  "https://linux.die.net/man/1/zip",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("0", "0", false, "store only")
	rootCmd.Flags().BoolS("1", "1", false, "compress faster")
	rootCmd.Flags().BoolS("9", "9", false, "compress better")
	rootCmd.Flags().BoolS("@", "@", false, "read names from stdin")
	rootCmd.Flags().BoolS("A", "A", false, "adjust self-extracting exe")
	rootCmd.Flags().BoolS("D", "D", false, "do not add directory entries")
	rootCmd.Flags().BoolS("F", "F", false, "fix zipfile (-FF try harder)")
	rootCmd.Flags().BoolS("J", "J", false, "junk zipfile prefix (unzipsfx)")
	rootCmd.Flags().BoolS("T", "T", false, "test zipfile integrity")
	rootCmd.Flags().BoolS("X", "X", false, "eXclude eXtra file attributes")
	rootCmd.Flags().StringS("b", "b", "", "")
	rootCmd.Flags().BoolS("c", "c", false, "add one-line comments")
	rootCmd.Flags().BoolS("d", "d", false, "delete entries in zipfile")
	rootCmd.Flags().BoolS("e", "e", false, "encrypt")
	rootCmd.Flags().BoolS("f", "f", false, "freshen: only changed files")
	rootCmd.Flags().BoolS("i", "i", false, "include only the following names")
	rootCmd.Flags().BoolS("j", "j", false, "junk (don't record) directory names")
	rootCmd.Flags().BoolS("l", "l", false, "convert LF to CR LF")
	rootCmd.Flags().BoolS("m", "m", false, "move into zipfile (delete OS files)")
	rootCmd.Flags().StringS("n", "n", "", "don't compress these suffixes")
	rootCmd.Flags().BoolS("o", "o", false, "make zipfile as old as latest entry")
	rootCmd.Flags().BoolS("q", "q", false, "quiet operation")
	rootCmd.Flags().BoolS("r", "r", false, "recurse into directories")
	rootCmd.Flags().StringS("t", "t", "", "")
	rootCmd.Flags().BoolS("u", "u", false, "update: only changed or new files")
	rootCmd.Flags().BoolS("v", "v", false, "verbose operation/print version info")
	rootCmd.Flags().BoolS("x", "x", false, "exclude the following names")
	rootCmd.Flags().BoolS("y", "y", false, "store symbolic links as the link instead of the referenced file")
	rootCmd.Flags().BoolS("z", "z", false, "add zipfile comment")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
