package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var secret_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a secret from a file or STDIN as content",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secret_createCmd).Standalone()

	secret_createCmd.Flags().StringP("driver", "d", "", "Secret driver")
	secret_createCmd.Flags().StringP("label", "l", "", "Secret labels")
	secret_createCmd.Flags().String("template-driver", "", "Template driver")
	secretCmd.AddCommand(secret_createCmd)

	carapace.Gen(secret_createCmd).PositionalCompletion(
		docker.ActionSecrets(),
		carapace.ActionFiles(),
	)
}
