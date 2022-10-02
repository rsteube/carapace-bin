package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cloud_publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Create and release a new Vagrant Box on Vagrant Cloud",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_publishCmd).Standalone()

	cloud_publishCmd.Flags().String("box-version", "", "Version of box to create")
	cloud_publishCmd.Flags().StringP("checksum", "c", "", "Checksum of the box for this provider. --checksum-type option is required.")
	cloud_publishCmd.Flags().StringP("checksum-type", "C", "", "Type of checksum used (md5, sha1, sha256, sha384, sha512). --checksum option is required.")
	cloud_publishCmd.Flags().StringP("description", "d", "", "Full description of box")
	cloud_publishCmd.Flags().Bool("direct-upload", false, "Upload asset directly to backend storage")
	cloud_publishCmd.Flags().BoolP("force", "f", false, "Disables confirmation to create or update box")
	cloud_publishCmd.Flags().Bool("no-direct-upload", false, "Do not upload asset directly to backend storage")
	cloud_publishCmd.Flags().Bool("no-force", false, "Enables confirmation to create or update box")
	cloud_publishCmd.Flags().Bool("no-private", false, "Makes box public")
	cloud_publishCmd.Flags().Bool("no-release", false, "Does not elease box")
	cloud_publishCmd.Flags().BoolP("private", "p", false, "Makes box private")
	cloud_publishCmd.Flags().BoolP("release", "r", false, "Releases box")
	cloud_publishCmd.Flags().StringP("short-description", "s", "", "Short description of the box")
	cloud_publishCmd.Flags().String("url", "", "Remote URL to download this provider (cannot be used with provider-file)")
	cloud_publishCmd.Flags().String("version-description", "", "Description of the version to create")
	cloudCmd.AddCommand(cloud_publishCmd)

	carapace.Gen(cloud_publishCmd).FlagCompletion(carapace.ActionMap{
		"checksum-type": carapace.ActionValues("md5", "sha1", "sha256", "sha384", "sha512"),
	})

	carapace.Gen(cloud_publishCmd).PositionalCompletion(
		action.ActionCloudBoxSearch(""),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionCloudBoxVersions(c.Args[0], "")
		}),
		action.ActionProviders(),
		carapace.ActionFiles(".box", ".json"),
	)
}
