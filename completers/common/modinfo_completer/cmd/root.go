package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "modinfo",
	Short: "Show information about a Linux Kernel module",
	Long:  "https://linux.die.net/man/8/modinfo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("author", "a", false, "Print only 'author'")
	rootCmd.Flags().StringP("basedir", "b", "", "Use DIR as filesystem root for /lib/modules")
	rootCmd.Flags().BoolP("description", "d", false, "Print only 'description'")
	rootCmd.Flags().StringP("field", "F", "", "Print only provided FIELD")
	rootCmd.Flags().BoolP("filename", "n", false, "Print only 'filename'")
	rootCmd.Flags().BoolP("help", "h", false, "Show this help")
	rootCmd.Flags().BoolP("license", "l", false, "Print only 'license'")
	rootCmd.Flags().BoolP("null", "0", false, "Use \\0 instead of \\n")
	rootCmd.Flags().BoolP("parameters", "p", false, "Print only 'parm'")
	rootCmd.Flags().StringP("set-version", "k", "", "Use VERSION instead of `uname -r`")
	rootCmd.Flags().BoolP("version", "V", false, "Show version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"basedir": carapace.ActionDirectories(),
		"field":   carapace.ActionValues("author", "description", "license", "parm", "filename"),
		"set-version": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return os.ActionKernelReleases(rootCmd.Flag("basedir").Value.String())
		}),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return os.ActionKernelModules(os.KernelModulesOpts{Basedir: rootCmd.Flag("basedir").Value.String(), Release: rootCmd.Flag("set-version").Value.String()})
		}),
	)
}
