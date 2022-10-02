package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "openscad",
	Short: "script file based graphical CAD environment",
	Long:  "https://openscad.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("D", "D", "", "pre-define variables")
	rootCmd.Flags().StringP("P", "P", "", "customizer parameter set")
	rootCmd.Flags().Bool("autocenter", false, "adjust camera to look at object's center")
	rootCmd.Flags().String("camera", "", "camera parameters when exporting png: ")
	rootCmd.Flags().String("check-parameter-ranges", "", "configure the parameter range check for builtin modules")
	rootCmd.Flags().String("check-parameters", "", "configure the parameter check for user modules and functions")
	rootCmd.Flags().String("colorscheme", "", "colorscheme")
	rootCmd.Flags().String("csglimit", "", "stop rendering at n CSG elements")
	rootCmd.Flags().StringP("d", "d", "", "generate a dependency file for make")
	rootCmd.Flags().String("debug", "", "special debug info")
	rootCmd.Flags().Bool("hardwarnings", false, "Stop on the first warning")
	rootCmd.Flags().BoolP("help", "h", false, "print this help message and exit")
	rootCmd.Flags().String("imgsize", "", "width,height of exported png")
	rootCmd.Flags().Bool("info", false, "print information about the build process")
	rootCmd.Flags().StringP("m", "m", "", "runs make_cmd file if file is missing")
	rootCmd.Flags().StringP("o", "o", "", "output specified file instead of running the GUI")
	rootCmd.Flags().StringP("p", "p", "", "customizer parameter file")
	rootCmd.Flags().String("preview", "", "preview ")
	rootCmd.Flags().String("projection", "", "=(o)rtho or (p)erspective when exporting png")
	rootCmd.Flags().BoolP("quiet", "q", false, "quiet mode")
	rootCmd.Flags().String("render", "", "for full geometry evaluation when exporting png")
	rootCmd.Flags().BoolP("version", "v", false, "print the version")
	rootCmd.Flags().String("view", "", "view option")
	rootCmd.Flags().Bool("viewall", false, "adjust camera to fit object")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"colorscheme": carapace.ActionValues("Cornfield", "Metallic", "Sunset", "Starnight", "BeforeDawn", "Nature", "DeepOcean", "Solarized", "Tomorrow", "Tomorrow Night", "Monotone"),
		"d":           carapace.ActionFiles(),
		"o":           carapace.ActionFiles(),
		"p":           carapace.ActionFiles(),
		"view":        carapace.ActionValues("axes", "crosshairs", "edges", "scales", "wireframe"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(carapace.ActionFiles(".scad"))
}
