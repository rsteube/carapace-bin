package action

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net/http"
	"github.com/spf13/cobra"
)

func actionApi(cmd *cobra.Command, query string, v interface{}, transform func() carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if flag := cmd.Flag("repo"); flag != nil && flag.Changed {
			project := strings.Join(strings.Split(flag.Value.String(), "/")[1:], "/")
			query = strings.Replace(query, ":fullpath", url.PathEscape(project), 1)
		}

		args := []string{"api"}
		if flag := cmd.Flag("repo"); flag != nil && flag.Changed {
			host := strings.Split(flag.Value.String(), "/")[0]
			args = append(args, "--hostname", host)
		}
		args = append(args, query)
		return carapace.ActionExecCommand("glab", args...)(func(output []byte) carapace.Action {
			if err := json.Unmarshal(output, &v); err != nil {
				return carapace.ActionMessage("failed to unmarshall response: " + err.Error())
			}
			return transform()
		})
	})
}

func actionGraphql(cmd *cobra.Command, query string, v interface{}, transform func() carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"api", "graphql"}
		if flag := cmd.Flag("repo"); flag != nil && flag.Changed {
			host := strings.Split(flag.Value.String(), "/")[0]
			args = append(args, "--hostname", host)
		}
		args = append(args, "-f", fmt.Sprintf("query=%v", query))
		return carapace.ActionExecCommand("glab", args...)(func(output []byte) carapace.Action {
			if err := json.Unmarshal(output, &v); err != nil {
				return carapace.ActionMessage("failed to unmarshall response: " + err.Error())
			}
			return transform()
		})
	})
}

func ActionApiPaths(cmd *cobra.Command) carapace.Action {
	return http.ActionApiPaths(v4paths, ":[^/]+", func(c carapace.Context, matchedData map[string]string, segment string) carapace.Action {
		switch segment {
		// TODO completion for other placeholders
		case ":group_id":
			return ActionGroups(cmd)
		case ":user_id":
			return ActionUsers(cmd)
		default:
			// static value or placeholder not yet handled
			return carapace.ActionValues(segment)
		}
	})
}
