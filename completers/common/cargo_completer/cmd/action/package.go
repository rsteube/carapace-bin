package action

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionPackageSearch() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// TODO crates.io has no prefix search so the results aren't very helpful at the moment
		return carapace.ActionExecCommand("cargo", "search", "--limit", "100", c.CallbackValue)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			r := regexp.MustCompile(`^(?P<name>[^ ]+) = "(?P<version>[^"]+)" +#( (?P<description>.*))?$`)

			vals := make([]string, 0)
			for _, line := range lines {
				if r.MatchString(line) {
					matches := r.FindStringSubmatch(line)
					if len(matches) > 3 {
						vals = append(vals, matches[1], matches[3])
					} else {
						vals = append(vals, matches[1], "")
					}
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

type content struct {
	Name string
}

func ActionGithubPackageSearch() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		batch := carapace.Batch()

		if len(c.Args) < 1 {
			batch = append(batch, actionGithubPackageIndex("1"))
		}

		if len(c.Args) < 2 {
			batch = append(batch, actionGithubPackageIndex("2"))
			batch = append(batch, carapace.ActionCallback(func(c carapace.Context) carapace.Action {
				return actionGithubPackageIndex("").Invoke(c).Filter([]string{"1", "2", "3", "config.json"}).ToA()
			}))
		}

		if len(c.Args) < 3 {
			batch = append(batch, actionGithubPackageIndex("3"))
		}

		if len(c.CallbackValue) > 3 {
			batch = append(batch, actionGithubPackageIndex(fmt.Sprintf("%v/%v", c.CallbackValue[0:2], c.CallbackValue[2:4])))
		} else if len(c.CallbackValue) > 1 {
			batch = append(batch, carapace.ActionCallback(func(c carapace.Context) carapace.Action {
				return actionGithubPackageIndex(c.CallbackValue[0:2]).Invoke(c).Prefix(c.CallbackValue[0:2]).ToA()
			}))
		}

		return batch.ToA().NoSpace()
	})
}

func actionGithubPackageIndex(path string) carapace.Action {
	return carapace.ActionExecCommand("gh", "api", "repos/rust-lang/crates.io-index/contents/"+url.PathEscape(path))(func(output []byte) carapace.Action {
		var contents []content
		if err := json.Unmarshal(output, &contents); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0, len(contents))
		for _, c := range contents {
			vals = append(vals, c.Name)
		}
		return carapace.ActionValues(vals...)
	})
}

type version struct {
	Features map[string][]string `json:"features,omitempty"`
	Vers     string
}

func githubPath(pkg string) string {
	path := ""
	switch len(pkg) {
	case 0:
		return ""
	case 1:
		path = "1/" + pkg
	case 2:
		path = "2/" + pkg
	case 3:
		path = "3/" + pkg
	default:
		path = fmt.Sprintf("%v/%v/%v", pkg[0:2], pkg[2:4], pkg)
	}
	return path
}

func ActionGithubPackageVersions(pkg string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("gh", "api", "repos/rust-lang/crates.io-index/contents/"+url.PathEscape(githubPath(pkg)), "--jq", ".content")(func(output []byte) carapace.Action {
			decoded, err := base64.StdEncoding.DecodeString(string(output))
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}

			lines := strings.Split(string(decoded), "\n")

			vals := make([]string, 0)
			var v version
			for _, line := range lines[:len(lines)-1] {
				if err := json.Unmarshal([]byte(line), &v); err != nil {
					return carapace.ActionMessage(err.Error())
				}
				vals = append(vals, v.Vers)
			}
			return carapace.ActionValues(vals...)
		})
	})
}

func ActionGithubPackageFeatures(pkg string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("gh", "api", "repos/rust-lang/crates.io-index/contents/"+url.PathEscape(githubPath(pkg)), "--jq", ".content")(func(output []byte) carapace.Action {
			decoded, err := base64.StdEncoding.DecodeString(string(output))
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}

			lines := strings.Split(string(decoded), "\n")

			vals := make([]string, 0)
			var v version
			for _, line := range lines[:len(lines)-1] {
				if err := json.Unmarshal([]byte(line), &v); err != nil {
					return carapace.ActionMessage(err.Error())
				}
				for feature := range v.Features {
					vals = append(vals, feature)
				}
			}
			return carapace.ActionValues(vals...)
		})
	})
}
