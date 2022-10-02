package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "List images used by the created containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagesCmd).Standalone()
	imagesCmd.Flags().BoolP("quiet", "q", false, "Only display IDs")
	rootCmd.AddCommand(imagesCmd)

	carapace.Gen(imagesCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionServices(imagesCmd).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
