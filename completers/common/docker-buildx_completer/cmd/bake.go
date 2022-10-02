package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var bakeCmd = &cobra.Command{
	Use:   "bake",
	Short: "Build from a file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bakeCmd).Standalone()
	bakeCmd.Flags().StringArrayP("file", "f", []string{}, "Build definition file")
	bakeCmd.Flags().Bool("load", false, "Shorthand for \"--set=*.output=type=docker\"")
	bakeCmd.Flags().String("metadata-file", "", "Write build result metadata to the file")
	bakeCmd.Flags().Bool("no-cache", false, "Do not use cache when building the image")
	bakeCmd.Flags().Bool("print", false, "Print the options without building")
	bakeCmd.Flags().String("progress", "auto", "Set type of progress output (\"auto\", \"plain\", \"tty\"). Use plain to show container output")
	bakeCmd.Flags().Bool("pull", false, "Always attempt to pull a newer version of the image")
	bakeCmd.Flags().Bool("push", false, "Shorthand for \"--set=*.output=type=registry\"")
	bakeCmd.Flags().StringArray("set", []string{}, "Override target value (e.g., \"targetpattern.key=value\")")
	rootCmd.AddCommand(bakeCmd)

	// TODO set completion
	carapace.Gen(bakeCmd).FlagCompletion(carapace.ActionMap{
		"file":          carapace.ActionFiles(),
		"metadata-file": carapace.ActionFiles(),
		"progress":      carapace.ActionValues("auto", "plain", "tty"),
	})

	// TODO positional completion
}
