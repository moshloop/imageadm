package pkg

import (
	"fmt"
	"os"
	"runtime"
)

func getQemuArgs(vars Variables) [][]string {
	args := [][]string{
		{"-m", fmt.Sprintf("%s", vars.Memory)},
		{"-smp", fmt.Sprintf("cpus=%s", vars.Cpus)},
	}

	if runtime.GOOS == "darwin" {
		args = append(args, []string{"-machine", "accel=hvf"})
	}

	if _, err := os.Stat("/dev/kvm"); err == nil {
		args = append(args, []string{"-machine", "accel=kvm"})
	}
	return args
}
