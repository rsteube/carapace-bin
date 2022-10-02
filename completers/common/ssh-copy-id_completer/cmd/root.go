package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/rsteube/carapace-bin/pkg/actions/net/ssh"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ssh-copy-id",
	Short: "use locally available keys to authorise logins on a remote machine",
	Long:  "https://linux.die.net/man/1/ssh-copy-id",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("?", "?", false, "Print Usage summary")
	rootCmd.Flags().BoolS("f", "f", false, "Forced mode")
	rootCmd.Flags().BoolS("h", "h", false, "Print Usage summary")
	rootCmd.Flags().StringS("i", "i", "", "Use only the key(s) contained in identity_file")
	rootCmd.Flags().BoolS("n", "n", false, "do a dry-run")
	rootCmd.Flags().StringS("o", "o", "", "ssh option")
	rootCmd.Flags().StringS("p", "p", "", "ssh port")
	rootCmd.Flags().BoolS("s", "s", false, "SFTP mode")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"i": carapace.Batch(
			ssh.ActionPrivateKeys(),
			carapace.ActionFiles(),
		).ToA(),
		"o": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return ssh.ActionOptions()
		}),
		"p": net.ActionPorts(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionMultiParts("@", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return os.ActionUsers().Invoke(c).Suffix("@").ToA()
			case 1:
				return net.ActionHosts()
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
