package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/conda_completer/cmd/action"
	"github.com/spf13/cobra"
)

var compareCmd = &cobra.Command{
	Use:   "compare",
	Short: "Compare packages between conda environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compareCmd).Standalone()

	compareCmd.Flags().BoolP("help", "h", false, "Show this help message and exit.")
	compareCmd.Flags().Bool("json", false, "Report all output as json.")
	compareCmd.Flags().StringP("name", "n", "", "Name of environment.")
	compareCmd.Flags().StringP("prefix", "p", "", "Full path to environment location (i.e. prefix).")
	compareCmd.Flags().BoolP("quiet", "q", false, "Do not display progress bar.")
	compareCmd.Flags().BoolP("verbose", "v", false, "Use once for info, twice for debug, three times for trace.")
	rootCmd.AddCommand(compareCmd)

	carapace.Gen(compareCmd).FlagCompletion(carapace.ActionMap{
		"name":   action.ActionEnvironments(),
		"prefix": carapace.ActionDirectories(),
	})

	carapace.Gen(compareCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
