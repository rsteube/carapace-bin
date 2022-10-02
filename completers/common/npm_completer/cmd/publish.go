package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(publishCmd).Standalone()
	publishCmd.Flags().String("access", "restricted", "publish as public or restricted")
	publishCmd.Flags().Bool("dry-run", false, "only report changes")
	publishCmd.Flags().String("otp", "", "one-time password")
	publishCmd.Flags().String("tag", "latest", "register with given tag")
	addWorkspaceFlags(publishCmd)

	rootCmd.AddCommand(publishCmd)

	carapace.Gen(publishCmd).FlagCompletion(carapace.ActionMap{
		"access": carapace.ActionValues("public", "restricted"),
	})

	carapace.Gen(publishCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
