package module

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

func ActionInvokeVenv() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		cmd := &cobra.Command{
			Use:   "venv",
			Short: "Creates virtual Python environments in one or more target directories",
			Run:   func(cmd *cobra.Command, args []string) {},
		}

		carapace.Gen(cmd).Standalone()

		cmd.Flags().Bool("clear", false, "Delete the contents of the environment directory if it already exists, before environment creation.")
		cmd.Flags().Bool("copies", false, "Try to use copies rather than symlinks, even when symlinks are the default for the platform.")
		cmd.Flags().BoolP("help", "h", false, "show this help message and exit")
		cmd.Flags().String("prompt", "", "Provides an alternative prompt prefix for this environment.")
		cmd.Flags().Bool("symlinks", false, "Try to use symlinks rather than copies, when symlinks are not the default for the platform.")
		cmd.Flags().Bool("system-site-packages", false, "Give the virtual environment access to the system site-packages dir.")
		cmd.Flags().Bool("upgrade", false, "Upgrade the environment directory to use this version of Python, assuming Python has been upgraded in-place.")
		cmd.Flags().Bool("upgrade-deps", false, "Upgrade core dependencies: pip setuptools to the latest version in PyPI")
		cmd.Flags().Bool("without-pip", false, "Skips installing or upgrading pip in the virtual environment (pip is bootstrapped by default)")

		carapace.Gen(cmd).PositionalAnyCompletion(
			carapace.ActionDirectories(),
		)

		return carapace.ActionExecute(cmd)
	})
}
