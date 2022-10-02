package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gdu",
	Short: "Pretty fast disk usage analyzer written in Go",
	Long:  "https://github.com/dundee/gdu",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().BoolP("const-gc", "g", false, "Enable memory garbage collection during analysis with constant level set by GOGC")
	rootCmd.Flags().Bool("enable-profiling", false, "Enable collection of profiling data and provide it on http://localhost:6060/debug/pprof/")
	rootCmd.Flags().BoolP("help", "h", false, "help for gdu")
	rootCmd.Flags().StringSliceP("ignore-dirs", "i", []string{"/proc", "/dev", "/sys", "/run"}, "Absolute paths to ignore (separated by comma)")
	rootCmd.Flags().StringSliceP("ignore-dirs-pattern", "I", []string{}, "Absolute path patterns to ignore (separated by comma)")
	rootCmd.Flags().StringP("ignore-from", "X", "", "Read absolute path patterns to ignore from file")
	rootCmd.Flags().StringP("input-file", "f", "", "Import analysis from JSON file")
	rootCmd.Flags().StringP("log-file", "l", "/dev/null", "Path to a logfile")
	rootCmd.Flags().IntP("max-cores", "m", 4, "Set max cores that GDU will use. 4 cores available")
	rootCmd.Flags().BoolP("no-color", "c", false, "Do not use colorized output")
	rootCmd.Flags().BoolP("no-cross", "x", false, "Do not cross filesystem boundaries")
	rootCmd.Flags().BoolP("no-hidden", "H", false, "Ignore hidden directories (beginning with dot)")
	rootCmd.Flags().BoolP("no-progress", "p", false, "Do not show progress in non-interactive mode")
	rootCmd.Flags().BoolP("non-interactive", "n", false, "Do not run in interactive mode")
	rootCmd.Flags().StringP("output-file", "o", "", "Export all info into file as JSON")
	rootCmd.Flags().BoolP("show-apparent-size", "a", false, "Show apparent size")
	rootCmd.Flags().BoolP("show-disks", "d", false, "Show all mounted disks")
	rootCmd.Flags().BoolP("show-relative-size", "B", false, "Show relative size")
	rootCmd.Flags().Bool("si", false, "Show sizes with decimal SI prefixes (kB, MB, GB) instead of binary prefixes (KiB, MiB, GiB)")
	rootCmd.Flags().BoolP("summarize", "s", false, "Show only a total in non-interactive mode")
	rootCmd.Flags().BoolP("version", "v", false, "Print version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"ignore-dirs": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionDirectories()
		}),
		"ignore-from": carapace.ActionFiles(),
		"input-file":  carapace.ActionFiles(),
		"log-file":    carapace.ActionFiles(),
		"output-file": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
