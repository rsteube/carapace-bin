package cmd

import (
	"encoding/json"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tldr",
	Short: "Python command line client for tldr",
	Long:  "https://github.com/tldr-pages/tldr-python-client",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("color", "c", false, "Override color stripping")
	rootCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.Flags().StringP("language", "L", "", "Override the default language")
	rootCmd.Flags().BoolP("list", "l", false, "List all available commands for operating system")
	rootCmd.Flags().StringP("platform", "p", "", "Override the operating system")
	rootCmd.Flags().BoolP("render", "r", false, "Render local markdown files")
	rootCmd.Flags().StringP("source", "s", "", "Override the default page source")
	rootCmd.Flags().BoolP("update_cache", "u", false, "Update the local cache of pages and exit")
	rootCmd.Flags().BoolP("version", "v", false, "show program's version number and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"platform": carapace.ActionValues("linux", "osx", "sunos", "windows", "common"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		ActionCommands(),
	)

}

func ActionCommands() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("tldr", "--list")(func(output []byte) carapace.Action {
			var commands []string
			fixedOutput := strings.Replace(string(output), `'`, `"`, -1)
			if err := json.Unmarshal([]byte(fixedOutput), &commands); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				if len(commands) == 0 {
					return carapace.ActionMessage("execute `tldr -u` first")
				} else {
					return carapace.ActionValues(commands...)
				}
			}
		})
	})
}
