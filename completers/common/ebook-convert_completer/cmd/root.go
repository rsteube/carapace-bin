package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ebook-convert",
	Short: "Convert an e-book from one format to another",
	Long:  "https://manual.calibre-ebook.com/de/generated/de/ebook-convert.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.Flags().Bool("list-recipes", false, "List builtin recipe names. You can create an e-book from a")
	rootCmd.Flags().Bool("version", false, "show program's version number and exit")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.Batch(
			carapace.ActionFiles(),
			carapace.ActionCallback(func(c carapace.Context) carapace.Action {
				if index := strings.LastIndex(c.CallbackValue, "."); index != -1 {
					return carapace.ActionValuesDescribed(
						"mobi", "Mobipocket",
						"epub", "Electronic Publication",
						"azw3", "Amazon KF8",
						"fb2", "Fiction Book",
						"htmlz", "Zipped HTML eBook",
						"lit", "Microsoft LIT",
						"lrf", "Sony LRF",
						"pdb", "PalmDoc Book",
						"pdf", "Portable Document",
						"pmlz", "Zipped Palm Markup Language",
						"rb", "Rocket eBook",
						"rtf", "Rich Text Format",
						"snb", "S Note File",
						"tcr", "Psion Series 3 eBook",
						"txt", "Plain Text File",
						"txtz", "Zipped Plain Text File",
						"zip", "Zipped File",
					).Invoke(c).Prefix(c.CallbackValue[:index+1]).ToA().StyleF(style.ForPathExt)
				}
				return carapace.ActionValues()
			}),
		).ToA(),
	)
}
