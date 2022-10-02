package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/aws"
	"github.com/spf13/cobra"
)

var registryLoginCmd = &cobra.Command{
	Use:   "registry-login",
	Short: "Generate and save the registry authentication file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	registryLoginCmd.Flags().String("account-id", "", "Your AWS Account id")
	registryLoginCmd.Flags().Bool("ecr", false, "If we are using ECR we need a different set of flags, so if this is set, we need to set --username and --password")
	registryLoginCmd.Flags().String("password", "", "The registry password")
	registryLoginCmd.Flags().BoolP("password-stdin", "s", false, "Reads the docker password from stdin, either pipe to the command or remember to press ctrl+d when reading interactively")
	registryLoginCmd.Flags().String("region", "", "Your AWS region")
	registryLoginCmd.Flags().String("server", "https://index.docker.io/v1/", "The server URL, it is defaulted to the docker registry")
	registryLoginCmd.Flags().StringP("username", "u", "", "The Registry Username")
	rootCmd.AddCommand(registryLoginCmd)

	carapace.Gen(registryLoginCmd).FlagCompletion(carapace.ActionMap{
		"region": aws.ActionRegions(),
	})
}
