package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ncdu",
	Short: "NCurses Disk Usage",
	Long:  "https://linux.die.net/man/1/ncdu",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("0", "0", false, "UI to use when scanning")
	rootCmd.Flags().BoolS("1", "1", false, "UI to use when scanning")
	rootCmd.Flags().BoolS("2", "2", false, "UI to use when scanning")
	rootCmd.Flags().BoolS("V", "V", false, "Print version")
	rootCmd.Flags().String("color", "", "Set color scheme")
	rootCmd.Flags().Bool("confirm-quit", false, "Confirm quitting ncdu")
	rootCmd.Flags().BoolS("e", "e", false, "Enable extended information")
	rootCmd.Flags().String("exclude", "", "Exclude files that match PATTERN")
	rootCmd.Flags().Bool("exclude-caches", false, "Exclude directories containing CACHEDIR.TAG")
	rootCmd.Flags().StringP("exclude-from", "X", "", "Exclude files that match any pattern in FILE")
	rootCmd.Flags().Bool("exclude-kernfs", false, "Exclude Linux pseudo filesystems")
	rootCmd.Flags().StringS("f", "f", "", "Import scanned directory from FILE")
	rootCmd.Flags().BoolP("follow-symlinks", "L", false, "Follow symbolic links")
	rootCmd.Flags().BoolP("help", "h", false, "This help message")
	rootCmd.Flags().StringS("o", "o", "", "Export scanned directory to FILE")
	rootCmd.Flags().BoolS("q", "q", false, "Quiet mode, refresh interval 2 seconds")
	rootCmd.Flags().BoolS("r", "r", false, "Read only")
	rootCmd.Flags().Bool("si", false, "Use base 10 (SI) prefixes instead of base 2")
	rootCmd.Flags().BoolP("version", "v", false, "Print version")
	rootCmd.Flags().BoolS("x", "x", false, "Same filesystem")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color":        carapace.ActionValues("off", "dark"),
		"exclude-from": carapace.ActionFiles(),
		"f":            carapace.ActionFiles(),
		"o":            carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
