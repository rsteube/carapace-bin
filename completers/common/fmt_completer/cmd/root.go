package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fmt",
	Short: "simple optimal text formatter",
	Long:  "https://linux.die.net/man/1/fmt",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("crown-margin", "c", false, "preserve indentation of first two lines")
	rootCmd.Flags().StringP("goal", "g", "", "goal width (default of 93% of width)")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().StringP("prefix", "p", "", "reformat only lines beginning with STRING,")
	rootCmd.Flags().BoolP("split-only", "s", false, "split long lines, but do not refill")
	rootCmd.Flags().BoolP("tagged-paragraph", "t", false, "indentation of first line different from second")
	rootCmd.Flags().BoolP("uniform-spacing", "u", false, "one space between words, two after sentences")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
	rootCmd.Flags().StringP("width", "w", "", "maximum line width (default of 75 columns)")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
