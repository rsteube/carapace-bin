package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rustc",
	Short: "compiler for the Rust programming language",
	Long:  "https://doc.rust-lang.org/rustc/index.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringArrayS("L", "L", []string{}, "Add a directory to the library search path")
	rootCmd.Flags().BoolS("O", "O", false, "Equivalent to -C opt-level=2")
	rootCmd.Flags().StringP("allow", "A", "", "Set lint allowed")
	rootCmd.Flags().String("cap-lints", "", "Set the most restrictive lint level")
	rootCmd.Flags().String("cfg", "", "Configure the compilation environment")
	rootCmd.Flags().StringP("codegen", "C", "", "Set a codegen option")
	rootCmd.Flags().String("crate-name", "", "Specify the name of the crate being built")
	rootCmd.Flags().String("crate-type", "", "separated list of types of crates for the compiler to emit")
	rootCmd.Flags().StringP("deny", "D", "", "Set lint denied")
	rootCmd.Flags().String("edition", "", "Specify which edition of the compiler to use when compiling code")
	rootCmd.Flags().String("emit", "", "Comma separated list of types of output for the compiler to emit")
	rootCmd.Flags().String("explain", "", "Provide a detailed explanation of an error message")
	rootCmd.Flags().StringP("forbid", "F", "", "Set lint forbidden")
	rootCmd.Flags().BoolS("g", "g", false, "Equivalent to -C debuginfo=2")
	rootCmd.Flags().BoolP("help", "h", false, "Display this message")
	rootCmd.Flags().StringArrayS("l", "l", []string{}, "Link the generated crate(s) to the specified native library NAME")
	rootCmd.Flags().StringS("o", "o", "", "Write output to <filename>")
	rootCmd.Flags().String("out-dir", "", "Write output to compiler-chosen filename in <dir>")
	rootCmd.Flags().String("print", "", "Compiler information to print on stdout")
	rootCmd.Flags().String("target", "", "Target triple for which the code is compiled")
	rootCmd.Flags().Bool("test", false, "Build a test harness")
	rootCmd.Flags().BoolP("verbose", "v", false, "Use verbose output")
	rootCmd.Flags().BoolP("version", "V", false, "Print version info and exit")
	rootCmd.Flags().StringP("warn", "W", "", "Set lint warnings")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"L": carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues("dependency", "crate", "native", "framework", "all").Invoke(c).Suffix("=").ToA()
			case 1:
				return carapace.ActionFiles()
			default:
				return carapace.ActionValues()
			}
		}),
		"crate-type": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("bin", "lib", "rlib", "dylib", "cdylib", "staticlib", "proc-macro").Invoke(c).Filter(c.Parts).ToA()
		}),
		"edition": carapace.ActionValues("2015", "2018", "2021"),
		"emit": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("asm", "llvm-bc", "llvm-ir", "obj", "metadata", "link", "dep-info", "mir").Invoke(c).Filter(c.Parts).ToA()
		}),
		"l": carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues("static", "framework", "dylib").Invoke(c).Suffix("=").ToA()
			case 1:
				return carapace.ActionFiles()
			default:
				return carapace.ActionValues()
			}
		}),
		"o":       carapace.ActionFiles(),
		"out-dir": carapace.ActionFiles(),
		"print":   carapace.ActionValues("crate-name", "file-names", "sysroot", "target-libdir", "cfg", "target-list", "target-cpus", "target-features", "relocation-models", "code-models", "tls-models", "target-spec-json", "native-static-libs"),
		// TODO target triple (see rustup action)
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(".rs"),
	)
}
