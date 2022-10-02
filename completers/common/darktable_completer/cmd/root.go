package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "darktable",
	Short: "a digital photography workflow application",
	Long:  "https://darktable-org.github.io/dtdocs/en-us/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("cachedir", "", "user cache directory")
	rootCmd.Flags().String("conf", "", "override config values")
	rootCmd.Flags().String("configdir", "", "user config directory")
	rootCmd.Flags().StringS("d", "d", "", "enable debug output")
	rootCmd.Flags().String("datadir", "", "data directory")
	rootCmd.Flags().Bool("disable-opencl", false, "prevent OpenCL initialization")
	rootCmd.Flags().BoolP("help", "h", false, "")
	rootCmd.Flags().String("library", "", "library file")
	rootCmd.Flags().String("localedir", "", "locale directory")
	rootCmd.Flags().String("luacmd", "", "lua command")
	rootCmd.Flags().String("moduledir", "", "module directory")
	rootCmd.Flags().String("noiseprofiles", "", "noiseprofiles json file")
	rootCmd.Flags().StringS("t", "t", "", "num openmp threads")
	rootCmd.Flags().String("tmpdir", "", "tmp directory")
	rootCmd.Flags().Bool("version", false, "print version number")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cachedir":      carapace.ActionDirectories(),
		"configdir":     carapace.ActionDirectories(),
		"d":             carapace.ActionValues("all", "cache", "camctl", "camsupport", "control", "dev", "fswatch", "input", "lighttable", "lua", "masks", "memory", "nan", "opencl", "perf", "pwstorage", "print", "sql", "ioporder", "imageio"),
		"datadir":       carapace.ActionDirectories(),
		"library":       carapace.ActionFiles(".db"),
		"localedir":     carapace.ActionDirectories(),
		"moduledir":     carapace.ActionDirectories(),
		"noiseprofiles": carapace.ActionFiles(".json"),
		"tmpdir":        carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
