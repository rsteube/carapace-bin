package cmd

import (
	"github.com/spf13/cobra"
)

var show_refCmd = &cobra.Command{
	Use:   "show-ref",
	Short: "List references in a local repository",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	show_refCmd.Flags().String("abbrev", "", "use <n> digits to display SHA-1s")
	show_refCmd.Flags().BoolP("dereference", "d", false, "dereference tags into object IDs")
	show_refCmd.Flags().String("exclude-existing", "", "show refs from stdin that aren't in local repository")
	show_refCmd.Flags().StringP("hash", "s", "", "only show SHA1 hash using <n> digits")
	show_refCmd.Flags().Bool("head", false, "show the HEAD reference, even if it would be filtered out")
	show_refCmd.Flags().Bool("heads", false, "only show heads (can be combined with tags)")
	show_refCmd.Flags().BoolP("quiet", "q", false, "do not print results to stdout (useful with --verify)")
	show_refCmd.Flags().Bool("tags", false, "only show tags (can be combined with heads)")
	show_refCmd.Flags().Bool("verify", false, "stricter reference checking, requires exact ref path")
	rootCmd.AddCommand(show_refCmd)
}
