package action

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

const (
	GH_CONFIG_DIR   = "GH_CONFIG_DIR"
	XDG_CONFIG_HOME = "XDG_CONFIG_HOME"
	XDG_STATE_HOME  = "XDG_STATE_HOME"
	XDG_DATA_HOME   = "XDG_DATA_HOME"
	APP_DATA        = "AppData"
	LOCAL_APP_DATA  = "LocalAppData"
)

// Data path precedence
// 1. XDG_DATA_HOME
// 2. LocalAppData (windows only)
// 3. HOME
func dataDir() string {
	var path string
	if a := os.Getenv(XDG_DATA_HOME); a != "" {
		path = filepath.Join(a, "gh")
	} else if b := os.Getenv(LOCAL_APP_DATA); runtime.GOOS == "windows" && b != "" {
		path = filepath.Join(b, "GitHub CLI")
	} else {
		c, _ := os.UserHomeDir()
		path = filepath.Join(c, ".local", "share", "gh")
	}

	return path
}

func Extensions() ([]string, error) {
	files, err := os.ReadDir(dataDir() + "/extensions")
	if err != nil {
		return nil, err
	}

	vals := make([]string, 0)
	for _, file := range files {
		if file.IsDir() {
			vals = append(vals, strings.TrimPrefix(file.Name(), "gh-"))
		}
	}
	return vals, nil
}

func ActionExtensions() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		extensions, err := Extensions()
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return carapace.ActionValues(extensions...)
	})
}

func ActionTopExtensions() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult struct {
			Items []struct {
				FullName        string `json:"full_name"`
				Description     string
				StargazersCount int `json:"stargazers_count"`
			}
		}

		return ApiV3Action(&cobra.Command{}, fmt.Sprintf(`search/repositories?per_page=100&q=%v`, url.QueryEscape("topic:gh-extension sort:stars-desc")), &queryResult, func() carapace.Action {
			vals := make([]string, 0)
			for _, item := range queryResult.Items {
				vals = append(vals, item.FullName, fmt.Sprintf("%v - %v", item.StargazersCount, item.Description))
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	}).Cache(24 * time.Hour)
}
