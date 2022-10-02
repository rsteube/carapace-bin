package action

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/styles"
	"github.com/spf13/cobra"
)

func ActionKeywordLinks(cmd *cobra.Command) carapace.Action {
	return carapace.ActionMultiParts(" ", func(c carapace.Context) carapace.Action {
		if strings.HasPrefix(c.CallbackValue, "@") {
			c.CallbackValue = strings.TrimPrefix(c.CallbackValue, "@")
			return ActionMentionableUsers(cmd).Invoke(c).Prefix("@").ToA()
		}

		keywordsOfficial := []string{"close", "closes", "closed", "fix", "fixes", "fixed", "resolve", "resolves", "resolved"}
		keywordsCustom := []string{"causes", "caused-by", "closed_by", "fixed_by", "related", "resolved_by", "see"}

		re := regexp.MustCompile(fmt.Sprintf("^(%v)$", strings.Join(append(keywordsOfficial, keywordsCustom...), "|")))
		if last := len(c.Parts) - 1; last < 0 || !re.MatchString(c.Parts[last]) {
			return carapace.Batch(
				carapace.ActionValues(keywordsOfficial...).Style(styles.Gh.StateClosed),
				carapace.ActionValues(keywordsCustom...),
			).ToA().Invoke(c).Suffix(" #").ToA()
		}

		return carapace.ActionMultiParts("#", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				if strings.Contains(c.CallbackValue, "/") {
					return ActionOwnerRepositories(cmd).Invoke(c).Suffix("#").ToA()
				}
				return ActionOwnerRepositories(cmd)
			case 1:
				if repo := c.Parts[0]; repo != "" {
					if err := cmd.Flags().Set("repo", repo); err != nil {
						return carapace.ActionMessage(err.Error())
					}
				}
				return carapace.Batch(
					ActionIssues(cmd, IssueOpts{Open: true, Closed: true}),
					ActionPullRequests(cmd, PullRequestOpts{Open: true, Closed: true, Merged: true}),
				).ToA()
			default:
				return carapace.ActionValues()
			}
		})
	})
}
