package action

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type SecretOpts struct {
	App string
	Org string
	Env string
}

func (o SecretOpts) format() []string {
	args := make([]string, 0)
	if o.App != "" {
		args = append(args, "--app", o.App)
	}
	if o.Org != "" {
		args = append(args, "--org", o.Org)
	}
	if o.Env != "" {
		args = append(args, "--env", o.Env)
	}
	return args
}

func ActionSecrets(cmd *cobra.Command, opts SecretOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		repo := ""
		if cmd.Flag("repo") != nil {
			repo = cmd.Flag("repo").Value.String()
		}

		return carapace.ActionExecCommand(ghExecutable(), append([]string{"secret", "--repo", repo, "list"}, opts.format()...)...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0, len(lines)-1)
			for _, line := range lines[:len(lines)-1] {
				vals = append(vals, strings.SplitN(line, "\t", 2)...)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
