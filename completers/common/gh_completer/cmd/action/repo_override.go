package action

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

func ActionRepoOverride(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// TODO support github enterprise (different host)
		actions := make([]carapace.Action, 0)
		if _, err := util.FindReverse(c.Dir, ".git"); err == nil {
			actions = append(actions, actionRemoteRepositories())
		}
		if c.CallbackValue != "" || len(actions) == 0 {
			actions = append(actions, ActionOwnerRepositories(cmd))
		}
		return carapace.Batch(actions...).Invoke(c).Merge().ToA()
	})
}

func actionRemoteRepositories() carapace.Action {
	return carapace.ActionExecCommand("git", "remote", "--verbose")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^(?P<remote>[^ ]+)\t+(?P<repo>(?P<url>[^ ]*?))(\.git)? (?P<type>.*)$`)

		urls := make(map[string]string)
		for _, line := range lines[:len(lines)-1] {
			if r.MatchString(line) {
				matches := r.FindStringSubmatch(line)
				urls[matches[3]] = matches[1]
			}
		}

		vals := make([]string, 0)
		for url, remote := range urls {
			if strings.HasPrefix(url, "https://github.com/") {
				vals = append(vals, strings.TrimPrefix(url, "https://github.com/"), remote)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
