package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/conda_completer/cmd/action"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"uninstall"},
	Short:   "Remove a list of packages from a specified conda environment",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeCmd).Standalone()

	removeCmd.Flags().Bool("all", false, "Remove all packages, i.e., the entire environment.")
	removeCmd.Flags().StringP("channel", "c", "", "Additional channel to search for packages.")
	removeCmd.Flags().Bool("dev", false, "Use `sys.executable -m conda` in wrapper scripts instead of CONDA_EXE.")
	removeCmd.Flags().BoolP("dry-run", "d", false, "Only display what would have been done.")
	removeCmd.Flags().Bool("features", false, "Remove features (instead of packages).")
	removeCmd.Flags().Bool("force", false, "Forces removal of a package without removing packages that depend on it.")
	removeCmd.Flags().Bool("force-remove", false, "Forces removal of a package without removing packages that depend on it.")
	removeCmd.Flags().BoolP("help", "h", false, "Show this help message and exit.")
	removeCmd.Flags().BoolP("insecure", "k", false, "Allow conda to perform \"insecure\" SSL connections and transfers.")
	removeCmd.Flags().Bool("json", false, "Report all output as json.")
	removeCmd.Flags().StringP("name", "n", "", "Name of environment.")
	removeCmd.Flags().Bool("no-pin", false, "Ignore pinned file.")
	removeCmd.Flags().Bool("offline", false, "Offline mode. Don't connect to the Internet.")
	removeCmd.Flags().Bool("override-channels", false, "Do not search default or .condarc channels.")
	removeCmd.Flags().StringP("prefix", "p", "", "Full path to environment location (i.e. prefix).")
	removeCmd.Flags().BoolP("quiet", "q", false, "Do not display progress bar.")
	removeCmd.Flags().String("repodata-fn", "", "Specify name of repodata on remote server.")
	removeCmd.Flags().BoolP("use-index-cache", "C", false, "Use cache of channel index files, even if it has expired.")
	removeCmd.Flags().Bool("use-local", false, "Use locally built packages.")
	removeCmd.Flags().BoolP("verbose", "v", false, "Once for INFO, twice for DEBUG, three times for TRACE.")
	removeCmd.Flags().BoolP("yes", "y", false, "Do not ask for confirmation.")
	rootCmd.AddCommand(removeCmd)

	carapace.Gen(removeCmd).FlagCompletion(carapace.ActionMap{
		"name":   action.ActionEnvironments(),
		"prefix": carapace.ActionDirectories(),
	})

	carapace.Gen(removeCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionPackages(removeCmd).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
