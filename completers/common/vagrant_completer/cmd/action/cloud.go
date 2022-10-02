package action

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"

	"github.com/rsteube/carapace"
)

type box struct {
	Tag         string
	Description string
}

type boxQuery struct {
	Boxes []box
}

func ActionCloudBoxSearch(provider string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.Batch(
			actionCloudBoxSearch(provider, 1),
			actionCloudBoxSearch(provider, 2),
			actionCloudBoxSearch(provider, 3),
			actionCloudBoxSearch(provider, 4),
			actionCloudBoxSearch(provider, 5),
		).Invoke(c).Merge().ToA()
	})
}

func actionCloudBoxSearch(provider string, page int) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		query := make([]string, 0)
		if provider != "" {
			query = append(query, "provider="+url.QueryEscape(provider))
		}
		query = append(query, "q="+url.QueryEscape(c.CallbackValue))
		query = append(query, "page="+strconv.Itoa(page))

		return carapace.ActionExecCommand("curl", "https://app.vagrantup.com/api/v1/search?"+strings.Join(query, "&"))(func(output []byte) carapace.Action {
			var result boxQuery
			if err := json.Unmarshal(output, &result); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0)
			for _, box := range result.Boxes {
				vals = append(vals, box.Tag, box.Description)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

type boxDetails struct {
	CurrentVersion struct {
		Providers []struct {
			Name string
		}
	} `json:"current_version"`
	Versions []struct {
		Version   string
		Providers []struct {
			Name string
		}
	}
}

func ActionCloudBoxProviders(name string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("curl", "https://app.vagrantup.com/api/v1/box/"+name)(func(output []byte) carapace.Action {
			var details boxDetails
			if err := json.Unmarshal(output, &details); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0)
			for _, provider := range details.CurrentVersion.Providers {
				vals = append(vals, provider.Name)
			}
			for _, version := range details.Versions {
				for _, provider := range version.Providers {
					vals = append(vals, provider.Name)
				}
			}
			return carapace.ActionValues(vals...)
		})
	})
}

func ActionCloudBoxVersions(name string, provider string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("curl", "https://app.vagrantup.com/api/v1/box/"+name)(func(output []byte) carapace.Action {
			var details boxDetails
			if err := json.Unmarshal(output, &details); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0)
			for _, version := range details.Versions {
				for _, _provider := range version.Providers {
					if provider == "" || provider == _provider.Name {
						vals = append(vals, version.Version)
					}
				}
			}
			return carapace.ActionValues(vals...)
		})
	})
}
