package action

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/cache"
	"github.com/spf13/cobra"
)

func ActionExtensions(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand(cmd.Root().Use, "--list-extensions")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		})
	})
}

type searchResult struct {
	Results []struct {
		Extensions []struct {
			ExtensionName    string
			ShortDescription string
			Publisher        struct {
				PublisherName string
			}
		}
	}
}

func ActionExtensionSearch(category string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.CallbackValue) < 2 {
			// skip search when less than two characters
			return ActionMicrosoftExtensions(category)
		}

		term := ""
		if splitted := strings.SplitN(c.CallbackValue, ".", 2); len(splitted) > 1 && len(splitted[1]) > 0 {
			term = fmt.Sprintf("publisher:%v +%v", splitted[0], splitted[1])
		} else {
			term = "publisher:" + splitted[0]
		}

		categoryFilter := ""
		if category != "" {
			categoryFilter = fmt.Sprintf(`,{"filterType":5,"value":"%v"}`, category)
		}

		args := []string{"https://marketplace.visualstudio.com/_apis/public/gallery/extensionquery",
			"-H", "content-type: application/json",
			"-H", "accept: application/json;api-version=6.1-preview.1;excludeUrls=true",
			"--data-raw", fmt.Sprintf(`{"assetTypes":["Microsoft.VisualStudio.Services.Icons.Default","Microsoft.VisualStudio.Services.Icons.Branding","Microsoft.VisualStudio.Services.Icons.Small"],"filters":[{"criteria":[{"filterType":8,"value":"Microsoft.VisualStudio.Code"},{"filterType":10,"value":"%v"},{"filterType":12,"value":"37888"}%v],"direction":2,"pageSize":100,"pageNumber":1,"sortBy":0,"sortOrder":0,"pagingToken":null}],"flags":870}`, term, categoryFilter),
			"--compressed"}

		return carapace.ActionExecCommand("curl", args...)(func(output []byte) carapace.Action {
			var result searchResult

			if err := json.Unmarshal(output, &result); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0)
			for _, results := range result.Results {
				for _, extension := range results.Extensions {
					vals = append(vals, fmt.Sprintf("%v.%v", extension.Publisher.PublisherName, extension.ExtensionName), extension.ShortDescription)
				}
			}

			if c.CallbackValue != "Microsoft" {
				return carapace.Batch(
					carapace.ActionValuesDescribed(vals...),
					ActionMicrosoftExtensions(category), // always add microsoft extensions since searching for the aliased publisher is not working
				).Invoke(c).Merge().ToA()
			}
			return carapace.ActionValuesDescribed(vals...)

		})
	})
}

func ActionMicrosoftExtensions(category string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// Microsoft extensions are aliased and only found by the 'Microsoft' publisher (though the actual name differs)
		return ActionExtensionSearch(category).Invoke(carapace.Context{CallbackValue: "Microsoft"}).ToA()
	}).Cache(24*time.Hour, cache.String(category))
}
