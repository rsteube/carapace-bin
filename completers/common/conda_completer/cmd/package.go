package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/conda_completer/cmd/action"
	"github.com/spf13/cobra"
)

var packageCmd = &cobra.Command{
	Use:   "package",
	Short: "Low-level conda package utility",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(packageCmd).Standalone()

	packageCmd.Flags().BoolP("help", "h", false, "Show this help message and exit.")
	packageCmd.Flags().StringP("name", "n", "", "Name of environment.")
	packageCmd.Flags().String("pkg-build", "", "Package build number of the created package.")
	packageCmd.Flags().String("pkg-name", "", "Package name of the created package.")
	packageCmd.Flags().String("pkg-version", "", "Package version of the created package.")
	packageCmd.Flags().StringP("prefix", "p", "", "Full path to environment location (i.e. prefix).")
	packageCmd.Flags().BoolP("reset", "r", false, "Remove all untracked files and exit.")
	packageCmd.Flags().BoolP("untracked", "u", false, "Display all untracked files and exit.")
	packageCmd.Flags().StringArrayP("which", "w", []string{}, "Given some PATH print which conda package the file came from.")
	rootCmd.AddCommand(packageCmd)

	carapace.Gen(packageCmd).FlagCompletion(carapace.ActionMap{
		"name":   action.ActionEnvironments(),
		"prefix": carapace.ActionDirectories(),
		"which":  carapace.ActionFiles(),
	})
}
