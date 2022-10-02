package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/fs"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mkswap",
	Short: "Set up a Linux swap area",
	Long:  "https://linux.die.net/man/8/mkswap",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("check", "c", false, "check bad blocks before creating the swap area")
	rootCmd.Flags().BoolP("force", "f", false, "allow swap size area be larger than device")
	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().StringP("label", "L", "", "specify label")
	rootCmd.Flags().String("lock", "", "use exclusive device lock (yes, no or nonblock)")
	rootCmd.Flags().StringP("pagesize", "p", "", "specify page size in bytes")
	rootCmd.Flags().StringP("swapversion", "v", "", "specify swap-space version number")
	rootCmd.Flags().StringP("uuid", "U", "", "specify the uuid to use")
	rootCmd.Flags().Bool("verbose", false, "verbose output")
	rootCmd.Flags().BoolP("version", "V", false, "display version")

	rootCmd.Flag("lock").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"lock": carapace.ActionValues("yes", "no", "nonblock").StyleF(style.ForKeyword),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		fs.ActionBlockDevices(),
	)
}
