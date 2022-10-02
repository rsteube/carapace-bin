package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gimp",
	Short: "an image manipulation and paint program",
	Long:  "https://www.gimp.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("as-new", "a", false, "Open images as new")
	rootCmd.Flags().StringP("batch", "b", "", "Batch command to run (can be used multiple times)")
	rootCmd.Flags().String("batch-interpreter", "", "The procedure to process batch commands with")
	rootCmd.Flags().String("class", "", "Program class as used by the window manager")
	rootCmd.Flags().BoolP("console-messages", "c", false, "Send messages to console instead of using a dialog")
	rootCmd.Flags().Bool("debug-handlers", false, "Enable non-fatal debugging signal handlers")
	rootCmd.Flags().String("display", "", "X display to use")
	rootCmd.Flags().Bool("dump-gimprc", false, "Output a gimprc file with default settings")
	rootCmd.Flags().Bool("g-fatal-warnings", false, "Make all warnings fatal")
	rootCmd.Flags().String("gegl-cache-size", "", "How much memory to (approximately) use for caching imagery")
	rootCmd.Flags().Bool("gegl-disable-opencl", false, "Disable OpenCL")
	rootCmd.Flags().String("gegl-quality", "", "The quality of rendering, a value between 0.0 (fast) and 1.0 (reference)")
	rootCmd.Flags().String("gegl-swap", "", "Where GEGL stores its swap")
	rootCmd.Flags().String("gegl-swap-compression", "", "Compression algorithm used for data stored in the swap")
	rootCmd.Flags().String("gegl-threads", "", "The number of concurrent processing threads to use")
	rootCmd.Flags().String("gegl-tile-size", "", "Default size of tiles in GeglBuffers")
	rootCmd.Flags().StringP("gimprc", "g", "", "Use an alternate user gimprc file")
	rootCmd.Flags().Bool("gtk-g-fatal-warnings", false, "Make all warnings fatal")
	rootCmd.Flags().String("gtk-module", "", "Load additional GTK+ modules")
	rootCmd.Flags().BoolP("help", "h", false, "Show help options")
	rootCmd.Flags().Bool("help-all", false, "Show all help options")
	rootCmd.Flags().Bool("help-gegl", false, "Show GEGL Options")
	rootCmd.Flags().Bool("help-gtk", false, "Show GTK+ Options")
	rootCmd.Flags().Bool("license", false, "Show license information and exit")
	rootCmd.Flags().String("name", "", "Program name as used by the window manager")
	rootCmd.Flags().BoolP("new-instance", "n", false, "Start a new GIMP instance")
	rootCmd.Flags().Bool("no-cpu-accel", false, "Do not use special CPU acceleration functions")
	rootCmd.Flags().BoolP("no-data", "d", false, "Do not load brushes, gradients, patterns, ...")
	rootCmd.Flags().BoolP("no-fonts", "f", false, "Do not load any fonts")
	rootCmd.Flags().BoolP("no-interface", "i", false, "Run without a user interface")
	rootCmd.Flags().Bool("no-shm", false, "Do not use shared memory between GIMP and plug-ins")
	rootCmd.Flags().BoolP("no-splash", "s", false, "Do not show a splash screen")
	rootCmd.Flags().String("pdb-compat-mode", "", "PDB compatibility mode (off|on|warn)")
	rootCmd.Flags().String("screen", "", "X screen to use")
	rootCmd.Flags().String("session", "", "Use an alternate sessionrc file")
	rootCmd.Flags().Bool("show-playground", false, "Show a preferences page with experimental features")
	rootCmd.Flags().String("stack-trace-mode", "", "Debug in case of a crash (never|query|always)")
	rootCmd.Flags().Bool("sync", false, "Make X calls synchronous")
	rootCmd.Flags().String("system-gimprc", "", "Use an alternate system gimprc file")
	rootCmd.Flags().Bool("verbose", false, "Be more verbose")
	rootCmd.Flags().BoolP("version", "v", false, "Show version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"display":          os.ActionDisplays(),
		"gimprc":           carapace.ActionFiles(),
		"pdb-compat-mode":  carapace.ActionValues("off", "on", "warn").StyleF(style.ForKeyword),
		"stack-trace-mode": carapace.ActionValues("never", "query", "always").StyleF(style.ForKeyword),
		"system-gimprc":    carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
