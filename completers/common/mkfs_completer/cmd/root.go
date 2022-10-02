package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mkfs",
	Short: "Make a Linux filesystem",
	Long:  "https://linux.die.net/man/8/mkfs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().StringP("type", "t", "", "filesystem type; when unspecified, ext2 is used")
	rootCmd.Flags().Bool("verbose", false, "explain what is being done;")
	rootCmd.Flags().BoolP("version", "V", false, "display version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"type": fs.ActionFilesystemTypes(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.Batch(
			fs.ActionBlockDevices(),
			carapace.ActionFiles(),
		).ToA(),
	)
}
