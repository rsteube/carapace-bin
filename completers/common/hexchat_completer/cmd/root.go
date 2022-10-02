package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hexchat",
	Short: "IRC Client",
	Long:  "https://hexchat.github.io/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("cfgdir", "d", "", "Use a different config directory")
	rootCmd.Flags().StringP("command", "c", "", "Execute command")
	rootCmd.Flags().BoolP("configdir", "u", false, "Show user config directory")
	rootCmd.Flags().String("display", "", "X display to use")
	rootCmd.Flags().BoolP("existing", "e", false, "Open URL or execute command in an existing HexChat")
	rootCmd.Flags().BoolP("help", "h", false, "Show help options")
	rootCmd.Flags().Bool("help-all", false, "Show all help options")
	rootCmd.Flags().Bool("help-gtk", false, "Show GTK+ Options")
	rootCmd.Flags().String("minimize", "", "Begin minimized. Level 0=Normal 1=Iconified 2=Tray")
	rootCmd.Flags().BoolP("no-auto", "a", false, "Don't auto connect to servers")
	rootCmd.Flags().BoolP("no-plugins", "n", false, "Don't auto load any plugins")
	rootCmd.Flags().BoolP("plugindir", "p", false, "Show plugin/script auto-load directory")
	rootCmd.Flags().BoolP("version", "v", false, "Show version information")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cfgdir":  carapace.ActionDirectories(),
		"command": carapace.ActionFiles(),
		"display": os.ActionDisplays(),
		"minimize": carapace.ActionValuesDescribed(
			"0", "Normal",
			"1", "Iconified",
			"2", "Tray",
		),
	})
}
