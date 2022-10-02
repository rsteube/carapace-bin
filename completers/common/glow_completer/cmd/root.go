package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "glow",
	Short: "Render markdown on the CLI, with pizzazz!",
	Long:  "https://github.com/charmbracelet/glow",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().BoolP("all", "a", false, "show system files and directories (TUI-mode only)")
	rootCmd.PersistentFlags().String("config", "", "config file")
	rootCmd.Flags().BoolP("help", "h", false, "help for glow")
	rootCmd.Flags().BoolP("local", "l", false, "show local files only; no network (TUI-mode only)")
	rootCmd.Flags().BoolP("mouse", "m", false, "enable mouse wheel (TUI-mode only)")
	rootCmd.Flags().BoolP("pager", "p", false, "display with pager")
	rootCmd.Flags().StringP("style", "s", "auto", "style name or JSON path")
	rootCmd.Flags().BoolP("version", "v", false, "version for glow")
	rootCmd.Flags().UintP("width", "w", 0, "word-wrap at width")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(".yml", ".yaml"),
		"style": carapace.Batch(
			carapace.ActionValues("dark", "light", "auto"),
			carapace.ActionFiles(".json"),
		).ToA(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionCallback(func(c carapace.Context) carapace.Action {
				if !strings.HasPrefix(c.CallbackValue, "https://") {
					return carapace.ActionFiles(".md")
				}
				return carapace.ActionValues()
			}),
			git.ActionRepositorySearch(),
		).ToA(),
	)
}
