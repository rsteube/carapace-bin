package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var image_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_lsCmd).Standalone()

	image_lsCmd.Flags().BoolP("all", "a", false, "Show all images (default hides intermediate images)")
	image_lsCmd.Flags().Bool("digests", false, "Show digests")
	image_lsCmd.Flags().StringArrayP("filter", "f", []string{}, "Filter output based on conditions provided")
	image_lsCmd.Flags().String("format", "", "Pretty-print images using a Go template")
	image_lsCmd.Flags().Bool("no-trunc", false, "Don't truncate output")
	image_lsCmd.Flags().BoolP("quiet", "q", false, "Only show numeric IDs")
	imageCmd.AddCommand(image_lsCmd)

	rootAlias(image_lsCmd, func(cmd *cobra.Command, isAlias bool) {
		if isAlias {
			cmd.Use = "images"
		}

		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			"filter": carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValuesDescribed(
						"dangling", "filter images by dangling state",
						"label", "filter images by label",
						"before", "filter images created before given id or references",
						"since", "filter images created since given id or references",
						"reference", "filter images whose reference matches the specified pattern",
					).Invoke(c).Suffix("=").ToA()
				case 1:
					switch c.Parts[0] {
					case "dangling":
						return carapace.ActionValues("true", "false")
					case "before":
						return docker.ActionRepositoryTags()
					case "since":
						return docker.ActionRepositoryTags()
					default:
						return carapace.ActionValues()
					}
				default:
					return carapace.ActionValues()
				}
			}),
		})

		carapace.Gen(cmd).PositionalCompletion(
			docker.ActionRepositoryTags(),
		)
	})
}
