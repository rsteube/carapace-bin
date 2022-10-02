package action

import (
	"encoding/json"
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

type snapshot struct {
	Hostname string
	Short_id string
	Tags     []string `json:"tags,omitempty"`
	Time     time.Time
}

type snapshotDetails struct {
	Hostname string
	Paths    []string
	Tags     []string `json:"tags,omitempty"`
	Time     time.Time
}

func actionSnapshots(cmd *cobra.Command, f func(snapshots []snapshot) carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"snapshots", "--json"}
		if f := cmd.Flag("repo"); f.Changed {
			args = append(args, "--repo", f.Value.String())
		}
		return carapace.ActionExecCommand("restic", args...)(func(output []byte) carapace.Action {
			var snapshots []snapshot
			if err := json.Unmarshal(output, &snapshots); err != nil {
				return carapace.ActionMessage(err.Error())
			}
			return f(snapshots)
		})
	})
}

func ActionSnapshotTags(cmd *cobra.Command) carapace.Action {
	return actionSnapshots(cmd, func(snapshots []snapshot) carapace.Action {
		tags := make(map[string]bool)
		for _, s := range snapshots {
			for _, t := range s.Tags {
				tags[t] = true
			}
		}

		vals := make([]string, 0)
		for tag := range tags {
			vals = append(vals, tag)
		}
		return carapace.ActionValues(vals...)
	})
}

func ActionSnapshotHosts(cmd *cobra.Command) carapace.Action {
	return actionSnapshots(cmd, func(snapshots []snapshot) carapace.Action {
		vals := make([]string, 0)
		for _, s := range snapshots {
			vals = append(vals, s.Hostname)
		}
		return carapace.ActionValues(vals...)
	})
}

func ActionSnapshotIDs(cmd *cobra.Command) carapace.Action {
	return actionSnapshots(cmd, func(snapshots []snapshot) carapace.Action {
		vals := make([]string, 0)
		for _, s := range snapshots {
			vals = append(vals, s.Short_id, util.FuzzyAgo(time.Since(s.Time)))
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

func ActionSnapshotPaths(cmd *cobra.Command, id string) carapace.Action {
	// TODO repo/path/alias flags
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("restic", "cat", "snapshot", id)(func(output []byte) carapace.Action {
			var details snapshotDetails
			if err := json.Unmarshal(output, &details); err != nil {
				return carapace.ActionMessage(err.Error())
			}
			return carapace.ActionValues(details.Paths...)
		})
	})
}
