package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/code_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "code",
	Short: "Visual Studio Code",
	Long:  "https://code.visualstudio.com/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func ExecuteInsiders() error {
	rootCmd.Use = "code-insiders"
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("add", "a", "", "Add folder(s) to the last active window.")
	rootCmd.Flags().String("category", "", "Filters installed extensions by provided category, when using --list-extensions.")
	rootCmd.Flags().Bool("diff", false, "Compare two files with each other.")
	rootCmd.Flags().String("disable-extension", "", "Disable an extension.")
	rootCmd.Flags().Bool("disable-extensions", false, "Disable all installed extensions.")
	rootCmd.Flags().Bool("disable-gpu", false, "Disable GPU hardware acceleration.")
	rootCmd.Flags().String("enable-proposed-api", "", "Enables proposed API features for extensions.")
	rootCmd.Flags().String("extensions-dir", "", "Set the root path for extensions.")
	rootCmd.Flags().StringP("goto", "g", "", "Open a file at the path on the specified line and character position.")
	rootCmd.Flags().BoolP("help", "h", false, "Print usage.")
	rootCmd.Flags().String("inspect-brk-extensions", "", "Allow debugging and profiling of extensions with the extension host being paused after start.")
	rootCmd.Flags().String("inspect-extensions", "", "Allow debugging and profiling of extensions.")
	rootCmd.Flags().String("install-extension", "", "Installs or updates the extension.")
	rootCmd.Flags().Bool("list-extensions", false, "List the installed extensions.")
	rootCmd.Flags().String("locale", "", "The locale to use (e.g. en-US or zh-TW).")
	rootCmd.Flags().String("log", "", "Log level to use.")
	rootCmd.Flags().String("max-memory", "", "Max memory size for a window (in Mbytes).")
	rootCmd.Flags().BoolP("new-window", "n", false, "Force to open a new window.")
	rootCmd.Flags().Bool("prof-startup", false, "Run CPU profiler during startup.")
	rootCmd.Flags().BoolP("reuse-window", "r", false, "Force to open a file or folder in an already opened window.")
	rootCmd.Flags().Bool("show-versions", false, "Show versions of installed extensions, when using --list-extensions.")
	rootCmd.Flags().BoolP("status", "s", false, "Print process usage and diagnostics information.")
	rootCmd.Flags().String("sync", "", "Turn sync on or off.")
	rootCmd.Flags().Bool("telemetry", false, "Shows all telemetry events which VS code collects.")
	rootCmd.Flags().String("uninstall-extension", "", "Uninstalls an extension.")
	rootCmd.Flags().String("user-data-dir", "", "Specifies the directory that user data is kept in.")
	rootCmd.Flags().Bool("verbose", false, "Print verbose output (implies --wait).")
	rootCmd.Flags().BoolP("version", "v", false, "Print version.")
	rootCmd.Flags().BoolP("wait", "w", false, "Wait for the files to be closed before returning.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"add":               carapace.ActionDirectories(),
		"category":          carapace.ActionValues("Programming Languages", "Snippets", "Linters", "Themes", "Debuggers", "Other"),
		"disable-extension": action.ActionExtensions(rootCmd),
		"extensions-dir":    carapace.ActionDirectories(),
		"goto": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionFiles()
			default:
				return carapace.ActionValues()
			}
		}),
		"install-extension": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.CallbackValue) {
				return carapace.ActionFiles(".vsix")
			}
			return action.ActionExtensionSearch(rootCmd.Flag("category").Value.String())
		}),
		"locale":              os.ActionLocales(),
		"log":                 carapace.ActionValues("critical", "error", "warn", "info", "debug", "trace", "off").StyleF(style.ForLogLevel),
		"sync":                carapace.ActionValues("on", "off"),
		"uninstall-extension": action.ActionExtensions(rootCmd),
		"user-data-dir":       carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("diff").Changed &&
				len(c.Args) > 1 {
				return carapace.ActionValues()
			}
			return carapace.ActionFiles()
		}),
	)
}
