package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/lslocks_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lslocks",
	Short: "List local system locks",
	Long:  "https://man7.org/linux/man-pages/man8/lslocks.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("bytes", "b", false, "print SIZE in bytes rather than in human readable format")
	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().BoolP("json", "J", false, "use JSON output format")
	rootCmd.Flags().BoolP("noheadings", "n", false, "don't print headings")
	rootCmd.Flags().BoolP("noinaccessible", "i", false, "ignore locks without read permissions")
	rootCmd.Flags().BoolP("notruncate", "u", false, "don't truncate text in columns")
	rootCmd.Flags().StringP("output", "o", "", "define which output columns to use")
	rootCmd.Flags().Bool("output-all", false, "output all columns")
	rootCmd.Flags().StringP("pid", "p", "", "display only locks held by this process")
	rootCmd.Flags().BoolP("raw", "r", false, "use the raw output format")
	rootCmd.Flags().BoolP("version", "V", false, "display version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionColumns().Invoke(c).Filter(c.Args).ToA()
		}),
		"pid": ps.ActionProcessIds(),
	})
}
