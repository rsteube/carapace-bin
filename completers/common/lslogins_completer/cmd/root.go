package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/lslogins_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lslogins",
	Short: "Display information about known users in the system",
	Long:  "https://man7.org/linux/man-pages/man1/lslogins.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("acc-expiration", "a", false, "display info about passwords expiration")
	rootCmd.Flags().String("btmp-file", "", "set an alternate path for btmp")
	rootCmd.Flags().BoolP("colon-separate", "c", false, "display data in a format similar to /etc/passwd")
	rootCmd.Flags().BoolP("context", "Z", false, "display SELinux contexts")
	rootCmd.Flags().BoolP("export", "e", false, "display in an export-able output format")
	rootCmd.Flags().BoolP("failed", "f", false, "display data about the users' last failed logins")
	rootCmd.Flags().StringP("groups", "g", "", "display users belonging to a group in <groups>")
	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().BoolP("last", "L", false, "show info about the users' last login sessions")
	rootCmd.Flags().String("lastlog", "", "set an alternate path for lastlog")
	rootCmd.Flags().StringP("logins", "l", "", "display only users from <logins>")
	rootCmd.Flags().BoolP("newline", "n", false, "display each piece of information on a new line")
	rootCmd.Flags().Bool("noheadings", false, "don't print headings")
	rootCmd.Flags().Bool("notruncate", false, "don't truncate output")
	rootCmd.Flags().StringP("output", "o", "", "define the columns to output")
	rootCmd.Flags().Bool("output-all", false, "output all columns")
	rootCmd.Flags().BoolP("print0", "z", false, "delimit user entries with a nul character")
	rootCmd.Flags().BoolP("pwd", "p", false, "display information related to login by password.")
	rootCmd.Flags().BoolP("raw", "r", false, "display in raw mode")
	rootCmd.Flags().BoolP("supp-groups", "G", false, "display information about groups")
	rootCmd.Flags().BoolP("system-accs", "s", false, "display system accounts")
	rootCmd.Flags().String("time-format", "", "display dates in short, full or iso format")
	rootCmd.Flags().BoolP("user-accs", "u", false, "display user accounts")
	rootCmd.Flags().BoolP("version", "V", false, "display version")
	rootCmd.Flags().String("wtmp-file", "", "set an alternate path for wtmp")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"btmp-file": carapace.ActionFiles(),
		"groups": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return os.ActionGroups().Invoke(c).Filter(c.Parts).ToA()
		}),
		"lastlog": carapace.ActionFiles(),
		"logins": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return os.ActionUsers().Invoke(c).Filter(c.Parts).ToA()
		}),
		"output": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionColumns().Invoke(c).Filter(c.Parts).ToA()
		}),
		"time-format": carapace.ActionValues("short", "full", "iso"),
		"wtmp-file":   carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		os.ActionUsers(),
	)
}
