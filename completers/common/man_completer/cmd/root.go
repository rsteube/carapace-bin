package cmd

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "man",
	Short: "an interface to the system reference manuals",
	Long:  "https://man7.org/linux/man-pages/man1/man.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("S", "S", false, "use colon separated section list")
	rootCmd.Flags().BoolP("all", "a", false, "find all matching manual pages")
	rootCmd.Flags().BoolP("apropos", "k", false, "equivalent to apropos")
	rootCmd.Flags().BoolP("ascii", "7", false, "display ASCII translation of certain latin1 chars")
	rootCmd.Flags().BoolP("catman", "c", false, "used by catman to reformat out of date cat pages")
	rootCmd.Flags().StringP("config-file", "C", "", "use this user configuration file")
	rootCmd.Flags().BoolP("debug", "d", false, "emit debugging messages")
	rootCmd.Flags().BoolP("default", "D", false, "reset all options to their default values")
	rootCmd.Flags().BoolP("ditroff", "Z", false, "use groff and force it to produce ditroff")
	rootCmd.Flags().StringP("encoding", "E", "", "use selected output encoding")
	rootCmd.Flags().StringP("extension", "e", "", "limit search to extension type EXTENSION")
	rootCmd.Flags().BoolP("global-apropos", "K", false, "search for text in all pages")
	rootCmd.Flags().StringP("gxditview", "X", "", "use groff and display through gxditview")
	rootCmd.Flags().BoolP("help", "?", false, "give this help list")
	rootCmd.Flags().StringP("html", "H", "", "use  or BROWSER to display HTML output")
	rootCmd.Flags().BoolP("ignore-case", "i", false, "look for pages case-insensitively (default)")
	rootCmd.Flags().BoolP("local-file", "l", false, "interpret PAGE argument(s) as local filename(s)")
	rootCmd.Flags().StringP("locale", "L", "", "define the locale for this particular man search")
	rootCmd.Flags().Bool("location", false, "print physical location of man page(s)")
	rootCmd.Flags().Bool("location-cat", false, "print physical location of cat file(s)")
	rootCmd.Flags().StringP("manpath", "M", "", "set search path for manual pages to PATH")
	rootCmd.Flags().BoolP("match-case", "I", false, "look for pages case-sensitively")
	rootCmd.Flags().Bool("names-only", false, "make --regex and --wildcard match page names only,")
	rootCmd.Flags().Bool("nh", false, "turn off hyphenation")
	rootCmd.Flags().Bool("nj", false, "turn off justification")
	rootCmd.Flags().Bool("no-hyphenation", false, "turn off hyphenation")
	rootCmd.Flags().Bool("no-justification", false, "turn off justification")
	rootCmd.Flags().Bool("no-subpages", false, "don't try subpages, e.g. 'man foo bar' => 'man")
	rootCmd.Flags().StringP("pager", "P", "", "use program PAGER to display output")
	rootCmd.Flags().Bool("path", false, "print physical location of man page(s)")
	rootCmd.Flags().StringP("preprocessor", "p", "", "STRING indicates which preprocessors to run")
	rootCmd.Flags().StringP("prompt", "r", "", "provide the `less' pager with a prompt")
	rootCmd.Flags().StringP("recode", "R", "", "output source page encoded in ENCODING")
	rootCmd.Flags().Bool("regex", false, "show all pages matching regex")
	rootCmd.Flags().StringP("sections", "s", "", "use colon separated section list")
	rootCmd.Flags().StringP("systems", "m", "", "use manual pages from other systems")
	rootCmd.Flags().BoolP("troff", "t", false, "use groff to format pages")
	rootCmd.Flags().StringP("troff-device", "T", "", "use groff with selected device")
	rootCmd.Flags().BoolP("update", "u", false, "force a cache consistency check")
	rootCmd.Flags().Bool("usage", false, "give a short usage message")
	rootCmd.Flags().BoolP("version", "V", false, "print program version")
	rootCmd.Flags().String("warnings", "", "enable warnings from groff")
	rootCmd.Flags().BoolP("whatis", "f", false, "equivalent to whatis")
	rootCmd.Flags().BoolP("where", "w", false, "print physical location of man page(s)")
	rootCmd.Flags().BoolP("where-cat", "W", false, "print physical location of cat file(s)")
	rootCmd.Flags().Bool("wildcard", false, "show all pages matching wildcard")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config-file": carapace.ActionFiles(),
		"manpath":     carapace.ActionDirectories(),
		"pager":       carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		ActionManPages(),
	)
}

func ActionManPages() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("man", "-k", c.CallbackValue)(func(output []byte) carapace.Action {
			r := regexp.MustCompile(`^(?P<name>.*) \(\d+\) +- (?P<description>.*)$`)

			vals := make([]string, 0)
			for _, line := range strings.Split(string(output), "\n") {
				if r.MatchString(line) {
					matches := r.FindStringSubmatch(line)
					vals = append(vals, matches[1], matches[2])
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
