package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "modprobe",
	Short: "Add and remove modules from the Linux Kernel",
	Long:  "https://linux.die.net/man/8/modprobe",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "Consider every non-argument to be a module name")
	rootCmd.Flags().StringP("config", "C", "", "Use FILE instead of default search paths")
	rootCmd.Flags().StringP("dirname", "d", "", "Use DIR as filesystem root for /lib/modules")
	rootCmd.Flags().BoolP("dry-run", "n", false, "Do not execute operations, just print out")
	rootCmd.Flags().Bool("dump-modversions", false, "Same as --show-modversions")
	rootCmd.Flags().Bool("first-time", false, "Fail if module already inserted or removed")
	rootCmd.Flags().BoolP("force", "f", false, "Force module insertion or removal.")
	rootCmd.Flags().Bool("force-modversion", false, "Ignore module's version")
	rootCmd.Flags().Bool("force-vermagic", false, "Ignore module's version magic")
	rootCmd.Flags().BoolP("help", "h", false, "show this help")
	rootCmd.Flags().BoolS("ignore", "i", false, "Ignore commands")
	rootCmd.Flags().Bool("ignore-install", false, "Ignore install commands")
	rootCmd.Flags().Bool("ignore-remove", false, "Ignore remove commands")
	rootCmd.Flags().BoolP("quiet", "q", false, "disable messages")
	rootCmd.Flags().BoolP("remove", "r", false, "Remove modules instead of inserting")
	rootCmd.Flags().Bool("remove-dependencies", false, "Also remove modules depending on it")
	rootCmd.Flags().BoolP("resolve-alias", "R", false, "Only lookup and print alias and exit")
	rootCmd.Flags().StringP("set-version", "S", "", "Use VERSION instead of `uname -r`")
	rootCmd.Flags().Bool("show", false, "Same as --dry-run")
	rootCmd.Flags().Bool("show-config", false, "Same as --showconfig")
	rootCmd.Flags().BoolP("show-depends", "D", false, "Only print module dependencies and exit")
	rootCmd.Flags().Bool("show-exports", false, "Only print module exported symbol versions and exit")
	rootCmd.Flags().Bool("show-modversions", false, "Dump module symbol version and exit")
	rootCmd.Flags().BoolP("showconfig", "c", false, "Print out known configuration and exit")
	rootCmd.Flags().BoolP("syslog", "s", false, "print to syslog, not stderr")
	rootCmd.Flags().BoolP("use-blacklist", "b", false, "Apply blacklist to resolved alias.")
	rootCmd.Flags().BoolP("verbose", "v", false, "enables more messages")
	rootCmd.Flags().BoolP("version", "V", false, "show version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config":  carapace.ActionFiles(),
		"dirname": carapace.ActionFiles(),
		"set-version": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return os.ActionKernelReleases(rootCmd.Flag("dirname").Value.String())
		}),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("showconfig").Changed || rootCmd.Flag("show-config").Changed {
				return carapace.ActionValues()
			}

			if rootCmd.Flag("dump-modversions").Changed {
				if len(c.Args) == 0 {
					return carapace.ActionFiles()
				}
				return carapace.ActionValues()
			}

			if rootCmd.Flag("remove").Changed {
				return os.ActionKernelModulesLoaded()
			}
			return os.ActionKernelModules(os.KernelModulesOpts{Basedir: rootCmd.Flag("dirname").Value.String(), Release: rootCmd.Flag("set-version").Value.String()})
		}),
	)
}
