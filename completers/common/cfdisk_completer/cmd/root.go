package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/fs"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cfdisk",
	Short: "display or manipulate a disk partition table",
	Long:  "https://en.wikipedia.org/wiki/Cfdisk",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("color", "L", "", "colorize output")
	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().String("lock", "", "use exclusive device lock")
	rootCmd.Flags().BoolP("version", "V", false, "display version")
	rootCmd.Flags().BoolP("zero", "z", false, "start with zeroed partition table")

	rootCmd.Flag("color").NoOptDefVal = " "
	rootCmd.Flag("lock").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color": carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"lock":  carapace.ActionValues("yes", "no", "nonblock").StyleF(style.ForKeyword),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.Batch(
			fs.ActionBlockDevices(),
			carapace.ActionFiles(),
		).ToA(),
	)
}
