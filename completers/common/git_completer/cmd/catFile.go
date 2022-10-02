package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var catFileCmd = &cobra.Command{
	Use:   "cat-file",
	Short: "Provide content or type and size information for repository objects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(catFileCmd).Standalone()

	catFileCmd.Flags().Bool("allow-unknown-type", false, "allow -s and -t to work with broken/corrupt objects")
	catFileCmd.Flags().String("batch", "", "show info and content of objects fed from the standard input")
	catFileCmd.Flags().Bool("batch-all-objects", false, "show all objects with --batch or --batch-check")
	catFileCmd.Flags().String("batch-check", "", "show info about objects fed from the standard input")
	catFileCmd.Flags().Bool("buffer", false, "buffer --batch output")
	catFileCmd.Flags().BoolS("e", "e", false, "exit with zero when there's no error")
	catFileCmd.Flags().Bool("filters", false, "for blob objects, run filters on object's content")
	catFileCmd.Flags().Bool("follow-symlinks", false, "follow in-tree symlinks (used with --batch or --batch-check)")
	catFileCmd.Flags().BoolS("p", "p", false, "pretty-print object's content")
	catFileCmd.Flags().String("path", "", "use a specific path for --textconv/--filters")
	catFileCmd.Flags().BoolS("s", "s", false, "show object size")
	catFileCmd.Flags().BoolS("t", "t", false, "show object type")
	catFileCmd.Flags().Bool("textconv", false, "for blob objects, run textconv on object's content")
	catFileCmd.Flags().Bool("unordered", false, "do not order --batch-all-objects output")
	rootCmd.AddCommand(catFileCmd)

	// TODO completions
	carapace.Gen(catFileCmd).PositionalCompletion(
		carapace.ActionValues("blob", "tree", "commit", "tag"),
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
