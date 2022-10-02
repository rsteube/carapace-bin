package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "apt-cache",
	Short: "query the APT cache",
	Long:  "https://linux.die.net/man/8/apt-cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("all-names", false, "Make pkgnames print all names")
	rootCmd.Flags().BoolP("all-versions", "a", false, "Print full records for all available versions")
	rootCmd.Flags().StringP("config-file", "c", "", "Configuration File")
	rootCmd.Flags().BoolP("full", "f", false, "Print full package records when searching")
	rootCmd.Flags().BoolP("generate", "g", false, "Perform automatic package cache regeneration")
	rootCmd.Flags().BoolP("help", "h", false, "Show a short usage summary")
	rootCmd.Flags().Bool("implicit", false, "Per default depends and rdepends print only dependencies explicitly expressed in the metadata")
	rootCmd.Flags().BoolP("important", "i", false, "Print only important dependencies")
	rootCmd.Flags().Bool("installed", false, "Limit the output of depends and rdepends to packages which are currently installed")
	rootCmd.Flags().BoolP("names-only,", "n", false, "Only search on the package and provided package names")
	rootCmd.Flags().Bool("no-breaks", false, "Omit breaks")
	rootCmd.Flags().Bool("no-conflicts", false, "Omit conflicts")
	rootCmd.Flags().Bool("no-depends", false, "Omit depends")
	rootCmd.Flags().Bool("no-enhances", false, "Omit enhances")
	rootCmd.Flags().Bool("no-pre-depends", false, "Omit pre-depends")
	rootCmd.Flags().Bool("no-recommends", false, "Omit recommends")
	rootCmd.Flags().Bool("no-replaces", false, "Omit replaces")
	rootCmd.Flags().Bool("no-suggests", false, "Omit suggests")
	rootCmd.Flags().StringArrayP("option", "o", []string{}, "Set a Configuration Option")
	rootCmd.Flags().StringP("pkg-cache", "p", "", "Select the file to store the package cache")
	rootCmd.Flags().BoolP("quiet", "q", false, "Quiet")
	rootCmd.Flags().Bool("recurse", false, "Make depends and rdepends recursive so that all packages mentioned are printed once")
	rootCmd.Flags().StringP("src-cache", "s", "", "Select the file to store the source cache")
	rootCmd.Flags().BoolP("version", "v", false, "Show the program version")
	rootCmd.Flags().String("with-source", "", "Adds the given file as a source for metadata")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config-file": carapace.ActionFiles(),
		"pkg-cache":   carapace.ActionFiles(),
		"src-cache":   carapace.ActionFiles(),
		"with-source": carapace.ActionFiles(),
	})
}
