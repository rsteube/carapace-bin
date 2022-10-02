package action

import (
	"fmt"
	"net/url"
	"path/filepath"
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/styles"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

type release struct {
	TagName   string    `json:"tag_name"`
	CreatedAt time.Time `json:"created_at"`
	Assets    struct {
		Sources []struct {
			Format string
			Url    string
		}
	}
}

func ActionReleases(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult []release
		return actionApi(cmd, "projects/:fullpath/releases?per_page=100", &queryResult, func() carapace.Action {
			vals := make([]string, 0, len(queryResult)*2)
			for _, release := range queryResult {
				vals = append(vals, release.TagName, util.FuzzyAgo(time.Since(release.CreatedAt)))
			}
			return carapace.ActionValuesDescribed(vals...).Style(styles.Git.Tag)
		})
	})
}

func ActionReleaseAssets(cmd *cobra.Command, id string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult release
		return actionApi(cmd, fmt.Sprintf("projects/:fullpath/releases/%v", url.PathEscape(id)), &queryResult, func() carapace.Action {
			vals := make([]string, 0)
			for _, source := range queryResult.Assets.Sources {
				vals = append(vals, filepath.Base(source.Url), source.Format)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
