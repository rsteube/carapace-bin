package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "csview",
	Short: "A high performance csv viewer with cjk/emoji support",
	Long:  "https://github.com/wfxr/csview",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("delimiter", "d", "", "Specify the field delimiter")
	rootCmd.Flags().BoolP("help", "h", false, "Print help information")
	rootCmd.Flags().StringP("indent", "i", "", "Specify global indent for table")
	rootCmd.Flags().BoolP("no-headers", "H", false, "Specify that the input has no header row")
	rootCmd.Flags().StringP("padding", "p", "", "Specify padding for table cell")
	rootCmd.Flags().String("sniff", "", "Limit column widths sniffing to the specified number of rows")
	rootCmd.Flags().StringP("style", "s", "", "Specify the border style")
	rootCmd.Flags().BoolP("tsv", "t", false, "Use '\\t' as delimiter for tsv")
	rootCmd.Flags().BoolP("version", "V", false, "Print version information")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"style": carapace.ActionValues("None", "Ascii", "Sharp", "Rounded", "Reinforced", "Markdown", "Grid"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
