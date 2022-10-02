package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ps",
	Short: "report a snapshot of the current processes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("A", "A", false, "Select all processes")
	rootCmd.Flags().BoolS("F", "F", false, "Extra full format")
	rootCmd.Flags().StringP("Group", "G", "", "Select by real group ID  or name")
	rootCmd.Flags().BoolS("H", "H", false, "Show process hierarchy (forest).")
	rootCmd.Flags().BoolS("L", "L", false, "Show threads, possibly with LWP and NLWP columns.")
	rootCmd.Flags().BoolS("M", "M", false, "Add a column of security data.")
	rootCmd.Flags().BoolS("N", "N", false, "Select all processes except those that fulfill the specified conditions")
	rootCmd.Flags().StringS("O", "O", "", "Like -o, but preloaded with some default columns")
	rootCmd.Flags().BoolS("T", "T", false, "Show threads, possibly with SPID column.")
	rootCmd.Flags().StringP("User", "U", "", "Select by real user ID (RUID) or name")
	rootCmd.Flags().BoolS("V", "V", false, "Print the procps-ng version.")
	rootCmd.Flags().BoolS("a", "a", false, "Select all processes except both session leaders and processes not associated with a terminal")
	rootCmd.Flags().BoolS("c", "c", false, "Show different scheduler information for the -l option.")
	rootCmd.Flags().String("cols", "", "Set screen width.")
	rootCmd.Flags().String("columns", "", "Set screen width.")
	rootCmd.Flags().Bool("context", false, "Display security context format (for SELinux).")
	rootCmd.Flags().Bool("cumulative", false, "Include some dead child process data (as a sum with the parent).")
	rootCmd.Flags().BoolS("d", "d", false, "Select all processes except session leaders")
	rootCmd.Flags().Bool("deselect", false, "Select all processes except those that fulfill the specified conditions")
	rootCmd.Flags().BoolS("e", "e", false, "Select all processes")
	rootCmd.Flags().BoolS("f", "f", false, "Do full-format listing")
	rootCmd.Flags().Bool("forest", false, "ASCII art process tree.")
	rootCmd.Flags().StringP("format", "o", "", "user-defined format")
	rootCmd.Flags().StringP("group", "g", "", "Select by session OR by effective group name")
	rootCmd.Flags().Bool("headers", false, "Repeat header lines, one per page of output.")
	rootCmd.Flags().String("help", "", "Print a help message")
	rootCmd.Flags().Bool("info", false, "Print debugging info.")
	rootCmd.Flags().BoolS("j", "j", false, "Jobs format.")
	rootCmd.Flags().BoolS("l", "l", false, "Long format.  The -y option is often useful with this.")
	rootCmd.Flags().String("lines", "", "Set screen height.")
	rootCmd.Flags().BoolS("m", "m", false, "Show threads after processes.")
	rootCmd.Flags().Bool("no-headers", false, "Print no header line at all")
	rootCmd.Flags().StringP("pid", "p", "", "Select by PID")
	rootCmd.Flags().String("ppid", "", "Select by parent process ID")
	rootCmd.Flags().StringP("quick-pid", "q", "", "Select by process ID (quick mode)")
	rootCmd.Flags().String("rows", "", "Set screen height.")
	rootCmd.Flags().StringS("s", "s", "", "Select by session ID")
	rootCmd.Flags().String("sid", "", "Select by session ID")
	rootCmd.Flags().String("sort", "", "Specify sorting order")
	rootCmd.Flags().StringP("tty", "t", "", "Select by tty")
	rootCmd.Flags().StringP("user", "u", "", "Select by effective user ID (EUID) or name")
	rootCmd.Flags().Bool("version", false, "Print the procps-ng version.")
	rootCmd.Flags().BoolS("w", "w", false, "Wide output")
	rootCmd.Flags().String("width", "", "Set screen width.")
	rootCmd.Flags().BoolS("y", "y", false, "Do not show flags")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"Group": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return os.ActionGroups().Invoke(c).Filter(c.Parts).ToA()
		}),
		"User": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return os.ActionUsers().Invoke(c).Filter(c.Parts).ToA()
		}),
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return ps.ActionFormatSpecifiers().Invoke(c).Filter(c.Parts).ToA()
		}),
		"group": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return os.ActionGroups().Invoke(c).Filter(c.Parts).ToA()
		}),
		"pid": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return ps.ActionProcessIds().Invoke(c).Filter(c.Parts).ToA()
		}),
		"ppid": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return ps.ActionProcessIds().Invoke(c).Filter(c.Parts).ToA()
		}),
		"quick-pid": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return ps.ActionProcessIds().Invoke(c).Filter(c.Parts).ToA()
		}),
		"sid": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return os.ActionSessionIds().Invoke(c).Filter(c.Parts).ToA()
		}),
		"sort": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return ps.ActionFormatSpecifiers().Invoke(c).Filter(c.Parts).ToA()
		}),
		"tty": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return os.ActionTerminals().Invoke(c).Filter(c.Parts).ToA()
		}),
		"user": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return os.ActionUsers().Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
