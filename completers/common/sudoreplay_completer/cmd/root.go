package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sudoreplay",
	Short: "replay sudo session logs",
	Long:  "https://linux.die.net/man/8/sudoreplay",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("directory", "d", "", "specify directory for session logs")
	rootCmd.Flags().StringP("filter", "f", "", "specify which I/O type(s) to display")
	rootCmd.Flags().BoolP("help", "h", false, "display help message and exit")
	rootCmd.Flags().BoolP("list", "l", false, "list available session IDs, with optional expression")
	rootCmd.Flags().StringP("max-wait", "m", "", "max number of seconds to wait between events")
	rootCmd.Flags().BoolP("no-resize", "R", false, "do not attempt to re-size the terminal")
	rootCmd.Flags().BoolP("non-interactive", "n", false, "no prompts, session is sent to the standard output")
	rootCmd.Flags().StringP("speed", "s", "", "speed up or slow down output")
	rootCmd.Flags().BoolP("suspend-wait", "S", false, "wait while the command was suspended")
	rootCmd.Flags().BoolP("version", "V", false, "display version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"directory": carapace.ActionDirectories(),
		"filter": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("stdin", "stdout", "stderr", "ttyin", "ttyout").Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	// TODO positional completion
}
