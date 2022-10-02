package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var container_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update configuration of one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_updateCmd).Standalone()

	container_updateCmd.Flags().String("blkio-weight", "", "Block IO (relative weight), between 10 and 1000, or 0 to")
	container_updateCmd.Flags().String("cpu-period", "", "Limit CPU CFS (Completely Fair Scheduler) period")
	container_updateCmd.Flags().String("cpu-quota", "", "Limit CPU CFS (Completely Fair Scheduler) quota")
	container_updateCmd.Flags().String("cpu-rt-period", "", "Limit the CPU real-time period in microseconds")
	container_updateCmd.Flags().String("cpu-rt-runtime", "", "Limit the CPU real-time runtime in microseconds")
	container_updateCmd.Flags().StringP("cpu-shares", "c", "", "CPU shares (relative weight)")
	container_updateCmd.Flags().String("cpus", "", "Number of CPUs")
	container_updateCmd.Flags().String("cpuset-cpus", "", "CPUs in which to allow execution (0-3, 0,1)")
	container_updateCmd.Flags().String("cpuset-mems", "", "MEMs in which to allow execution (0-3, 0,1)")
	container_updateCmd.Flags().String("kernel-memory", "", "Kernel memory limit")
	container_updateCmd.Flags().StringP("memory", "m", "", "Memory limit")
	container_updateCmd.Flags().String("memory-reservation", "", "Memory soft limit")
	container_updateCmd.Flags().String("memory-swap", "", "Swap limit equal to memory plus swap: '-1' to enable")
	container_updateCmd.Flags().String("pids-limit", "", "Tune container pids limit (set -1 for unlimited)")
	container_updateCmd.Flags().String("restart", "", "Restart policy to apply when a container exits")
	containerCmd.AddCommand(container_updateCmd)

	rootAlias(container_updateCmd, func(cmd *cobra.Command, isAlias bool) {
		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			"restart": carapace.ActionValues("always", "no", "on-failure", "unless-stopped").StyleF(style.ForKeyword),
		})

		carapace.Gen(cmd).PositionalAnyCompletion(docker.ActionContainers())
	})
}
