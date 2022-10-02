package cmd

import (
	"os"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kotlinc",
	Short: "Kotlin command-line compiler",
	Long:  "https://kotlinlang.org/docs/command-line.html#manual-install",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "-J") {
			name := strings.TrimPrefix(arg, "-")
			rootCmd.Flags().Bool(name, false, "")
		}
	}

	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("J", "J", false, "Pass an option directly to JVM")
	rootCmd.Flags().StringS("P", "P", "", "Pass an option to a plugin")
	rootCmd.Flags().BoolS("Werror", "Werror", false, "Report an error if there are any warnings")
	rootCmd.Flags().BoolS("X", "X", false, "Print a synopsis of advanced options")
	rootCmd.Flags().StringS("api-version", "api-version", "", "Allow using declarations only from the specified version of bundled libraries")
	rootCmd.Flags().StringS("classpath", "classpath", "", "List of directories and JAR/ZIP archives to search for user class files")
	rootCmd.Flags().StringS("d", "d", "", "Destination for generated class files")
	rootCmd.Flags().BoolP("expression", "e", false, "Evaluate the given string as a Kotlin script")
	rootCmd.Flags().BoolP("help", "h", false, "Print a synopsis of standard options")
	rootCmd.Flags().BoolS("include-runtime", "include-runtime", false, "Include Kotlin runtime into the resulting JAR")
	rootCmd.Flags().BoolS("java-parameters", "java-parameters", false, "Generate metadata for Java 1.8 reflection on method parameters")
	rootCmd.Flags().StringS("jdk-home", "jdk-home", "", "Include a custom JDK from the specified location into the classpath")
	rootCmd.Flags().StringS("jvm-target", "jvm-target", "", "Target version of the generated JVM bytecode")
	rootCmd.Flags().StringS("kotlin-home", "kotlin-home", "", "Path to the home directory of Kotlin compiler used for discovery of runtime libraries")
	rootCmd.Flags().StringS("language-version", "language-version", "", "Provide source compatibility with the specified version of Kotlin")
	rootCmd.Flags().StringS("module-name", "module-name", "", "Name of the generated .kotlin_module file")
	rootCmd.Flags().BoolS("no-jdk", "no-jdk", false, "Don't automatically include the Java runtime into the classpath")
	rootCmd.Flags().BoolS("no-reflect", "no-reflect", false, "Don't automatically include Kotlin reflection into the classpath")
	rootCmd.Flags().BoolS("no-stdlib", "no-stdlib", false, "Don't automatically include the Kotlin/JVM stdlib and Kotlin reflection into the classpath")
	rootCmd.Flags().BoolS("nowarn", "nowarn", false, "Generate no warnings")
	rootCmd.Flags().StringS("opt-in", "opt-in", "", "Enable usages of API that requires opt-in with an opt-in requirement marker with the given fully qualified name")
	rootCmd.Flags().BoolS("progressive", "progressive", false, "Enable progressive compiler mode.")
	rootCmd.Flags().BoolS("script", "script", false, "Evaluate the given Kotlin script (*.kts) file")
	rootCmd.Flags().StringS("script-templates", "script-templates", "", "Script definition template classes")
	rootCmd.Flags().BoolS("verbose", "verbose", false, "Enable verbose logging output")
	rootCmd.Flags().BoolS("version", "version", false, "Display compiler version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"classpath":   carapace.ActionDirectories(),
		"d":           carapace.ActionDirectories(),
		"jdk-home":    carapace.ActionDirectories(),
		"jvm-target":  carapace.ActionValues("1.6", "1.8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18"),
		"kotlin-home": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".kt"),
	)
}
