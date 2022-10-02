package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fc-list",
	Short: "list available fonts",
	Long:  "https://linux.die.net/man/1/fc-list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("brief", "b", false, "display entire font pattern briefly")
	rootCmd.Flags().StringP("format", "f", "", "use the given output format")
	rootCmd.Flags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().BoolP("quiet", "q", false, "suppress all normal output, exit 1 if no fonts matched")
	rootCmd.Flags().BoolP("verbose", "v", false, "display entire font pattern verbosely")
	rootCmd.Flags().BoolP("version", "V", false, "display font config version and exit")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues(":lang=")
			case 1:
				return os.ActionLanguages()
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
