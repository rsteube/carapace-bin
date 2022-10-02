package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/golang"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-carpet",
	Short: "show test coverage for Go source files",
	Long:  "https://github.com/msoap/go-carpet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("256colors", false, "use more colors on 256-color terminal")
	rootCmd.Flags().String("args", "", "pass additional arguments for go test")
	rootCmd.Flags().String("file", "", "comma-separated list of files to test")
	rootCmd.Flags().String("func", "", "comma-separated functions list")
	rootCmd.Flags().Bool("include-vendor", false, "include vendor directories for show coverage")
	rootCmd.Flags().Bool("summary", false, "only show summary for each file")
	rootCmd.Flags().Bool("version", false, "get version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionFiles()
		}),
		"func": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			files := []string{}
			if f := rootCmd.Flag("file"); f.Changed {
				files = strings.Split(f.Value.String(), ",")
			}
			return golang.ActionFuncs(files...).Invoke(c).Filter(c.Parts).ToA() // TODO only for given files
		}),
	})
}
