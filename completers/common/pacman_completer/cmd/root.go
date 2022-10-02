package cmd

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pacman",
	Short: "package manager utility",
	Long:  "https://wiki.archlinux.de/title/Pacman",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PreRun(func(cmd *cobra.Command, args []string) {
		if len(args) < 1 || (len(args) == 1 && regexp.MustCompile("^-$|^--").MatchString(args[0])) {
			rootCmd.Flags().BoolP("database", "D", false, "Operate on the package database")
			rootCmd.Flags().BoolP("deptest", "T", false, "Check dependencies")
			rootCmd.Flags().BoolP("files", "F", false, "Query the files database")
			rootCmd.Flags().BoolP("help", "h", false, "Display syntas for the given operation")
			rootCmd.Flags().BoolP("query", "Q", false, "Query the package database")
			rootCmd.Flags().BoolP("remove", "R", false, "Remove packages from the system")
			rootCmd.Flags().BoolP("sync", "S", false, "Synchronize packages")
			rootCmd.Flags().BoolP("upgrade", "U", false, "Upgrade or add packages")
			rootCmd.Flags().BoolP("version", "V", false, "Display version")
		} else {
			subCmd := args[0]
			if !strings.HasPrefix(subCmd, "--") && len(subCmd) > 1 {
				subCmd = subCmd[:2] // assume shorthand flag
			}
			switch subCmd {
			case "--database", "-D":
				rootCmd.Flags().BoolP("database", "D", false, "Operate on the package database")
				initDatabaseCmd(rootCmd)
			case "--deptest", "-T":
				rootCmd.Flags().BoolP("deptest", "T", false, "Check dependencies")
				initDeptestCmd(rootCmd)
			case "--files", "-F":
				rootCmd.Flags().BoolP("files", "F", false, "Query the files database")
				initFilesCmd(rootCmd)
			case "--help", "-h":
				rootCmd.Flags().BoolP("help", "h", false, "Display syntas for the given operation")
			case "--query", "-Q":
				rootCmd.Flags().BoolP("query", "Q", false, "Query the package database")
				initQueryCmd(rootCmd)
			case "--remove", "-R":
				rootCmd.Flags().BoolP("remove", "R", false, "Remove packages from the system")
				initRemoveCmd(rootCmd)
			case "--sync", "-S":
				rootCmd.Flags().BoolP("sync", "S", false, "Synchronize packages")
				initSyncCmd(rootCmd)
			case "--upgrade", "-U":
				rootCmd.Flags().BoolP("upgrade", "U", false, "Upgrade or add packages")
				initUpgradeCmd(rootCmd)
			case "--version", "-V":
				rootCmd.Flags().BoolP("version", "V", false, "Display version")
			}
		}
	})
}
