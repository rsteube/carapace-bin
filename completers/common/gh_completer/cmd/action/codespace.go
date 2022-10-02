package action

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

type codespace struct {
	Name       string
	Owner      string
	Repository string
	State      string
}

func ActionCodespaces() carapace.Action {
	return carapace.ActionExecCommand("gh", "codespace", "list", "--json", "name,owner,repository,state")(func(output []byte) carapace.Action {
		var codespaces []codespace
		if err := json.Unmarshal(output, &codespaces); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		vals := make([]string, 0)
		for _, c := range codespaces {
			vals = append(vals, c.Name, fmt.Sprintf("%v - %v - %v", c.Owner, c.Repository, c.State))
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

func ActionCodespacePath(codespace string, expand bool) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		path := filepath.Dir(c.CallbackValue)

		if !expand {
			// scp treats each filename argument as a shell expression,
			// subjecting it to expansion of environment variables, braces,
			// tilde, backticks, globs and so on. Because these present a
			// security risk (see https://lwn.net/Articles/835962/), we
			// disable them by shell-escaping the argument unless the user
			// provided the -e flag.
			path = `'` + strings.Replace(path, `'`, `'\''`, -1) + `'`
		}

		args := []string{"codespace", "ssh", "--codespace", codespace, "--", "ls", "-1", "-p", path}
		if splitted := strings.Split(c.CallbackValue, "/"); strings.HasPrefix(splitted[len(splitted)-1], ".") {
			args = append(args, "-a") // show hidden
		}
		return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
			return carapace.ActionExecCommand("gh", args...)(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\n")
				return carapace.ActionValues(lines[:len(lines)-1]...)
			})
		}).StyleF(style.ForPathExt)
	})
}

func ActionCodespaceMachines() carapace.Action {
	return carapace.ActionValues("TODO")
}

type codespacePort struct {
	Label      string
	SourcePort int
	Visibility string
}

func ActionCodespacePorts(codespace string) carapace.Action {
	return carapace.ActionExecCommand("gh", "codespace", "ports", "--codespace", codespace, "--json", "label,sourcePort,visibility")(func(output []byte) carapace.Action {
		var ports []codespacePort
		if err := json.Unmarshal(output, &ports); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, port := range ports {
			vals = append(vals, strconv.Itoa(port.SourcePort), fmt.Sprintf("%v - %v", port.Label, port.Visibility))
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

func ActionCodespaceFields() carapace.Action {
	return carapace.ActionValues(
		"name",
		"owner",
		"repository",
		"state",
		"gitStatus",
		"createdAt",
		"lastUsedAt",
	)
}

func ActionCodespacePortFields() carapace.Action {
	return carapace.ActionValues(
		"sourcePort",
		// "destinationPort", // TODO(mislav): this appears to always be blank?
		"visibility",
		"label",
		"browseUrl",
	)
}
