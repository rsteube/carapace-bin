package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "julia",
	Short: "high-level, high-performance dynamic programming language for technical computing",
	Long:  "https://julialang.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("banner", "", "Enable or disable startup banner")
	rootCmd.Flags().String("bug-report", "", "Launch a bug report session.")
	rootCmd.Flags().String("check-bounds", "", "Emit bounds checks always, never, or respect @inbounds declarations")
	rootCmd.Flags().String("code-coverage", "", "Count executions of source lines")
	rootCmd.Flags().String("color", "", "Enable or disable color text")
	rootCmd.Flags().String("compiled-modules", "", "Enable or disable incremental precompilation of modules")
	rootCmd.Flags().String("depwarn", "", "Enable or disable syntax and method deprecation warnings (\"error\" turns warnings into errors)")
	rootCmd.Flags().StringP("eval", "e", "", "Evaluate <expr>")
	rootCmd.Flags().String("handle-signals", "", "Enable or disable Julia's default signal handlers")
	rootCmd.Flags().BoolP("help", "h", false, "Print this message (--help-hidden for more)")
	rootCmd.Flags().Bool("help-hidden", false, "Uncommon options not shown by `-h`")
	rootCmd.Flags().String("history-file", "", "Load or save history")
	rootCmd.Flags().StringP("home", "H", "", "Set location of `julia` executable")
	rootCmd.Flags().BoolS("i", "i", false, "Interactive mode; REPL runs and isinteractive() is true")
	rootCmd.Flags().String("inline", "", "Control whether inlining is permitted, including overriding @inline declarations")
	rootCmd.Flags().StringP("load", "L", "", "Load <file> immediately on all processors")
	rootCmd.Flags().String("machine-file", "", "Run processes on hosts listed in <file>")
	rootCmd.Flags().String("math-mode", "", "Disallow or enable unsafe floating point optimizations")
	rootCmd.Flags().StringP("optimize", "O", "", "Set the optimization level")
	rootCmd.Flags().StringP("print", "E", "", "Evaluate <expr> and display the result")
	rootCmd.Flags().StringP("procs", "p", "", "Integer value N launches N additional local worker processes")
	rootCmd.Flags().String("project", "", "Set <dir> as the home project/environment")
	rootCmd.Flags().BoolP("quiet", "q", false, "Quiet startup: no banner, suppress REPL warnings")
	rootCmd.Flags().String("startup-file", "", "Load `~/.julia/config/startup.jl`")
	rootCmd.Flags().StringP("sysimage", "J", "", "Start up with the given system image file")
	rootCmd.Flags().String("sysimage-native-code", "", "Use native code from system image if available")
	rootCmd.Flags().StringP("threads", "t", "", "Enable N threads")
	rootCmd.Flags().String("track-allocation", "", "Count bytes allocated by each source line")
	rootCmd.Flags().BoolP("version", "v", false, "Display version information")
	rootCmd.Flags().String("warn-overwrite", "", "Enable or disable method overwrite warnings")
	rootCmd.Flags().String("warn-scope", "", "Enable or disable warning for ambiguous top-level scope")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"banner": carapace.ActionValues("yes", "no", "auto"),
		"bug-report": carapace.ActionValuesDescribed(
			"rr", "Run julia inside rr record and upload the recorded trace.",
			"rr-local", "Run julia inside rr record but do not upload the recorded trace.",
			"help", "Print help message and exit.",
		),
		"check-bounds":         carapace.ActionValues("yes", "no", "auto").StyleF(style.ForKeyword),
		"code-coverage":        carapace.ActionValues("none", "user", "all").StyleF(style.ForKeyword),
		"color":                carapace.ActionValues("yes", "no", "auto").StyleF(style.ForKeyword),
		"compiled-modules":     carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
		"depwarn":              carapace.ActionValues("yes", "no", "error").StyleF(style.ForKeyword),
		"handle-signals":       carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
		"history-file":         carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
		"home":                 carapace.ActionDirectories(),
		"inline":               carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
		"load":                 carapace.ActionFiles(),
		"machine-file":         carapace.ActionFiles(),
		"math-mode":            carapace.ActionValues("ieee", "fast"),
		"optimize":             carapace.ActionValues("0", "1", "2", "3"),
		"project":              carapace.ActionDirectories(),
		"startup-file":         carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
		"sysimage":             carapace.ActionFiles(),
		"sysimage-native-code": carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
		"track-allocation":     carapace.ActionValues("none", "user", "all").StyleF(style.ForKeyword),
		"warn-overwrite":       carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
		"warn-scope":           carapace.ActionValues("yes", "no").StyleF(style.ForKeyword),
	})

	rootCmd.Flag("optimize").NoOptDefVal = " "
	rootCmd.Flag("code-coverage").NoOptDefVal = " "
	rootCmd.Flag("track-allocation").NoOptDefVal = " "
	rootCmd.Flag("bug-report").NoOptDefVal = " "

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
