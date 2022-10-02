package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var fmtCmd = &cobra.Command{
	Use:   "fmt",
	Short: "Format source files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fmtCmd).Standalone()

	fmtCmd.Flags().Bool("check", false, "Check if the source files are formatted")
	fmtCmd.Flags().StringP("config", "c", "", "Load configuration file")
	fmtCmd.Flags().String("ext", "", "Set standard input (stdin) content type [default: ts]  [possible values: ts, tsx, js, jsx, md, json, jsonc]")
	fmtCmd.Flags().String("ignore", "", "Ignore formatting particular source files")
	fmtCmd.Flags().String("options-indent-width", "", "Define indentation width. Defaults to 2.")
	fmtCmd.Flags().String("options-line-width", "", "Define maximum line width. Defaults to 80.")
	fmtCmd.Flags().String("options-prose-wrap", "", "Define how prose should be wrapped. Defaults to always. [possible values: always, never, preserve]")
	fmtCmd.Flags().Bool("options-single-quote", false, "Use single quotes. Defaults to false.")
	fmtCmd.Flags().Bool("options-use-tabs", false, "Use tabs instead of spaces for indentation. Defaults to false.")
	fmtCmd.Flags().Bool("watch", false, "UNSTABLE: Watch for file changes and restart process automatically")
	rootCmd.AddCommand(fmtCmd)

	carapace.Gen(fmtCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
		"ext": carapace.ActionValues("ts", "tsx", "js", "jsx", "md", "json", "jsonc").StyleF(func(s string) string {
			return style.ForPathExt("." + s)
		}),
		"options-prose-wrap": carapace.ActionValues("always", "never", "preserve").StyleF(style.ForKeyword),
	})

	carapace.Gen(fmtCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
