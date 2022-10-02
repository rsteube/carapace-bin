package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

var cpCmd = &cobra.Command{
	Use:   "cp",
	Short: "Copy files/folders between a service container and the local filesystem",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cpCmd).Standalone()
	cpCmd.Flags().Bool("all", false, "Copy to all the containers of the service.")
	cpCmd.Flags().BoolP("archive", "a", false, "Archive mode (copy all uid/gid information)")
	cpCmd.Flags().BoolP("follow-link", "L", false, "Always follow symbol link in SRC_PATH")
	cpCmd.Flags().Int("index", 1, "Index of the container if there are multiple instances of a service [default: 1].")
	rootCmd.AddCommand(cpCmd)

	carapace.Gen(cpCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.CallbackValue) {
				return carapace.ActionFiles()
			}
			return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return action.ActionServices(cpCmd).Invoke(c).Suffix(":").ToA()
				case 1:
					if index, err := cpCmd.Flags().GetInt("index"); err != nil {
						return carapace.ActionMessage(err.Error())
					} else {
						return action.ActionFiles(cpCmd, c.Parts[0], index)
					}
				default:
					return carapace.ActionValues()
				}
			})
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if !util.HasPathPrefix(c.Args[0]) {
				return carapace.ActionFiles()
			}
			return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return action.ActionServices(cpCmd).Invoke(c).Suffix(":").ToA()
				case 1:
					if index, err := cpCmd.Flags().GetInt("index"); err != nil {
						return carapace.ActionMessage(err.Error())
					} else {
						return action.ActionFiles(cpCmd, c.Parts[0], index)
					}
				default:
					return carapace.ActionValues()
				}
			})
		}),
	)
}
