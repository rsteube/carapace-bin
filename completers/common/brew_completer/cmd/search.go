package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "search",
	Short: "Perform a substring search of cask tokens and formula names",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(Cmd).Standalone()

	Cmd.Flags().Bool("archlinux", false, "Search for text in the given database.")
	Cmd.Flags().String("cask", "", "Search online and locally for casks.")
	Cmd.Flags().Bool("closed", false, "Search for only closed GitHub pull requests.")
	Cmd.Flags().Bool("debian", false, "Search for text in the given database.")
	Cmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	Cmd.Flags().Bool("desc", false, "Search for formulae with a description")
	Cmd.Flags().Bool("fedora", false, "Search for text in the given database.")
	Cmd.Flags().Bool("fink", false, "Search for text in the given database.")
	Cmd.Flags().String("formula", "", "Search online and locally for formulae.")
	Cmd.Flags().BoolP("help", "h", false, "Show this message.")
	Cmd.Flags().Bool("macports", false, "Search for text in the given database.")
	Cmd.Flags().Bool("open", false, "Search for only open GitHub pull requests.")
	Cmd.Flags().Bool("opensuse", false, "Search for text in the given database.")
	Cmd.Flags().Bool("pull-request", false, "Search for GitHub pull requests containing")
	Cmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	Cmd.Flags().Bool("repology", false, "Search for text in the given database.")
	Cmd.Flags().Bool("ubuntu", false, "Search for text in the given database.")
	Cmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(Cmd)
}
