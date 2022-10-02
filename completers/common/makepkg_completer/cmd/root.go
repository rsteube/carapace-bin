package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "makepkg",
	Short: "make packages compatible for use with pacman",
	Long:  "https://wiki.archlinux.org/title/Makepkg",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("allsource", false, "Generate a source-only tarball including downloaded sources")
	rootCmd.Flags().Bool("asdeps", false, "Install packages as non-explicitly installed")
	rootCmd.Flags().Bool("check", false, "Run the check() function in the PKGBUILD")
	rootCmd.Flags().BoolP("clean", "c", false, "Clean up work files after build")
	rootCmd.Flags().String("config", "", "Use an alternate config file (instead of '/etc/makepkg.conf')")
	rootCmd.Flags().BoolP("force", "f", false, "Overwrite existing package")
	rootCmd.Flags().BoolP("geninteg", "g", false, "Generate integrity checks for source files")
	rootCmd.Flags().BoolP("help", "h", false, "Show this help message and exit")
	rootCmd.Flags().Bool("holdver", false, "Do not update VCS sources")
	rootCmd.Flags().BoolP("install", "i", false, "Install package after successful build")
	rootCmd.Flags().String("key", "", "Specify a key to use for gpg signing instead of the default")
	rootCmd.Flags().BoolP("log", "L", false, "Log package build process")
	rootCmd.Flags().Bool("needed", false, "Do not reinstall the targets that are already up to date")
	rootCmd.Flags().Bool("noarchive", false, "Do not create package archive")
	rootCmd.Flags().BoolP("nobuild", "o", false, "Download and extract files only")
	rootCmd.Flags().Bool("nocheck", false, "Do not run the check() function in the PKGBUILD")
	rootCmd.Flags().BoolP("nocolor", "m", false, "Disable colorized output messages")
	rootCmd.Flags().Bool("noconfirm", false, "Do not ask for confirmation when resolving dependencies")
	rootCmd.Flags().BoolP("nodeps", "d", false, "Skip all dependency checks")
	rootCmd.Flags().BoolP("noextract", "e", false, "Do not extract source files (use existing $srcdir/ dir)")
	rootCmd.Flags().Bool("noprepare", false, "Do not run the prepare() function in the PKGBUILD")
	rootCmd.Flags().Bool("noprogressbar", false, "Do not show a progress bar when downloading files")
	rootCmd.Flags().Bool("nosign", false, "Do not create a signature for the package")
	rootCmd.Flags().StringS("p", "p", "", "Use an alternate build script (instead of 'PKGBUILD')")
	rootCmd.Flags().Bool("packagelist", false, "Only list package filepaths that would be produced")
	rootCmd.Flags().Bool("printsrcinfo", false, "Print the generated SRCINFO and exit")
	rootCmd.Flags().BoolP("repackage", "R", false, "Repackage contents of the package without rebuilding")
	rootCmd.Flags().BoolP("rmdeps", "r", false, "Remove installed dependencies after a successful build")
	rootCmd.Flags().Bool("sign", false, "Sign the resulting package with gpg")
	rootCmd.Flags().Bool("skipchecksums", false, "Do not verify checksums of the source files")
	rootCmd.Flags().Bool("skipinteg", false, "Do not perform any verification checks on source files")
	rootCmd.Flags().Bool("skippgpcheck", false, "Do not verify source files with PGP signatures")
	rootCmd.Flags().BoolP("source", "S", false, "Generate a source-only tarball without downloaded sources")
	rootCmd.Flags().BoolP("syncdeps", "s", false, "Install missing dependencies with pacman")
	rootCmd.Flags().Bool("verifysource", false, "Download source files (if needed) and perform integrity checks")
	rootCmd.Flags().BoolP("version", "V", false, "Show version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
		"p":      carapace.ActionFiles(),
	})
}
