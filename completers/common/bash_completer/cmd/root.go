package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bash",
	Short: "GNU Bourne-Again SHell",
	Long:  "https://www.gnu.org/software/bash/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("c", "c", false, "commands are read from the first non-option argument command_string")
	rootCmd.Flags().Bool("debugger", false, "Arrange for the debugger profile to be executed before the shell  starts.")
	rootCmd.Flags().Bool("dump-po-strings", false, "Equivalent to -D, but the output is in the GNU gettext po file format")
	rootCmd.Flags().BoolP("dump-strings", "D", false, "A  list of all double-quoted strings preceded by $ is printed on the standard output.")
	rootCmd.Flags().BoolS("i", "i", false, "If the -i option is present, the shell is interactive.")
	rootCmd.Flags().String("init-file", "", "")
	rootCmd.Flags().BoolP("login", "l", false, "act as if  invoked as a login shell")
	rootCmd.Flags().Bool("noediting", false, "do not use the GNU readline library to read command lines")
	rootCmd.Flags().Bool("noprofile", false, "skip profile initialization files")
	rootCmd.Flags().Bool("norc", false, "skip ~/.bashrc")
	rootCmd.Flags().Bool("posix", false, "match posix behavior")
	rootCmd.Flags().BoolS("r", "r", false, "restricted shell")
	rootCmd.Flags().String("rcfile", "", "execute  commands  from  file instead of the default")
	rootCmd.Flags().Bool("restricted", false, "The shell becomes restricted.")
	rootCmd.Flags().BoolS("s", "s", false, "then commands are read from the standard input")
	rootCmd.Flags().BoolP("verbose", "v", false, "verbose mode")
	rootCmd.Flags().Bool("version", false, "Show version information ")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"init-file": carapace.ActionFiles(),
		"rcfile":    carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("c").Changed || rootCmd.Flag("s").Changed {
				return carapace.ActionValues()
			} else {
				return carapace.ActionFiles()
			}
		}),
	)
}
