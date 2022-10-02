package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove unused packages and caches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanCmd).Standalone()

	cleanCmd.Flags().BoolP("all", "a", false, "Remove index cache, lock files, unused cache packages, and tarballs.")
	cleanCmd.Flags().StringArrayS("c", "c", []string{}, "Remove temporary files that could not be deleted earlier due to being in-use.")
	cleanCmd.Flags().BoolP("dry-run", "d", false, "Only display what would have been done.")
	cleanCmd.Flags().BoolP("force-pkgs-dirs", "f", false, "Remove *all* writable package caches.")
	cleanCmd.Flags().BoolP("help", "h", false, "Show this help message and exit.")
	cleanCmd.Flags().BoolP("index-cache", "i", false, "Remove index cache.")
	cleanCmd.Flags().Bool("json", false, "Report all output as json.")
	cleanCmd.Flags().BoolP("packages", "p", false, "Remove unused packages from writable package caches.")
	cleanCmd.Flags().BoolP("quiet", "q", false, "Do not display progress bar.")
	cleanCmd.Flags().BoolP("tarballs", "t", false, "Remove cached package tarballs.")
	cleanCmd.Flags().BoolP("verbose", "v", false, "Can be used multiple times.")
	cleanCmd.Flags().BoolP("yes", "y", false, "Do not ask for confirmation.")
	rootCmd.AddCommand(cleanCmd)

	carapace.Gen(cleanCmd).FlagCompletion(carapace.ActionMap{
		"c": carapace.ActionFiles(),
	})
}
