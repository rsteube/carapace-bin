package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "unzip",
	Short: "list, test and extract compressed files in a ZIP archive",
	Long:  "https://linux.die.net/man/1/unzip",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("C", "C", false, "match filenames case-insensitively")
	rootCmd.Flags().BoolS("K", "K", false, "keep setuid/setgid/tacky permissions")
	rootCmd.Flags().BoolS("L", "L", false, "make (some) names lowercase")
	rootCmd.Flags().BoolS("M", "M", false, "pipe through \"more\" pager")
	rootCmd.Flags().BoolS("T", "T", false, "timestamp archive to latest")
	rootCmd.Flags().BoolS("U", "U", false, "use escapes for all non-ASCII Unicode")
	rootCmd.Flags().BoolS("V", "V", false, "retain VMS version numbers")
	rootCmd.Flags().BoolS("X", "X", false, "restore UID/GID info")
	rootCmd.Flags().BoolS("Z", "Z", false, "ZipInfo mode")
	rootCmd.Flags().BoolS("a", "a", false, "auto-convert any text files")
	rootCmd.Flags().StringS("d", "d", "", "extract files into exdir")
	rootCmd.Flags().BoolS("f", "f", false, "freshen existing files, create none")
	rootCmd.Flags().BoolS("j", "j", false, "junk paths (do not make directories)")
	rootCmd.Flags().BoolS("l", "l", false, "list files (short format)")
	rootCmd.Flags().BoolS("n", "n", false, "never overwrite existing files")
	rootCmd.Flags().BoolS("o", "o", false, "overwrite files WITHOUT prompting")
	rootCmd.Flags().BoolS("p", "p", false, "extract files to pipe, no messages")
	rootCmd.Flags().BoolS("q", "q", false, "quiet mode (-qq => quieter)")
	rootCmd.Flags().BoolS("t", "t", false, "test compressed archive data")
	rootCmd.Flags().BoolS("u", "u", false, "update files, create if necessary")
	rootCmd.Flags().BoolS("v", "v", false, "list verbosely/show version info")
	rootCmd.Flags().StringS("x", "x", "", "exclude files that follow (in xlist)")
	rootCmd.Flags().BoolS("z", "z", false, "display archive comment only")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"d": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return fs.ActionZipFileContents(c.Args[0])
		}),
	)
}
