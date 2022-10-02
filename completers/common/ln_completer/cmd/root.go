package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ln",
	Short: "make links between files",
	Long:  "https://linux.die.net/man/1/ln",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("b", "b", false, "like --backup but does not accept an argument")
	rootCmd.Flags().String("backup", "", "make a backup of each existing destination file")
	rootCmd.Flags().BoolP("directory", "d", false, "allow the superuser to attempt to hard link directories")
	rootCmd.Flags().BoolP("force", "f", false, "remove existing destination files")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("interactive", "i", false, "prompt whether to remove destinations")
	rootCmd.Flags().BoolP("logical", "L", false, "dereference TARGETs that are symbolic links")
	rootCmd.Flags().BoolP("no-dereference", "n", false, "treat LINK_NAME as a normal file if it is a symbolic link to a directory")
	rootCmd.Flags().BoolP("no-target-directory", "T", false, "treat LINK_NAME as a normal file always")
	rootCmd.Flags().BoolP("physical", "P", false, "make hard links directly to symbolic links")
	rootCmd.Flags().BoolP("relative", "r", false, "create symbolic links relative to link location")
	rootCmd.Flags().StringP("suffix", "S", "", "override the usual backup suffix")
	rootCmd.Flags().BoolP("symbolic", "s", false, "make symbolic links instead of hard links")
	rootCmd.Flags().StringP("target-directory", "t", "", "specify the DIRECTORY in which to create the links")
	rootCmd.Flags().BoolP("verbose", "v", false, "print name of each linked file")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"backup":           carapace.ActionValues("existing", "nil", "none", "off", "numbered", "t", "simple", "never"),
		"target-directory": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
