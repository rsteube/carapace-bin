package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/conda_completer/cmd/action"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"upgrade"},
	Short:   "Updates conda packages to the latest compatible version",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	updateCmd.Flags().Bool("all", false, "Update all installed packages in the environment.")
	updateCmd.Flags().Bool("clobber", false, "Allow clobbering of overlapping file paths within packages, and suppress related warnings.")
	updateCmd.Flags().Bool("copy", false, "Install all packages using copies instead of hard- or soft-linking.")
	updateCmd.Flags().Bool("download-only", false, "Solve an environment and ensure package caches are populated, but exit prior to unlinking and linking packages into the prefix.")
	updateCmd.Flags().BoolP("dry-run", "d", false, "Only display what would have been done.")
	updateCmd.Flags().String("file", "", "Read package versions from the given file.")
	updateCmd.Flags().Bool("force-reinstall", false, "Ensure that any user-requested package for the current operation is uninstalled and reinstalled, even if that package already exists in the environment.")
	updateCmd.Flags().Bool("freeze-installed", false, "Do not update or change already-installed dependencies.")
	updateCmd.Flags().BoolP("help", "h", false, "Show this help message and exit.")
	updateCmd.Flags().BoolP("insecure", "k", false, "Allow conda to perform \"insecure\" SSL connections and transfers.")
	updateCmd.Flags().Bool("json", false, "Report all output as json. Suitable for using conda programmatically.")
	updateCmd.Flags().StringP("name", "n", "", "Name of environment.")
	updateCmd.Flags().Bool("no-channel-priority", false, "Package version takes precedence over channel priority.")
	updateCmd.Flags().Bool("no-deps", false, "Do not install, update, remove, or change dependencies.")
	updateCmd.Flags().Bool("no-pin", false, "Ignore pinned file.")
	updateCmd.Flags().Bool("no-update-deps", false, "Do not update or change already-installed dependencies.")
	updateCmd.Flags().Bool("offline", false, "Offline mode.")
	updateCmd.Flags().Bool("only-deps", false, "Only install dependencies.")
	updateCmd.Flags().Bool("override-channels", false, "Do not search default or .condarc channels.")
	updateCmd.Flags().BoolP("quiet", "q", false, "Do not display progress bar.")
	updateCmd.Flags().String("repodata-fn", "", "Specify name of repodata on remote server.")
	updateCmd.Flags().BoolP("satisfied-skip-solve", "S", false, "Exit early and do not run the solver if the requested specs are satisfied.")
	updateCmd.Flags().Bool("show-channel-urls", false, "Show channel urls.")
	updateCmd.Flags().Bool("strict-channel-priority", false, "Packages in lower priority channels are not considered if a package with the same name appears in a higher priority channel.")
	updateCmd.Flags().Bool("update-all", false, "Update all installed packages in the environment.")
	updateCmd.Flags().Bool("update-deps", false, "Update dependencies.")
	updateCmd.Flags().Bool("update-specs", false, "Update based on provided specifications.")
	updateCmd.Flags().BoolP("use-index-cache", "C", false, "Use cache of channel index files, even if it has expired.")
	updateCmd.Flags().Bool("use-local", false, "Use locally built packages.")
	updateCmd.Flags().BoolP("verbose", "v", false, "Once for INFO, twice for DEBUG, three times for TRACE.")
	updateCmd.Flags().BoolP("yes", "y", false, "Do not ask for confirmation.")
	rootCmd.AddCommand(updateCmd)

	carapace.Gen(updateCmd).FlagCompletion(carapace.ActionMap{
		"name": action.ActionEnvironments(),
	})

	carapace.Gen(updateCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionPackages(updateCmd).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
