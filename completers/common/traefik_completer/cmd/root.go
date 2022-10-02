package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/traefik_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "traefik",
	Short:              "Traefik is a modern HTTP reverse proxy and load balancer made to deploy microservices with ease",
	Long:               "https://traefik.io/",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	carapace.Gen(rootCmd).PositionalAnyCompletion(
		action.ActionParameters(),
	)
}
