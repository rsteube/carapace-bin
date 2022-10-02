package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "install",
	Short: "copy files and set attributes",
	Long:  "https://linux.die.net/man/1/install",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("D", "D", false, "create all leading components of DEST except the last,")
	rootCmd.Flags().BoolS("Z", "Z", false, "set SELinux security context of destination")
	rootCmd.Flags().BoolS("b", "b", false, "like --backup but does not accept an argument")
	rootCmd.Flags().String("backup", "", "make a backup of each existing destination file")
	rootCmd.Flags().BoolS("c", "c", false, "(ignored)")
	rootCmd.Flags().BoolP("compare", "C", false, "compare each pair of source and destination files, and")
	rootCmd.Flags().String("context", "", "like -Z, or if CTX is specified then set the")
	rootCmd.Flags().BoolP("directory", "d", false, "treat all arguments as directory names; create all")
	rootCmd.Flags().StringP("group", "g", "", "set group ownership, instead of process' current group")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().StringP("mode", "m", "", "set permission mode (as in chmod), instead of rwxr-xr-x")
	rootCmd.Flags().BoolP("no-target-directory", "T", false, "treat DEST as a normal file")
	rootCmd.Flags().StringP("owner", "o", "", "set ownership (super-user only)")
	rootCmd.Flags().Bool("preserve-context", false, "preserve SELinux security context")
	rootCmd.Flags().BoolP("preserve-timestamps", "p", false, "apply access/modification times of SOURCE files")
	rootCmd.Flags().BoolP("strip", "s", false, "strip symbol tables")
	rootCmd.Flags().String("strip-program", "", "program used to strip binaries")
	rootCmd.Flags().StringP("suffix", "S", "", "override the usual backup suffix")
	rootCmd.Flags().StringP("target-directory", "t", "", "copy all SOURCE arguments into DIRECTORY")
	rootCmd.Flags().BoolP("verbose", "v", false, "print the name of each directory as it is created")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		// TODO complete mode
		"group":            os.ActionGroups(),
		"owner":            os.ActionUsers(),
		"strip-program":    carapace.ActionDirectories(),
		"target-directory": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionDirectories())
}
