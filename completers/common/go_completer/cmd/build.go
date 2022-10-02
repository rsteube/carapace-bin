package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/golang"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "compile packages and dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().StringS("o", "o", "", "set output file or directory")
	addBuildFlags(buildCmd)
	rootCmd.AddCommand(buildCmd)

	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"o": carapace.ActionFiles(),
	})
}

func addBuildFlags(cmd *cobra.Command) {
	cmd.Flags().BoolS("a", "a", false, "force rebuilding of packages that are already up-to-date")
	cmd.Flags().BoolS("asan", "asan", false, "enable interoperation with address sanitizer")
	cmd.Flags().StringS("asmflags", "asmflags", "", "arguments to pass on each go tool asm invocation")
	cmd.Flags().StringS("buildmode", "buildmode", "", "build mode to use")
	cmd.Flags().StringS("buildvcs", "buildvcs", "", "whether to stamp binaries with version control information")
	cmd.Flags().StringS("compiler", "compiler", "", "name of compiler to use")
	cmd.Flags().StringS("gccgoflags", "gccgoflags", "", "arguments to pass on each gccgo compiler/linker invocation")
	cmd.Flags().StringS("gcflags", "gcflags", "", "arguments to pass on each go tool compile invocation")
	cmd.Flags().StringS("installsuffix", "installsuffix", "", "a suffix to use in the name of the package installation directory")
	cmd.Flags().StringS("ldflags", "ldflags", "", "arguments to pass on each go tool link invocation")
	cmd.Flags().BoolS("linkshared", "linkshared", false, "build code that will be linked against shared libraries")
	cmd.Flags().StringS("mod", "mod", "", "module download mode to use")
	cmd.Flags().BoolS("modcacherw", "modcacherw", false, "leave newly-created directories in the module cache read-write")
	cmd.Flags().StringS("modfile", "modfile", "", "read and possibly write an alternate go.mod file")
	cmd.Flags().BoolS("msan", "msan", false, "enable interoperation with memory sanitizer")
	cmd.Flags().BoolS("n", "n", false, "print the commands but do not run them")
	cmd.Flags().StringS("overlay", "overlay", "", "read a JSON config file that provides an overlay for build operations")
	cmd.Flags().StringS("p", "p", "", "the number of programs to run in parallel")
	cmd.Flags().StringS("pkgdir", "pkgdir", "", "install and load all packages from dir")
	cmd.Flags().BoolS("race", "race", false, "enable data race detection")
	cmd.Flags().StringS("tags", "tags", "", "a comma-separated list of build tags to consider satisfied during the")
	cmd.Flags().StringS("toolexec", "toolexec", "", "a program to use to invoke toolchain programs like vet and asm")
	cmd.Flags().BoolS("trimpath", "trimpath", false, "remove all file system paths from the resulting executable")
	cmd.Flags().BoolS("v", "v", false, "print the names of packages as they are compiled")
	cmd.Flags().BoolS("work", "work", false, "print the name of the temporary work directory")
	cmd.Flags().BoolS("x", "x", false, "print the commands")

	cmd.Flag("buildvcs").NoOptDefVal = "auto"

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"buildmode": carapace.ActionValues("archive", "c-archive", "c-shared", "default", "shared", "exe", "pie", "plugin"),
		"buildvcs":  carapace.ActionValues("true", "false", "auto").StyleF(style.ForKeyword),
		"compiler":  carapace.ActionValues("gccgo", "gc"),
		"mod":       carapace.ActionValues("readonly", "vendor", "mod"),
		"modfile":   carapace.ActionFiles(".mod"),
		"n":         carapace.ActionValues("1", "2", "3", "4", "5", "6", "7", "8"),
		"pkgdir":    carapace.ActionDirectories(),
		"tags": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return golang.ActionBuildTags().Invoke(c).Filter(c.Parts).ToA()
		}),
	})

}
