package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/cmd/carapace/cmd/completers"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gh",
	Short: "GitHub CLI",
	Long:  "https://cli.github.com/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.PersistentFlags().Bool("help", false, "Show help for command")
	rootCmd.Flags().Bool("version", false, "Show gh version")

	carapace.Gen(rootCmd).PreRun(func(cmd *cobra.Command, args []string) {
		if aliases, err := action.Aliases(); err == nil {
			for key, value := range aliases {
				cmd.Root().AddCommand(&cobra.Command{Use: key, Short: value, Run: func(cmd *cobra.Command, args []string) {}})
			}
		}

		if extensions, err := action.Extensions(); err == nil {
			for _, extension := range extensions {
				extensionCmd := &cobra.Command{
					Use:                extension,
					Short:              completers.Description("gh-" + extension),
					Run:                func(cmd *cobra.Command, args []string) {},
					DisableFlagParsing: true,
				}

				if extensionCmd.Short == "" {
					extensionCmd.Short = "extension"
				}

				carapace.Gen(extensionCmd).PositionalAnyCompletion(
					bridge.ActionCarapaceBin("gh-" + extension),
				)

				rootCmd.AddCommand(extensionCmd)
			}
		}
	})
}
