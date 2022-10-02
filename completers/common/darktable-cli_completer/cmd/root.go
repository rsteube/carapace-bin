package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "darktable-cli",
	Short: "a command line darktable variant",
	Long:  "https://darktable-org.github.io/dtdocs/en-us/special-topics/program-invocation/darktable-cli/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("apply-custom-presets", "", "custom preset")
	rootCmd.Flags().String("bpp", "", "bit depth")
	rootCmd.Flags().String("height", "", "max height")
	rootCmd.Flags().Bool("help,-h", false, "")
	rootCmd.Flags().String("hq", "", "high quality resampling")
	rootCmd.Flags().String("style", "", "style name")
	rootCmd.Flags().Bool("style-overwrite", false, "")
	rootCmd.Flags().String("upscale", "", "upscaling")
	rootCmd.Flags().Bool("verbose", false, "")
	rootCmd.Flags().Bool("version", false, "")
	rootCmd.Flags().String("width", "", "max width")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"apply-custom-presets": carapace.ActionValues("0", "1", "false", "true"),
		"hq":                   carapace.ActionValues("0", "1", "false", "true"),
		"upscale":              carapace.ActionValues("0", "1", "false", "true"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
