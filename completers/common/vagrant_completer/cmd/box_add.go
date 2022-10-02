package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

var box_addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a box",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(box_addCmd).Standalone()

	box_addCmd.Flags().String("box-version", "", "Constrain version of the added box")
	box_addCmd.Flags().String("cacert", "", "CA certificate for SSL download")
	box_addCmd.Flags().String("capath", "", "CA certificate directory for SSL download")
	box_addCmd.Flags().String("cert", "", "A client SSL cert, if needed")
	box_addCmd.Flags().String("checksum", "", "Checksum for the box")
	box_addCmd.Flags().String("checksum-type", "", "Checksum type (md5, sha1, sha256)")
	box_addCmd.Flags().BoolP("clean", "c", false, "Clean any temporary download files")
	box_addCmd.Flags().BoolP("force", "f", false, "Overwrite an existing box if it exists")
	box_addCmd.Flags().Bool("insecure", false, "Do not validate SSL certificates")
	box_addCmd.Flags().Bool("location-trusted", false, "Trust 'Location' header from HTTP redirects and use the same credentials for subsequent urls as for the initial one")
	box_addCmd.Flags().String("name", "", "Name of the box")
	box_addCmd.Flags().String("provider", "", "Provider the box should satisfy")
	boxCmd.AddCommand(box_addCmd)

	carapace.Gen(box_addCmd).FlagCompletion(carapace.ActionMap{
		"cacert":        carapace.ActionFiles(),
		"capath":        carapace.ActionDirectories(),
		"cert":          carapace.ActionFiles(),
		"checksum-type": carapace.ActionValues("md5", "sha1", "sha256"),
		"provider":      action.ActionProviders(),
	})

	carapace.Gen(box_addCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.CallbackValue) {
				return carapace.ActionFiles(".box", ".json")
			}
			return action.ActionCloudBoxSearch(box_addCmd.Flag("provider").Value.String())
		}),
	)
}
