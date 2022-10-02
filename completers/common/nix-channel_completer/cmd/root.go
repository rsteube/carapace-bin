package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nix-channel",
	Short: "manage Nix channels",
	Long:  "https://nixos.org/manual/nix/stable/command-ref/nix-channel.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("add", false, "Adds a channel named name with URL url to the list  of subscribed channels")
	rootCmd.Flags().Bool("help", false, "show help")
	rootCmd.Flags().Bool("list", false, "Prints the names and URLs of all subscribed channels on standard output")
	rootCmd.Flags().Bool("remove", false, "Removes the channel named name from the list of sub‐ scribed channels")
	rootCmd.Flags().Bool("rollback", false, "Reverts the previous call to nix-channel --update")
	rootCmd.Flags().Bool("update", false, "Downloads the Nix expressions of all subscribed channels")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if (rootCmd.Flag("remove").Changed && len(c.Args) == 0) ||
				rootCmd.Flag("update").Changed {
				return nix.ActionLocalChannels().Invoke(c).Filter(c.Args).ToA()
			} else if rootCmd.Flag("add").Changed {
				switch len(c.Args) {
				case 0:
					return nix.ActionRemoteChannels()
				case 1:
					return nix.ActionLocalChannels()
				}
			}
			return carapace.ActionValues()
		}),
	)
}
