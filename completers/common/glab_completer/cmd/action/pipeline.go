package action

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action/utils"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

type pipeline struct {
	Id        int
	Ref       string
	CreatedAt *time.Time `json:"created_at"`
	Status    string
}

func ActionPipelines(cmd *cobra.Command, status string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		statusFilter := ""
		if status != "" {
			statusFilter = "&status=" + url.QueryEscape(status)
		}

		var queryResult []pipeline
		return actionApi(cmd, fmt.Sprintf("projects/:fullpath/pipelines?order_by=updated_at&per_page=100%v", statusFilter), &queryResult, func() carapace.Action {
			vals := make([]string, 0, len(queryResult)*2)
			for _, pipeline := range queryResult {
				var s string
				switch pipeline.Status {
				case "success":
					s = style.Green
				case "failed":
					s = style.Red
				default:
					s = style.Gray
				}

				description := pipeline.Ref
				if pipeline.CreatedAt != nil {
					description = fmt.Sprintf("%v (%v)", pipeline.Ref, utils.TimeToPrettyTimeAgo(*pipeline.CreatedAt))
				}
				vals = append(vals, strconv.Itoa(pipeline.Id), description, s)
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	})
}
