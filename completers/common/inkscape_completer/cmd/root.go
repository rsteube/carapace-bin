package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "inkscape",
	Short: "an SVG (Scalable Vector Graphics) editing program",
	Long:  "https://inkscape.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("action-list", false, "List all available actions")
	rootCmd.Flags().String("actions", "", "List of actions (with optional arguments) to execute")
	rootCmd.Flags().Bool("batch-process", false, "Close GUI after executing all actions/verbs")
	rootCmd.Flags().String("class", "", "Program class as used by the window manager")
	rootCmd.Flags().String("convert-dpi-method", "", "Method used to convert pre-0.92 document dpi")
	rootCmd.Flags().Bool("dbus-listen", false, "Enter a listening loop for D-Bus messages in console mode")
	rootCmd.Flags().String("dbus-name", "", "Specify the D-Bus name; default is 'org.inkscape'")
	rootCmd.Flags().String("display", "", "X display to use")
	rootCmd.Flags().StringP("export-area", "a", "", "Area to export in SVG user units")
	rootCmd.Flags().BoolP("export-area-drawing", "D", false, "Area to export is whole drawing (ignoring page size)")
	rootCmd.Flags().BoolP("export-area-page", "C", false, "Area to export is page")
	rootCmd.Flags().Bool("export-area-snap", false, "Snap the bitmap export area outwards to the nearest integer values")
	rootCmd.Flags().StringP("export-background", "b", "", "Background color for exported bitmaps (any SVG color string)")
	rootCmd.Flags().StringP("export-background-opacity", "y", "", "Background opacity for exported bitmaps (0.0 to 1.0, or 1 to 255)")
	rootCmd.Flags().StringP("export-dpi", "d", "", "Resolution for bitmaps and rasterized filters; default is 96")
	rootCmd.Flags().StringP("export-filename", "o", "", "Output file name (file type is guessed from extension)")
	rootCmd.Flags().StringP("export-height", "h", "", "Bitmap height in pixels (overrides --export-dpi)")
	rootCmd.Flags().StringP("export-id", "i", "", "ID(s) of object(s) to export")
	rootCmd.Flags().BoolP("export-id-only", "j", false, "Hide all objects except object with ID selected by export-id")
	rootCmd.Flags().Bool("export-ignore-filters", false, "Render objects without filters instead of rasterizing (PS/EPS/PDF)")
	rootCmd.Flags().Bool("export-latex", false, "Export text separately to LaTeX file (PS/EPS/PDF)")
	rootCmd.Flags().String("export-margin", "", "Margin around export area: units of page size for SVG, mm for PS/EPS/PDF")
	rootCmd.Flags().Bool("export-overwrite", false, "Overwrite input file")
	rootCmd.Flags().String("export-pdf-version", "", "PDF version (1.4 or 1.5)")
	rootCmd.Flags().BoolP("export-plain-svg", "l", false, "Remove Inkscape-specific SVG attributes/properties")
	rootCmd.Flags().String("export-ps-level", "", "Postscript level (2 or 3); default is 3")
	rootCmd.Flags().BoolP("export-text-to-path", "T", false, "Convert text to paths (PS/EPS/PDF/SVG)")
	rootCmd.Flags().String("export-type", "", "File type(s) to export")
	rootCmd.Flags().BoolP("export-use-hints", "t", false, "Use stored filename and DPI hints when exporting object selected by --export-id")
	rootCmd.Flags().StringP("export-width", "w", "", "Bitmap width in pixels (overrides --export-dpi)")
	rootCmd.Flags().Bool("g-fatal-warnings", false, "Make all warnings fatal")
	rootCmd.Flags().Bool("gapplication-service", false, "Enter GApplication service mode (use from D-Bus service files)")
	rootCmd.Flags().String("gdk-debug", "", "GDK debugging flags to set")
	rootCmd.Flags().String("gdk-no-debug", "", "GDK debugging flags to unset")
	rootCmd.Flags().String("gtk-debug", "", "GTK+ debugging flags to set")
	rootCmd.Flags().String("gtk-module", "", "Load additional GTK+ modules")
	rootCmd.Flags().String("gtk-no-debug", "", "GTK+ debugging flags to unset")
	rootCmd.Flags().BoolP("help", "?", false, "Show help options")
	rootCmd.Flags().Bool("help-all", false, "Show all help options")
	rootCmd.Flags().Bool("help-gapplication", false, "Show GApplication options")
	rootCmd.Flags().Bool("help-gtk", false, "Show GTK+ Options")
	rootCmd.Flags().String("name", "", "Program name as used by the window manager")
	rootCmd.Flags().Bool("no-convert-text-baseline-spacing", false, "Do not fix pre-0.92 document's text baseline spacing on opening")
	rootCmd.Flags().String("pdf-page", "", "PDF page number to import")
	rootCmd.Flags().Bool("pdf-poppler", false, "Use poppler when importing via commandline")
	rootCmd.Flags().BoolP("pipe", "p", false, "Read input file from standard input (stdin)")
	rootCmd.Flags().BoolP("query-all", "S", false, "Print bounding boxes of all objects")
	rootCmd.Flags().BoolP("query-height", "H", false, "Height of drawing or object (if specified by --query-id)")
	rootCmd.Flags().StringP("query-id", "I", "", "ID(s) of object(s) to be queried")
	rootCmd.Flags().BoolP("query-width", "W", false, "Width of drawing or object (if specified by --query-id)")
	rootCmd.Flags().BoolP("query-x", "X", false, "X coordinate of drawing or object (if specified by --query-id)")
	rootCmd.Flags().BoolP("query-y", "Y", false, "Y coordinate of drawing or object (if specified by --query-id)")
	rootCmd.Flags().String("select", "", "Select objects: comma-separated list of IDs")
	rootCmd.Flags().Bool("shell", false, "Start Inkscape in interactive shell mode")
	rootCmd.Flags().Bool("system-data-directory", false, "Print system data directory")
	rootCmd.Flags().Bool("user-data-directory", false, "Print user data directory")
	rootCmd.Flags().Bool("vacuum-defs", false, "Remove unused definitions from the <defs> section(s) of document")
	rootCmd.Flags().String("verb", "", "List of verbs to execute")
	rootCmd.Flags().Bool("verb-list", false, "List all available verbs")
	rootCmd.Flags().BoolP("version", "V", false, "Print Inkscape version")
	rootCmd.Flags().BoolP("with-gui", "g", false, "With graphical user interface (required by some actions/verbs)")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"convert-dpi-method": carapace.ActionValues("none", "scale-viewbox", "scale-document"),
		"display":            os.ActionDisplays(),
		"export-filename":    carapace.ActionFiles(),
		"export-pdf-version": carapace.ActionValues("1.4", "1.5"),
		"export-type": carapace.ActionValues("svg", "png", "ps", "eps", "pdf", "emf", "wmf", "xaml").StyleF(func(s string) string {
			return style.ForPathExt("." + s)
		}),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
