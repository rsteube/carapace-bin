package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "Create an archive of files from a named tree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(archiveCmd).Standalone()

	archiveCmd.Flags().BoolS("0", "0", false, "set compression level")
	archiveCmd.Flags().BoolS("1", "1", false, "set compression level")
	archiveCmd.Flags().BoolS("2", "2", false, "set compression level")
	archiveCmd.Flags().BoolS("3", "3", false, "set compression level")
	archiveCmd.Flags().BoolS("4", "4", false, "set compression level")
	archiveCmd.Flags().BoolS("5", "5", false, "set compression level")
	archiveCmd.Flags().BoolS("6", "6", false, "set compression level")
	archiveCmd.Flags().BoolS("7", "7", false, "set compression level")
	archiveCmd.Flags().BoolS("8", "8", false, "set compression level")
	archiveCmd.Flags().BoolS("9", "9", false, "set compression level")
	archiveCmd.Flags().String("add-file", "", "add untracked file to archive")
	archiveCmd.Flags().String("exec", "", "path to the remote git-upload-archive command")
	archiveCmd.Flags().String("format", "", "archive format")
	archiveCmd.Flags().BoolP("list", "l", false, "list supported archive formats")
	archiveCmd.Flags().StringP("output", "o", "", "write the archive to this file")
	archiveCmd.Flags().String("prefix", "", "prepend prefix to each pathname in the archive")
	archiveCmd.Flags().String("remote", "", "retrieve the archive from remote repository <repo>")
	archiveCmd.Flags().BoolP("verbose", "v", false, "report archived files on stderr")
	archiveCmd.Flags().Bool("worktree-attributes", false, "read .gitattributes in working directory")
	rootCmd.AddCommand(archiveCmd)

	carapace.Gen(archiveCmd).FlagCompletion(carapace.ActionMap{
		"add-file": carapace.ActionFiles(),
		"format":   carapace.ActionValues("tar", "zip"),
		"output":   carapace.ActionFiles(),
		"remote":   git.ActionRemotes(),
	})

	carapace.Gen(archiveCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)

	carapace.Gen(archiveCmd).PositionalAnyCompletion(
		carapace.ActionDirectories(), // TODO shouldn't thi bee the path from the ref?
	)
}
