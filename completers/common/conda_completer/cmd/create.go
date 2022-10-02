package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/conda_completer/cmd/action"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new conda environment from a list of specified packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createCmd).Standalone()

	createCmd.Flags().StringP("channel", "c", "", "Additional channel to search for packages.")
	createCmd.Flags().String("clone", "", "Path to (or name of) existing local environment.")
	createCmd.Flags().Bool("copy", false, "Install all packages using copies instead of hard- or soft-linking.")
	createCmd.Flags().Bool("dev", false, "Use `sys.executable -m conda` in wrapper scripts instead of CONDA_EXE.")
	createCmd.Flags().Bool("download-only", false, "Solve an environment and ensure package caches are populated")
	createCmd.Flags().BoolP("dry-run", "d", false, "Only display what would have been done.")
	createCmd.Flags().StringArray("file", []string{}, "Read package versions from the given file.")
	createCmd.Flags().BoolP("help", "h", false, "Show this help message and exit.")
	createCmd.Flags().BoolP("insecure", "k", false, "Allow conda to perform \"insecure\" SSL connections and transfers.")
	createCmd.Flags().Bool("json", false, "Report all output as json.")
	createCmd.Flags().StringP("name", "n", "", "Name of environment.")
	createCmd.Flags().Bool("no-channel-priority", false, "Package version takes precedence over channel priority.")
	createCmd.Flags().Bool("no-default-packages", false, "Ignore create_default_packages in the .condarc file.")
	createCmd.Flags().Bool("no-deps", false, "Do not install, update, remove, or change dependencies.")
	createCmd.Flags().Bool("no-pin", false, "Ignore pinned file.")
	createCmd.Flags().Bool("offline", false, "Offline mode.")
	createCmd.Flags().Bool("only-deps", false, "Only install dependencies.")
	createCmd.Flags().Bool("override-channels", false, "Do not search default or .condarc channels.")
	createCmd.Flags().StringP("prefix", "p", "", "Full path to environment location (i.e. prefix).")
	createCmd.Flags().BoolP("quiet", "q", false, "Do not display progress bar.")
	createCmd.Flags().String("repodata-fn", "", "Specify name of repodata on remote server.")
	createCmd.Flags().Bool("show-channel-urls", false, "Show channel urls.")
	createCmd.Flags().Bool("strict-channel-priority", false, "Packages in lower priority channels are not considered if a package with the same name appears in a higher priority channel.")
	createCmd.Flags().BoolP("use-index-cache", "C", false, "Use cache of channel index files, even if it has expired.")
	createCmd.Flags().Bool("use-local", false, "Use locally built packages. Identical to '-c local'.")
	createCmd.Flags().BoolP("verbose", "v", false, "Once for INFO, twice for DEBUG, three times for TRACE.")
	createCmd.Flags().BoolP("yes", "y", false, "Do not ask for confirmation.")
	rootCmd.AddCommand(createCmd)

	carapace.Gen(createCmd).FlagCompletion(carapace.ActionMap{
		"file":   carapace.ActionFiles(),
		"name":   action.ActionEnvironments(),
		"prefix": carapace.ActionDirectories(),
	})

	carapace.Gen(createCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
