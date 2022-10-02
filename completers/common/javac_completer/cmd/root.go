package cmd

import (
	"os"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/javac_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "javac",
	Short: "Reads Java class and interface definitions and compiles them into bytecode and class files",
	Long:  "https://en.wikipedia.org/wiki/Javac",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	for _, arg := range os.Args {
		// yuck - these are flags that require an argument but don't allow space between name and value
		// which essentially makes them new flags. To prevent unkown flag errors these are added as fake
		// flags just as they are.
		for _, flag := range []string{"-A", "-J"} { // TODO support `=` in `-A`
			if strings.HasPrefix(arg, flag) && len(arg) > 2 {
				fakeflag := strings.SplitN(arg[1:], "=", 2)[0]
				if rootCmd.Flag(fakeflag) == nil {
					rootCmd.Flags().String(fakeflag, "", "") // fake flag to prevent errors
					rootCmd.Flag(fakeflag).NoOptDefVal = " "
				}
			}
		}
	}

	carapace.Override(carapace.Opts{
		OptArgDelimiter: ":",
	})
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("A", "A", "", "Options to pass to annotation processors")
	rootCmd.Flags().BoolS("J", "J", false, "Pass flag directly to the runtime system")
	rootCmd.Flags().BoolS("Werror", "Werror", false, "Terminate compilation if warnings occur")
	rootCmd.Flags().StringS("bootclasspath", "bootclasspath", "", "Override location of bootstrap class files")
	rootCmd.Flags().StringS("classpath", "classpath", "", "Specify where to find user class files and annotation processors")
	rootCmd.Flags().StringS("cp", "cp", "", "Specify where to find user class files and annotation processors")
	rootCmd.Flags().BoolS("d", "d", false, "Specify where to place generated class files")
	rootCmd.Flags().BoolS("deprecation", "deprecation", false, "Output source locations where deprecated APIs are used")
	rootCmd.Flags().StringS("encoding", "encoding", "", "Specify character encoding used by source files")
	rootCmd.Flags().StringS("endorseddirs", "endorseddirs", "", "Override location of endorsed standards path")
	rootCmd.Flags().StringS("extdirs", "extdirs", "", "Override location of installed extensions")
	rootCmd.Flags().StringS("g", "g", "", "Generates all debugging information")
	rootCmd.Flags().BoolS("h", "h", false, "Specify where to place generated native header files")
	rootCmd.Flags().BoolS("help", "help", false, "Print a synopsis of standard options")
	rootCmd.Flags().StringS("implicit", "implicit", "", "Specify whether or not to generate class files for implicitly referenced files")
	rootCmd.Flags().BoolS("nowarn", "nowarn", false, "Generate no warnings")
	rootCmd.Flags().BoolS("parameters", "parameters", false, "Generate metadata for reflection on method parameters")
	rootCmd.Flags().StringS("proc", "proc", "", "Control whether annotation processing and/or compilation is done.")
	rootCmd.Flags().StringS("processor", "processor", "", "Names of the annotation processors to run; bypasses default discovery process")
	rootCmd.Flags().StringS("processorpath", "processorpath", "", "Specify where to find annotation processors")
	rootCmd.Flags().StringS("profile", "profile", "", "Check that API used is available in the specified profile")
	rootCmd.Flags().BoolS("s", "s", false, "Specify where to place generated source files")
	rootCmd.Flags().StringS("source", "source", "", "Provide source compatibility with specified release")
	rootCmd.Flags().StringS("sourcepath", "sourcepath", "", "Specify where to find input source files")
	rootCmd.Flags().StringS("target", "target", "", "Generate class files for specific VM version")
	rootCmd.Flags().BoolS("verbose", "verbose", false, "Output messages about what the compiler is doing")
	rootCmd.Flags().BoolS("version", "version", false, "Version information")

	rootCmd.Flag("g").NoOptDefVal = " "
	rootCmd.Flag("implicit").NoOptDefVal = " "
	rootCmd.Flag("proc").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"bootclasspath": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			return carapace.ActionFiles(".jar", ".zip")
		}),
		"classpath": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			return carapace.ActionFiles(".jar", ".zip")
		}),
		"cp": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			return carapace.ActionFiles(".jar", ".zip")
		}),
		"d": carapace.ActionDirectories(),
		// TODO encoding
		"endorseddirs": carapace.ActionDirectories(),
		"extdirs":      carapace.ActionDirectories(),
		"g": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValuesDescribed(
				"none", "Does not generate any debugging information",
				"source", "Source file debugging information",
				"lines", "Line number debugging information",
				"vars", "Local variable debugging information",
			).Invoke(c).Filter(c.Parts).ToA()
		}),
		"h":             carapace.ActionDirectories(),
		"implicit":      carapace.ActionValues("none", "class"),
		"proc":          carapace.ActionValues("none", "only"),
		"processorpath": carapace.ActionDirectories(),
		"profile":       carapace.ActionValues("compact1", "compact2", "compact3"),
		"s":             carapace.ActionDirectories(),
		"source":        action.ActionReleases(),
		"target":        action.ActionReleases(), // TODO limit to `>+ source version`
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if strings.HasPrefix(c.CallbackValue, "@") {
				return carapace.ActionFiles().Invoke(c).Prefix("@").ToA()
			}
			return carapace.ActionFiles(".java")
		}),
	)
}
