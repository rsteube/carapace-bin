package action

import (
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/cache"
	"github.com/spf13/cobra"
)

type resource struct {
	Urn string
	Id  string
}

func ActionUrns(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			cwd := cmd.Flag("cwd").Value.String()
			stack := cmd.Flag("stack").Value.String()

			c.Setenv("PULUMI_SKIP_UPDATE_CHECK", "1")
			return carapace.ActionExecCommand("pulumi", "--cwd", cwd, "stack", "--stack", stack, "--show-urns")(func(output []byte) carapace.Action {
				reUrn := regexp.MustCompile(`URN: (?P<urn>urn:pulumi.*)`)
				reId := regexp.MustCompile(`ID: (?P<id>.*)`)

				resources := make([]resource, 0)
				for _, line := range strings.Split(string(output), "\n") {
					if reUrn.MatchString(line) {
						resources = append(resources, resource{Urn: reUrn.FindStringSubmatch(line)[1]})
					}
					if reId.MatchString(line) {
						resources[len(resources)-1].Id = reId.FindStringSubmatch(line)[1]
					}
				}

				vals := make([]string, 0)
				for _, resource := range resources {
					vals = append(vals, resource.Urn, resource.Id)
				}
				return carapace.ActionValuesDescribed(vals...)
			}).Invoke(c).ToA()
		}).Cache(5*time.Second, func() (string, error) {
			workdir := c.Dir
			if cmd.Flag("cwd").Changed { // TODO use preinvoke?
				workdir = cmd.Flag("cwd").Value.String()
			}
			stack := cmd.Flag("stack").Value.String()

			absWd, err := filepath.Abs(workdir)
			if err != nil {
				return "", err
			}
			return cache.String(absWd, stack)()
		}).Invoke(c).ToMultiPartsA("::")
	})
}
