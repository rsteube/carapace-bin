package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "apt-get",
	Short: "APT package handling utility",
	Long:  "https://linux.die.net/man/8/apt-get",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().Bool("allow-downgrades", false, "Continue without prompting if it is doing downgrades.")
	rootCmd.PersistentFlags().Bool("allow-releaseinfo-change", false, "Allow the update command to continue downloading data from a repository which changed its information")
	rootCmd.PersistentFlags().Bool("allow-remove-essential", false, "Force yes")
	rootCmd.PersistentFlags().Bool("allow-unauthenticated", false, "Ignore if packages can't be authenticated and don't prompt about it.")
	rootCmd.PersistentFlags().Bool("arch-only", false, "Only process architecture-dependent build-dependencies.")
	rootCmd.PersistentFlags().Bool("assume-no", false, "Automatic \"no\" to all prompts")
	rootCmd.PersistentFlags().Bool("assume-yes", false, "Automatic yes to prompts")
	rootCmd.PersistentFlags().Bool("auto-remove", false, "acts like running the autoremove command")
	rootCmd.PersistentFlags().Bool("autoremove", false, "acts like running the autoremove command")
	rootCmd.PersistentFlags().Bool("build", false, "Compile source packages after downloading them.")
	rootCmd.PersistentFlags().BoolP("build-profiles", "P", false, "This option controls the activated build profiles")
	rootCmd.PersistentFlags().BoolP("compile", "b", false, "Compile source packages after downloading them.")
	rootCmd.PersistentFlags().BoolP("config-file", "c", false, "Configuration File")
	rootCmd.PersistentFlags().Bool("default-release", false, "This option controls the default input to the policy engine")
	rootCmd.PersistentFlags().Bool("diff-only", false, "Download only the diff file of a source archive.")
	rootCmd.PersistentFlags().BoolP("download-only", "d", false, "Download only")
	rootCmd.PersistentFlags().Bool("dry-run", false, "No action")
	rootCmd.PersistentFlags().Bool("dsc-only", false, "Download only the dsc file of a source archive.")
	rootCmd.PersistentFlags().String("error-on", "", "Fail the update command if any error occured, even a transient one.")
	rootCmd.PersistentFlags().BoolP("fix-broken", "f", false, "Fix")
	rootCmd.PersistentFlags().Bool("fix-missing", false, "Ignore missing packages")
	rootCmd.PersistentFlags().Bool("force-yes", false, "Force yes")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Show a short usage summary.")
	rootCmd.PersistentFlags().BoolP("host-architecture", "a", false, "This option controls the architecture packages are built for")
	rootCmd.PersistentFlags().Bool("ignore-hold", false, "Ignore package holds")
	rootCmd.PersistentFlags().BoolP("ignore-missing", "m", false, "Ignore missing packages")
	rootCmd.PersistentFlags().Bool("indep-only", false, "Only process architecture-independent build-dependencies.")
	rootCmd.PersistentFlags().Bool("install-suggests", false, "Consider suggested packages as a dependency for installing.")
	rootCmd.PersistentFlags().Bool("just-print", false, "No action")
	rootCmd.PersistentFlags().Bool("list-cleanup", false, "Automatically manage the contents of /var/lib/apt/lists")
	rootCmd.PersistentFlags().Bool("no-act", false, "No action")
	rootCmd.PersistentFlags().Bool("no-allow-insecure-repositories", false, "Forbid the update command to acquire unverifiable data from configured sources.")
	rootCmd.PersistentFlags().Bool("no-download", false, "Disables downloading of packages.")
	rootCmd.PersistentFlags().Bool("no-install-recommends", false, "Do not consider recommended packages as a dependency for installing.")
	rootCmd.PersistentFlags().Bool("no-remove", false, "If any packages are to be removed apt-get immediately aborts without prompting.")
	rootCmd.PersistentFlags().Bool("no-show-upgraded", false, "Do not show a list of all packages that are to be upgraded.")
	rootCmd.PersistentFlags().Bool("no-upgrade", false, "Do not upgrade packages")
	rootCmd.PersistentFlags().Bool("only-source", false, "Indicates that the given source names are not to be mapped through the binary table.")
	rootCmd.PersistentFlags().Bool("only-upgrade", false, "Do not install new packages")
	rootCmd.PersistentFlags().BoolP("option", "o", false, "Set a Configuration Option")
	rootCmd.PersistentFlags().Bool("print-uris", false, "Instead of fetching the files to install their URIs are printed.")
	rootCmd.PersistentFlags().Bool("purge", false, "Use purge instead of remove for anything that would be removed.")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "Quiet")
	rootCmd.PersistentFlags().Bool("recon", false, "No action")
	rootCmd.PersistentFlags().Bool("reinstall", false, "Re-install packages that are already installed and at the newest version.")
	rootCmd.PersistentFlags().Bool("show-progress", false, "Show user friendly progress information in the terminal window when packages are installed")
	rootCmd.PersistentFlags().BoolP("simulate", "s", false, "No action")
	rootCmd.PersistentFlags().Bool("tar-only", false, "Download only the tar file of a source archive.")
	rootCmd.PersistentFlags().BoolP("target-release", "t", false, "This option controls the default input to the policy engine")
	rootCmd.PersistentFlags().Bool("trivial-only", false, "Only perform operations that are 'trivial'.")
	rootCmd.PersistentFlags().BoolP("verbose-versions", "V", false, "Show full versions for upgraded and installed packages.")
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Show the program version.")
	rootCmd.PersistentFlags().Bool("with-new-pkgs", false, "Allow installing new packages when used in conjunction with upgrade.")
	rootCmd.PersistentFlags().String("with-source", "", "Adds the given file as a source for metadata.")
	rootCmd.PersistentFlags().BoolP("yes", "y", false, "Automatic yes to prompts")

	rootCmd.Flag("error-on").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"with-source": carapace.ActionFiles(),
	})
}
