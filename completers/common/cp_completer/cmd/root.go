package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cp",
	Short: "copy files and directories",
	Long:  "https://linux.die.net/man/1/cp",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("H", "H", false, "follow command-line symbolic links in SOURCE")
	rootCmd.Flags().BoolS("R", "R", false, "copy directories recursively")
	rootCmd.Flags().BoolS("Z", "Z", false, "set SELinux security context of destination file to default type")
	rootCmd.Flags().BoolP("archive", "a", false, "same as -dR --preserve=all")
	rootCmd.Flags().Bool("attributes-only", false, "don't copy the file data, just the attributes")
	rootCmd.Flags().BoolS("b", "b", false, "like --backup but does not accept an argument")
	rootCmd.Flags().String("backup", "", "make a backup of each existing destination file")
	rootCmd.Flags().String("context", "", "like -Z, or if CTX is specified then set the SELinux or SMACK security context to CTX")
	rootCmd.Flags().Bool("copy-contents", false, "copy contents of special files when recursive")
	rootCmd.Flags().BoolS("d", "d", false, "same as --no-dereference --preserve=links")
	rootCmd.Flags().BoolP("dereference", "L", false, "always follow symbolic links in SOURCE")
	rootCmd.Flags().BoolP("force", "f", false, "if an existing destination file cannot be opened, remove it and try again (this option is ignored when the -n option is also used)")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("interactive", "i", false, "prompt before overwrite (overrides a previous -n option)")
	rootCmd.Flags().BoolP("link", "l", false, "hard link files instead of copying")
	rootCmd.Flags().BoolP("no-clobber", "n", false, "do not overwrite an existing file (overrides a previous -i option)")
	rootCmd.Flags().BoolP("no-dereference", "P", false, "never follow symbolic links in SOURCE")
	rootCmd.Flags().StringSlice("no-preserve", []string{}, "don't preserve the specified attributes")
	rootCmd.Flags().BoolP("no-target-directory", "T", false, "treat DEST as a normal file")
	rootCmd.Flags().BoolP("one-file-system", "x", false, "stay on this file system")
	rootCmd.Flags().BoolS("p", "p", false, "same as --preserve=mode,ownership,timestamps")
	rootCmd.Flags().Bool("parents", false, "use full source file name under DIRECTORY")
	rootCmd.Flags().StringSlice("preserve", []string{""}, "preserve the specified attributes (default: mode,ownership,timestamps), if possible additional attributes: context, links, xattr, all")
	rootCmd.Flags().BoolP("recursive", "r", false, "copy directories recursively")
	rootCmd.Flags().String("reflink", "", "control clone/CoW copies")
	rootCmd.Flags().Bool("remove-destination", false, "remove each existing destination file before attempting to open it (contrast with --force)")
	rootCmd.Flags().String("sparse", "", "control creation of sparse files.")
	rootCmd.Flags().Bool("strip-trailing-slashes", false, "remove any trailing slashes from each SOURCE argument")
	rootCmd.Flags().StringP("suffix", "S", "", "override the usual backup suffix")
	rootCmd.Flags().BoolP("symbolic-link", "s", false, "make symbolic links instead of copying")
	rootCmd.Flags().StringP("target-directory", "t", "", "copy all SOURCE arguments into DIRECTORY")
	rootCmd.Flags().BoolP("update", "u", false, "copy only when the SOURCE file is newer than the destination file or when the destination file is missing")
	rootCmd.Flags().BoolP("verbose", "v", false, "explain what is being done")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"backup":           carapace.ActionValues("never", "nil", "none", "numbered", "off", "simple", "t"),
		"no-preserve":      carapace.ActionValues("all", "context", "links", "mode", "ownership", "timestamps", "xattr"),
		"preserve":         carapace.ActionValues("all", "context", "links", "mode", "ownership", "timestamps", "xattr"),
		"reflink":          carapace.ActionValues("alway", "auto").StyleF(style.ForKeyword),
		"sparse":           carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"target-directory": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
