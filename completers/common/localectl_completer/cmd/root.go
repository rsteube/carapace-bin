package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "localectl",
	Short: "Query or change system locale and keyboard settings",
	Long:  "https://www.man7.org/linux/man-pages/man1/localectl.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Show this help")
	rootCmd.Flags().StringP("host", "H", "", "Operate on remote host")
	rootCmd.Flags().StringP("machine", "M", "", "Operate on local container")
	rootCmd.Flags().Bool("no-ask-password", false, "Do not prompt for password")
	rootCmd.Flags().Bool("no-convert", false, "Don't convert keyboard mappings")
	rootCmd.Flags().Bool("no-pager", false, "Do not pipe output into a pager")
	rootCmd.Flags().Bool("version", false, "Show package version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"host": carapace.ActionMultiParts("@", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.Batch(
					os.ActionUsers().Invoke(c).Suffix("@").ToA(),
					net.ActionHosts(),
				).ToA()
			case 1:
				return net.ActionHosts()
			default:
				return carapace.ActionValues()
			}
		}),
	})
}
