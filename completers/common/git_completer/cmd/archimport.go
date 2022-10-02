package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var archimportCmd = &cobra.Command{
	Use:   "archimport",
	Short: "Import a GNU Arch repository into Git",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(archimportCmd).Standalone()

	archimportCmd.Flags().StringS("D", "D", "", "Follow merge ancestry and attempt to import trees that have been merged from.")
	archimportCmd.Flags().BoolS("T", "T", false, "Many tags. Will create a tag for every commit, reflecting the commit name in the Arch repository.")
	archimportCmd.Flags().BoolS("a", "a", false, "Attempt to auto-register archives at http://mirrors.sourcecontrol.net")
	archimportCmd.Flags().BoolS("f", "f", false, "Use the fast patchset import strategy.")
	archimportCmd.Flags().BoolS("h", "h", false, "Display usage.")
	archimportCmd.Flags().BoolS("o", "o", false, "Use this for compatibility with old-style branch names used by earlier versions of git archimport.")
	archimportCmd.Flags().StringS("t", "t", "", "Override the default tempdir.")
	archimportCmd.Flags().BoolS("v", "v", false, "Verbose output.")
	rootCmd.AddCommand(archimportCmd)

	carapace.Gen(archimportCmd).FlagCompletion(carapace.ActionMap{
		"t": carapace.ActionDirectories(),
	})

	// TODO positional completion
}
