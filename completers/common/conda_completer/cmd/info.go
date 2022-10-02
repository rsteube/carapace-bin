package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display information about current conda install",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()

	infoCmd.Flags().BoolP("all", "a", false, "Show all information.")
	infoCmd.Flags().Bool("base", false, "Display base environment path.")
	infoCmd.Flags().BoolP("envs", "e", false, "List all known conda environments.")
	infoCmd.Flags().BoolP("help", "h", false, "Show this help message and exit.")
	infoCmd.Flags().Bool("json", false, "Report all output as json. Suitable for using conda")
	infoCmd.Flags().BoolP("quiet", "q", false, "Do not display progress bar.")
	infoCmd.Flags().BoolP("system", "s", false, "List environment variables.")
	infoCmd.Flags().Bool("unsafe-channels", false, "Display list of channels with tokens exposed.")
	infoCmd.Flags().BoolP("verbose", "v", false, "Use once for info, twice for debug, three times for")
	rootCmd.AddCommand(infoCmd)
}
