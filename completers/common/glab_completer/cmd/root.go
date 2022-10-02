package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "glab",
	Short: "A GitLab CLI Tool",
	Long:  "https://glab.readthedocs.io/en/latest/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.PersistentFlags().Bool("help", false, "Show help for command")
	rootCmd.Flags().BoolP("version", "v", false, "show glab version information")

	carapace.Gen(rootCmd).PreRun(func(cmd *cobra.Command, args []string) {
		if aliases, err := action.LoadAliases(); err == nil {
			for key, value := range aliases {
				cmd.Root().AddCommand(&cobra.Command{Use: key, Short: value, Run: func(cmd *cobra.Command, args []string) {}})
			}
		}
	})
}
