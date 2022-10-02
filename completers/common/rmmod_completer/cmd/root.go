package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rmmod",
	Short: "Simple program to remove a module from the Linux Kernel",
	Long:  "https://linux.die.net/man/8/rmmod",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("force", "f", false, "forces a module unload and may crash your")
	rootCmd.Flags().BoolP("help", "h", false, "show this help")
	rootCmd.Flags().BoolP("syslog", "s", false, "print to syslog, not stderr")
	rootCmd.Flags().BoolP("verbose", "v", false, "enables more messages")
	rootCmd.Flags().BoolP("version", "V", false, "show version")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		os.ActionKernelModules(os.KernelModulesOpts{Basedir: "", Release: ""}),
	)
}
