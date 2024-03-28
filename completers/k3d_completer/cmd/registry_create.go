package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var registry_createCmd = &cobra.Command{
	Use:   "create NAME",
	Short: "Create a new registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_createCmd).Standalone()

	registry_createCmd.Flags().StringSliceP("cluster", "c", []string{}, "[NotReady] Select the cluster(s) that the registry shall connect to.")
	registry_createCmd.Flags().String("default-network", "", "Specify the network connected to the registry")
	registry_createCmd.Flags().StringP("image", "i", "", "Specify image used for the registry")
	registry_createCmd.Flags().Bool("no-help", false, "Disable the help text (How-To use the registry)")
	registry_createCmd.Flags().StringP("port", "p", "", "Select which port the registry should be listening on on your machine (localhost) (Format: `[HOST:]HOSTPORT`)")
	registry_createCmd.Flags().String("proxy-password", "", "Specify the password of the proxied remote registry")
	registry_createCmd.Flags().String("proxy-remote-url", "", "Specify the url of the proxied remote registry")
	registry_createCmd.Flags().String("proxy-username", "", "Specify the username of the proxied remote registry")
	registry_createCmd.Flags().StringSliceP("volume", "v", []string{}, "Mount volumes into the registry node (Format: `[SOURCE:]DEST`")
	registry_createCmd.Flag("cluster").Hidden = true
	registryCmd.AddCommand(registry_createCmd)
}
