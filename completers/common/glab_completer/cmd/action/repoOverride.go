package action

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

func FakeRepoFlag(cmd *cobra.Command, value string) {
	cmd.Flags().StringP("repo", "R", value, "fake repo flag")
	cmd.Flag("repo").Changed = true
}

func ActionRepo(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if flag := cmd.Flag("repo"); flag == nil {
			FakeRepoFlag(cmd, c.CallbackValue)
		}

		configHosts, err := hosts()
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		if strings.Contains(c.CallbackValue, "/") {
			isKnownHost := false
			for _, host := range configHosts {
				if strings.HasPrefix(c.CallbackValue, host) {
					isKnownHost = true
					break
				}
			}
			if !isKnownHost {
				return carapace.ActionValues() // only complete full host style for simplicity
			}
		}

		return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues(configHosts...).Invoke(c).Suffix("/").ToA()
			case 1:
				return carapace.Batch(
					ActionUsers(cmd).Suppress("search needs at least 3 characters"),
					ActionGroups(cmd),
				).Invoke(c).Merge().Suffix("/").ToA()
			case 2:
				b := carapace.Batch(
					ActionSubgroups(cmd, c.Parts[1]).Suppress("failed to unmarshall"),
					ActionGroupProjects(cmd, c.Parts[1]).Suppress("failed to unmarshall"),
					ActionUserProjects(cmd, c.Parts[1]).Suppress("failed to unmarshall"),
				).Invoke(c)
				return b[0].Suffix("/").Merge(b[1:]...).ToA() // ActionSubgroups needs `/` suffix
			case 3:
				return ActionGroupProjects(cmd, strings.Join(c.Parts[1:], "/")).Suppress("failed to unmarshall")
			default:
				return carapace.ActionValues()
			}
		})
	})
}
