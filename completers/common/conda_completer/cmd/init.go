package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize conda for shell interaction",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().Bool("all", false, "Initialize all currently available shells.")
	initCmd.Flags().BoolP("dry-run", "d", false, "Only display what would have been done.")
	initCmd.Flags().BoolP("help", "h", false, "Show this help message and exit.")
	initCmd.Flags().Bool("json", false, "Report all output as json. Suitable for using conda")
	initCmd.Flags().BoolP("quiet", "q", false, "Do not display progress bar.")
	initCmd.Flags().Bool("reverse", false, "Undo past effects of conda init.")
	initCmd.Flags().BoolP("verbose", "v", false, "Use once for info, twice for debug, three times for trace.")
	rootCmd.AddCommand(initCmd)

	carapace.Gen(initCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("bash", "fish", "powershell", "tcsh", "xonsh", "zsh").Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
