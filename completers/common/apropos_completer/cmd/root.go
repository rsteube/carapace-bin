package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "apropos",
	Short: "search the manual page names and descriptions",
	Long:  "https://linux.die.net/man/1/apropos",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("and", "a", false, "require all keywords to match")
	rootCmd.Flags().StringP("config-file", "C", "", "use this user configuration file")
	rootCmd.Flags().BoolP("debug", "d", false, "emit debugging messages")
	rootCmd.Flags().BoolP("exact", "e", false, "search each keyword for exact match")
	rootCmd.Flags().BoolP("help", "?", false, "give this help list")
	rootCmd.Flags().StringP("locale", "L", "", "define the locale for this search")
	rootCmd.Flags().BoolP("long", "l", false, "do not trim output to terminal width")
	rootCmd.Flags().StringP("manpath", "M", "", "set search path for manual pages to PATH")
	rootCmd.Flags().BoolP("regex", "r", false, "interpret each keyword as a regex")
	rootCmd.Flags().String("section", "", "search only these sections (colon-separated)")
	rootCmd.Flags().StringP("sections", "s", "", "search only these sections (colon-separated)")
	rootCmd.Flags().StringP("systems", "m", "", "use manual pages from other systems")
	rootCmd.Flags().Bool("usage", false, "give a short usage message")
	rootCmd.Flags().BoolP("verbose", "v", false, "print verbose warning messages")
	rootCmd.Flags().BoolP("version", "V", false, "print program version")
	rootCmd.Flags().BoolP("wildcard", "w", false, "the keyword(s) contain wildcards")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config-file": carapace.ActionFiles(),
		"manpath":     carapace.ActionDirectories(),
	})
}
