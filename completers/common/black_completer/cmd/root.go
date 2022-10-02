package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "black",
	Short: "The uncompromising code formatter",
	Long:  "https://github.com/psf/black",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("check", false, "Don't write the files back , just return the status.")
	rootCmd.Flags().StringP("code", "c", "", "Format the code passed in as a string.")
	rootCmd.Flags().Bool("color", false, "Show colored diff")
	rootCmd.Flags().String("config", "", "Read configuration from FILE path.")
	rootCmd.Flags().Bool("diff", false, "Don't write the files back, just output a diff for each file on stdout.")
	rootCmd.Flags().String("exclude", "", "A regular expression that matches files and directories that should be excluded.")
	rootCmd.Flags().String("extend-exclude", "", "Like --exclude, but adds additional files and directories on top of the excluded ones.")
	rootCmd.Flags().Bool("fast", false, "Skip temporary sanity checks.")
	rootCmd.Flags().String("force-exclude", "", "Like --exclude, but files and directories matching this regex will be excluded.")
	rootCmd.Flags().BoolP("help", "h", false, "Show this message and exit.")
	rootCmd.Flags().String("include", "", "A regular expression that matches files and directories that should be included.")
	rootCmd.Flags().Bool("ipynb", false, "Format all input files like Jupyter Notebooks")
	rootCmd.Flags().StringP("line-length", "l", "", "How many characters per line to allow.")
	rootCmd.Flags().Bool("no-color", false, "Do not show colored diff")
	rootCmd.Flags().Bool("pyi", false, "Format all input files like typing stubs regardless of file extension ")
	rootCmd.Flags().BoolP("quiet", "q", false, "Don't emit non-error messages to stderr.")
	rootCmd.Flags().String("required-version", "", "Require a specific version of Black to be running.")
	rootCmd.Flags().Bool("safe", false, "Do not skip temporary sanity checks.")
	rootCmd.Flags().BoolP("skip-magic-trailing-comma", "C", false, "Don't use trailing commas as a reason to split lines.")
	rootCmd.Flags().BoolP("skip-string-normalization", "S", false, "Don't normalize string quotes or prefixes.")
	rootCmd.Flags().String("stdin-filename", "", "The name of the file when passing it through stdin.")
	rootCmd.Flags().StringP("target-version", "t", "", "Python versions that should be supported by Black's output.")
	rootCmd.Flags().BoolP("verbose", "v", false, "Also emit messages to stderr about files that were not changed.")
	rootCmd.Flags().Bool("version", false, "Show the version and exit.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config":         carapace.ActionValues(),
		"target-version": carapace.ActionValues("py27", "py33", "py34", "py35", "py36", "py37", "py38", "py39", "py310"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
