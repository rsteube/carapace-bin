package action

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

type TestOpts struct {
	Benchmark bool
	Example   bool
	Test      bool
}

func ActionTests(packages []string, opts TestOpts) carapace.Action {
	return carapace.ActionExecCommand("go", append([]string{"test", "-list=.*"}, packages...)...)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if opts.Benchmark && strings.HasPrefix(line, "Benchmark") {
				vals = append(vals, line, style.Blue)
			} else if opts.Example && strings.HasPrefix(line, "Example") {
				vals = append(vals, line, style.Green)
			} else if opts.Test && strings.HasPrefix(line, "Test") {
				vals = append(vals, line, style.Default)
			}
		}
		return carapace.ActionStyledValues(vals...)
	})
}
