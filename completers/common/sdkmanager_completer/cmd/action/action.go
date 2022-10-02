package action

import (
	"regexp"
	"strings"
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/cache"
	"github.com/spf13/cobra"
)

func ActionAvailablePackages(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"--list"}
		sdk_root := cmd.Flag("sdk_root").Value.String()
		if sdk_root != "" {
			args = append(args, "--sdk_root="+sdk_root)
		}
		channel := cmd.Flag("channel").Value.String()
		if channel != "" {
			args = append(args, "--channel="+channel)
		}

		return carapace.ActionExecCommand("sdkmanager", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			r := regexp.MustCompile(`^ +(?P<path>[^ ]+) +[|] (?P<version>[^ ]+) +[|] (?P<description>.*)$`)

			vals := make([]string, 0)
			for _, line := range lines {
				if r.MatchString(line) {
					matches := r.FindStringSubmatch(line)
					if matches[1] != "Path" && matches[1] != "-------" {
						vals = append(vals, matches[1], matches[3])
					}
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		}).Cache(1*time.Hour, cache.String(sdk_root, channel))
	})
}

func ActionInstalledPackages(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"--list_installed"}
		if sdk_root := cmd.Flag("sdk_root").Value.String(); sdk_root != "" {
			args = append(args, "--sdk_root="+sdk_root)
		}

		return carapace.ActionExecCommand("sdkmanager", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for _, line := range lines {
				if splitted := strings.Split(line, "|"); len(splitted) == 4 {
					path := strings.TrimSpace(splitted[0])
					if path != "Path" && path != "-------" {
						vals = append(vals, path, strings.TrimSpace(splitted[2]))
					}
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
