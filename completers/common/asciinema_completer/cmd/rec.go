package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var recCmd = &cobra.Command{
	Use:   "rec",
	Short: "Record terminal session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(recCmd).Standalone()

	recCmd.Flags().Bool("append", false, "append to existing recording")
	recCmd.Flags().StringP("command", "c", "", "command to record, defaults to $SHELL")
	recCmd.Flags().StringP("env", "e", "", "list of environment variables to capture")
	recCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	recCmd.Flags().StringP("idle-time-limit", "i", "", "limit recorded idle time to given number of seconds")
	recCmd.Flags().Bool("overwrite", false, "overwrite the file if it already exists")
	recCmd.Flags().BoolP("quiet", "q", false, "be quiet, suppress all notices/warnings (implies -y)")
	recCmd.Flags().Bool("raw", false, "save only raw stdout output")
	recCmd.Flags().Bool("stdin", false, "enable stdin recording, disabled by default")
	recCmd.Flags().StringP("title", "t", "", "title of the asciicast")
	recCmd.Flags().BoolP("yes", "y", false, "answer \"yes\" to all prompts (e.g. upload confirmation)")
	rootCmd.AddCommand(recCmd)

	carapace.Gen(recCmd).FlagCompletion(carapace.ActionMap{
		"command": carapace.Batch(
			carapace.ActionFiles(),
			os.ActionPathExecutables(),
		).ToA(),
		"env": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return os.ActionEnvironmentVariables().Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	carapace.Gen(recCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
