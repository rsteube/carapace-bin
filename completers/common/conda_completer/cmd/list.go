package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/conda_completer/cmd/action"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List linked packages in a conda environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().BoolP("canonical", "c", false, "Output canonical names of packages only.")
	listCmd.Flags().Bool("explicit", false, "List explicitly all installed conda packaged with URL")
	listCmd.Flags().BoolP("export", "e", false, "Output requirement string only")
	listCmd.Flags().BoolP("full-name", "f", false, "Only search for full names.")
	listCmd.Flags().BoolP("help", "h", false, "Show this help message and exit.")
	listCmd.Flags().Bool("json", false, "Report all output as json.")
	listCmd.Flags().Bool("md5", false, "Add MD5 hashsum when using --explicit")
	listCmd.Flags().StringP("name", "n", "", "Name of environment.")
	listCmd.Flags().Bool("no-pip", false, "Do not include pip-only installed packages.")
	listCmd.Flags().StringP("prefix", "p", "", "Full path to environment location (i.e. prefix).")
	listCmd.Flags().BoolP("quiet", "q", false, "Do not display progress bar.")
	listCmd.Flags().BoolP("revisions", "r", false, "List the revision history and exit.")
	listCmd.Flags().Bool("show-channel-urls", false, "Show channel urls.")
	listCmd.Flags().BoolP("verbose", "v", false, "Use once for info, twice for debug, three times for trace.")
	rootCmd.AddCommand(listCmd)

	carapace.Gen(listCmd).FlagCompletion(carapace.ActionMap{
		"name":   action.ActionEnvironments(),
		"prefix": carapace.ActionDirectories(),
	})
}
