package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Prints the full starship prompt",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(promptCmd).Standalone()

	promptCmd.Flags().StringP("cmd-duration", "d", "", "The execution duration of the last command, in milliseconds")
	promptCmd.Flags().BoolP("help", "h", false, "Prints help information")
	promptCmd.Flags().StringP("jobs", "j", "", "The number of currently running jobs")
	promptCmd.Flags().StringP("keymap", "k", "", "The keymap of fish/zsh")
	promptCmd.Flags().StringP("path", "p", "", "The path that the prompt should render for")
	promptCmd.Flags().StringP("status", "s", "", "The status code of the previously run command")
	promptCmd.Flags().BoolP("version", "V", false, "Prints version information")
	rootCmd.AddCommand(promptCmd)

	carapace.Gen(promptCmd).FlagCompletion(carapace.ActionMap{
		"path": carapace.ActionDirectories(),
	})
}
