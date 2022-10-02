package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/conda_completer/cmd/action"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs a list of packages into a specified conda environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().Bool("all", false, "Update all installed packages in the environment.")
	installCmd.Flags().StringP("channel", "c", "", "Additional channel to search for packages.")
	installCmd.Flags().Bool("clobber", false, "Allow clobbering of overlapping file paths within packages, and suppress related warnings.")
	installCmd.Flags().Bool("copy", false, "Install all packages using copies instead of hard- or soft-linking.")
	installCmd.Flags().Bool("dev", false, "Use `sys.executable -m conda` in wrapper scripts instead of CONDA_EXE.")
	installCmd.Flags().Bool("download-only", false, "Solve an environment and ensure package caches are populated")
	installCmd.Flags().BoolP("dry-run", "d", false, "Only display what would have been done.")
	installCmd.Flags().StringArray("file", []string{}, "Read package versions from the given file.")
	installCmd.Flags().Bool("force-reinstall", false, "Ensure that any user-requested package for the current operation is uninstalled and reinstalled")
	installCmd.Flags().Bool("freeze-installed", false, "Do not update or change already-installed dependencies.")
	installCmd.Flags().BoolP("help", "h", false, "Show this help message and exit.")
	installCmd.Flags().BoolP("insecure", "k", false, "Allow conda to perform \"insecure\" SSL connections and transfers.")
	installCmd.Flags().Bool("json", false, "Report all output as json.")
	installCmd.Flags().BoolP("mkdir", "m", false, "Create the environment directory if necessary.")
	installCmd.Flags().StringP("name", "n", "", "Name of environment.")
	installCmd.Flags().Bool("no-channel-priority", false, "Package version takes precedence over channel priority.")
	installCmd.Flags().Bool("no-deps", false, "Do not install, update, remove, or change dependencies.")
	installCmd.Flags().Bool("no-pin", false, "Ignore pinned file.")
	installCmd.Flags().Bool("no-update-deps", false, "Do not update or change already-installed dependencies.")
	installCmd.Flags().Bool("offline", false, "Offline mode.")
	installCmd.Flags().Bool("only-deps", false, "Only install dependencies.")
	installCmd.Flags().Bool("override-channels", false, "Do not search default or .condarc channels.")
	installCmd.Flags().StringP("prefix", "p", "", "Full path to environment location (i.e. prefix).")
	installCmd.Flags().BoolP("quiet", "q", false, "Do not display progress bar.")
	installCmd.Flags().String("repodata-fn", "", "Specify name of repodata on remote server.")
	installCmd.Flags().String("revision", "", "Revert to the specified REVISION.")
	installCmd.Flags().BoolP("satisfied-skip-solve", "S", false, "Exit early and do not run the solver if the requested specs are satisfied.")
	installCmd.Flags().Bool("show-channel-urls", false, "Show channel urls.")
	installCmd.Flags().Bool("strict-channel-priority", false, "Packages in lower priority channels are not considered if a package with the same name appears in a higher priority channel.")
	installCmd.Flags().Bool("update-all", false, "Update all installed packages in the environment.")
	installCmd.Flags().Bool("update-deps", false, "Update dependencies.")
	installCmd.Flags().Bool("update-specs", false, "Update based on provided specifications.")
	installCmd.Flags().BoolP("use-index-cache", "C", false, "Use cache of channel index files, even if it has expired.")
	installCmd.Flags().Bool("use-local", false, "Use locally built packages. Identical to '-c local'.")
	installCmd.Flags().BoolP("verbose", "v", false, "Once for INFO, twice for DEBUG, three times for TRACE.")
	installCmd.Flags().BoolP("yes", "y", false, "Do not ask for confirmation.")
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"file":   carapace.ActionFiles(),
		"name":   action.ActionEnvironments(),
		"prefix": carapace.ActionDirectories(),
	})
}
