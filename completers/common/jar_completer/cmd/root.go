package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jar",
	Short: "create an archive for classes and resources",
	Long:  "https://docs.oracle.com/en/java/javase/16/docs/specs/man/jar.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("?", "?", false, "Show help")
	rootCmd.Flags().StringS("C", "C", "", "Change to the specified directory and include the following file")
	rootCmd.Flags().BoolP("create", "c", false, "Create the archive")
	rootCmd.Flags().BoolP("describe-module", "d", false, "Print the module descriptor, or automatic module name")
	rootCmd.Flags().BoolP("extract", "x", false, "Extract named (or all) files from the archive")
	rootCmd.Flags().StringP("file", "f", "", "The archive file name")
	rootCmd.Flags().StringP("generate-index", "i", "", "Generate index information for the specified jar archives")
	rootCmd.Flags().String("hash-modules", "", "Compute and record the hashes of modules matched by the given pattern")
	rootCmd.Flags().BoolP("help", "h", false, "Show help ")
	rootCmd.Flags().Bool("help-extra", false, "Give help on extra options")
	rootCmd.Flags().Bool("help:compat", false, "Show compat help ")
	rootCmd.Flags().BoolP("list", "t", false, "List the table of contents for the archive")
	rootCmd.Flags().StringP("main-class", "e", "", "The application entry point for stand-alone applications")
	rootCmd.Flags().StringP("manifest", "m", "", "Include the manifest information from the given manifest file")
	rootCmd.Flags().BoolP("module-path", "p", false, "Location of module dependence for generating the hash")
	rootCmd.Flags().String("module-version", "", "The module version")
	rootCmd.Flags().BoolP("no-compress", "0", false, "Store only; use no ZIP compression")
	rootCmd.Flags().BoolP("no-manifest", "M", false, "Do not create a manifest file for the entries")
	rootCmd.Flags().String("release", "", "Places all following files in a versioned directory of the jar")
	rootCmd.Flags().BoolP("update", "u", false, "Update an existing jar archive")
	rootCmd.Flags().Bool("validate", false, "Validate the contents of the jar archive")
	rootCmd.Flags().BoolP("verbose", "v", false, "Generate verbose output on standard output")
	rootCmd.Flags().Bool("version", false, "Print program version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"C":              carapace.ActionDirectories(),
		"file":           carapace.ActionFiles(),
		"generate-index": carapace.ActionFiles(),
		"manifest":       carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if strings.HasPrefix(c.CallbackValue, "@") {
				c.CallbackValue = strings.TrimPrefix(c.CallbackValue, "@")
				return carapace.ActionFiles().Invoke(c).Prefix("@").ToA()
			}
			return carapace.ActionFiles()
		}),
	)
}
