package cmd

import (
	"os"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/java_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "java",
	Short: "Launches a Java application",
	Long:  "https://en.wikipedia.org/wiki/Java_(programming_language)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	for _, arg := range os.Args {
		// yuck - these are flags that require an argument but don't allow space between name and value
		// which essentially makes them new flags. To prevent unkown flag errors these are added as fake
		// flags just as they are.
		for _, flag := range []string{"-D", "-Xms", "-Xmx", "-Xss"} {
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

	rootCmd.Flags().BoolS("?", "?", false, "print help message")
	rootCmd.Flags().StringArrayS("D", "D", []string{}, "set a system property")
	rootCmd.Flags().BoolS("X", "X", false, "print help on non-standard options")
	rootCmd.Flags().StringS("agentlib", "agentlib", "", "load native agent library")
	rootCmd.Flags().StringS("agentpath", "agentpath", "", "load native agent library by full pathname")
	rootCmd.Flags().StringS("classpath", "classpath", "", "class search path of directories and zip/jar files")
	rootCmd.Flags().StringS("cp", "cp", "", "class search path of directories and zip/jar files")
	rootCmd.Flags().BoolS("d32", "d32", false, "use a 32-bit data model if available")
	rootCmd.Flags().BoolS("d64", "d64", false, "use a 64-bit data model if available")
	rootCmd.Flags().StringS("da", "da", "", "disable assertions with specified granularity")
	rootCmd.Flags().StringS("disableassertions", "disableassertions", "", "disable assertions with specified granularity")
	rootCmd.Flags().BoolS("disablesystemassertions", "disablesystemassertions", false, "disable system assertions")
	rootCmd.Flags().BoolS("dsa", "dsa", false, "disable system assertions")
	rootCmd.Flags().StringS("ea", "ea", "", "enable assertions with specified granularity")
	rootCmd.Flags().StringS("enableassertions", "enableassertions", "", "enable assertions with specified granularity")
	rootCmd.Flags().BoolS("enablesystemassertions", "enablesystemassertions", false, "enable system assertions")
	rootCmd.Flags().BoolS("esa", "esa", false, "enable system assertions")
	rootCmd.Flags().BoolS("help", "help", false, "print help message")
	rootCmd.Flags().StringS("jar", "jar", "", "jar file to execute")
	rootCmd.Flags().StringS("javaagent", "javaagent", "", "load Java programming language agent, see java.lang.instrument")
	rootCmd.Flags().BoolS("server", "server", false, "to select the \"server\" VM")
	rootCmd.Flags().BoolS("showversion", "showversion", false, "print product version and continue")
	rootCmd.Flags().StringS("splash", "splash", "", "show splash screen with specified image")
	rootCmd.Flags().StringS("verbose", "verbose", "", "enable verbose output")
	rootCmd.Flags().BoolS("version", "version", false, "print product version and exit")

	rootCmd.Flags().StringArrayS("XX", "XX", []string{}, "Advanced options")
	rootCmd.Flags().BoolS("Xbatch", "Xbatch", false, "disable background compilation")
	rootCmd.Flags().StringS("Xbootclasspath", "Xbootclasspath", "", "set search path for bootstrap classes and resources")
	rootCmd.Flags().StringS("Xbootclasspath/a", "Xbootclasspath/a", "", "append to end of bootstrap class path")
	rootCmd.Flags().StringS("Xbootclasspath/p", "Xbootclasspath/p", "", "prepend in front of bootstrap class path")
	rootCmd.Flags().StringS("Xcheck", "Xcheck", "", "perform additional checks")
	rootCmd.Flags().BoolS("Xdiag", "Xdiag", false, "show additional diagnostic messages")
	rootCmd.Flags().BoolS("Xfuture", "Xfuture", false, "enable strictest checks, anticipating future default")
	rootCmd.Flags().BoolS("Xincgc", "Xincgc", false, "enable incremental garbage collection")
	rootCmd.Flags().BoolS("Xint", "Xint", false, "interpreted mode execution only")
	rootCmd.Flags().StringS("Xloggc", "Xloggc", "", "log GC status to a file with time stamps")
	rootCmd.Flags().BoolS("Xmixed", "Xmixed", false, "mixed mode execution (default)")
	rootCmd.Flags().StringS("Xms", "Xms", "", "set initial Java heap size")
	rootCmd.Flags().StringS("Xmx", "Xmx", "", "set maximum Java heap size")
	rootCmd.Flags().BoolS("Xnoclassgc", "Xnoclassgc", false, "disable class garbage collection")
	rootCmd.Flags().BoolS("Xprof", "Xprof", false, "output cpu profiling data")
	rootCmd.Flags().BoolS("Xrs", "Xrs", false, "reduce use of OS signals by Java/VM (see documentation)")
	rootCmd.Flags().StringS("Xshare", "Xshare", "", "set shared data usage")
	rootCmd.Flags().StringS("XshowSettings", "XshowSettings", "", "show all settings and continue")
	rootCmd.Flags().StringS("Xss", "Xss", "", "set java thread stack size")

	rootCmd.Flag("D").NoOptDefVal = " "
	rootCmd.Flag("agentlib").NoOptDefVal = " "
	rootCmd.Flag("agentpath").NoOptDefVal = " "
	rootCmd.Flag("da").NoOptDefVal = " "
	rootCmd.Flag("disableassertions").NoOptDefVal = " "
	rootCmd.Flag("ea").NoOptDefVal = " "
	rootCmd.Flag("enableassertions").NoOptDefVal = " "
	rootCmd.Flag("javaagent").NoOptDefVal = " "
	rootCmd.Flag("splash").NoOptDefVal = " "
	rootCmd.Flag("verbose").NoOptDefVal = " "
	rootCmd.Flag("Xbootclasspath").NoOptDefVal = " "
	rootCmd.Flag("Xbootclasspath/a").NoOptDefVal = " "
	rootCmd.Flag("Xbootclasspath/p").NoOptDefVal = " "
	rootCmd.Flag("Xloggc").NoOptDefVal = " "
	rootCmd.Flag("Xms").NoOptDefVal = " "
	rootCmd.Flag("Xmx").NoOptDefVal = " "
	rootCmd.Flag("Xss").NoOptDefVal = " "
	rootCmd.Flag("Xcheck").NoOptDefVal = " "
	rootCmd.Flag("Xshare").NoOptDefVal = " "
	rootCmd.Flag("XshowSettings").NoOptDefVal = " "
	rootCmd.Flag("XX").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"XX": action.ActionAdvanced(),
		"Xbootclasspath": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			return carapace.ActionFiles(".zip", ".jar")
		}),
		"Xbootclasspath/a": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			return carapace.ActionFiles(".zip", ".jar")
		}),
		"Xbootclasspath/p": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			return carapace.ActionFiles(".zip", ".jar")
		}),
		"Xcheck": carapace.ActionValuesDescribed(
			"jni", "perform additional checks for JNI functions",
		),
		"Xloggc": carapace.ActionFiles(),
		"Xshare": carapace.ActionValuesDescribed(
			"auto", "use shared class data if possible",
			"jni", "perform additional checks for JNI functions",
			"off", "do not attempt to use shared class data",
			"on", "require using shared class data, otherwise fail",
		),
		"XshowSettings": carapace.ActionValuesDescribed(
			"all", "show all settings and continue",
			"vm", "show all vm related settings and continue",
			"system", "show host system or container configuration and continue",
			"properties", "show all property settings and continue",
			"locale", "show all locale related settings and continue",
		),
		"agentpath": carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionFiles(".jar", ".zip")
			default:
				return carapace.ActionValues()
			}
		}),
		"classpath": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			return carapace.ActionFiles()
		}),
		"cp": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			return carapace.ActionFiles()
		}),
		"da": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			return action.ActionClasspathClasses(rootCmd)
		}),
		"disableassertions": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			return action.ActionClasspathClasses(rootCmd)
		}),
		"ea": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			return action.ActionClasspathClasses(rootCmd)
		}),
		"enableassertions": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			return action.ActionClasspathClasses(rootCmd)
		}),
		"jar": carapace.ActionFiles(".jar", ".zip"),
		"javaagent": carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionFiles(".jar", ".zip")
			default:
				return carapace.ActionValues()
			}
		}),
		"splash": carapace.ActionFiles(),
		"verbose": carapace.ActionValuesDescribed(
			"class", "Displays information about each loaded class",
			"gc", "Displays information about each garbage collection (GC) event",
			"jni", "Displays information about the use of native methods and other Java Native Interface (JNI) activity",
		),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if !rootCmd.Flag("jar").Changed {
				_, exists := os.LookupEnv("CLASSPATH")
				if !exists && !rootCmd.Flag("classpath").Changed && !rootCmd.Flag("cp").Changed {
					return action.ActionLocalClasses()
				}
				return action.ActionClasspathClasses(rootCmd)
			}
			return carapace.ActionValues()
		}),
	)
}
