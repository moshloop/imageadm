package pkg

import (
	"fmt"
	"runtime"
)

func getQemuArgs(vars Variables) [][]string {
	accel := "kvm"
	if runtime.GOOS == "darwin" {
		accel = "hvf"
	}
	return [][]string{
		{"-m", fmt.Sprintf("%s", vars.Memory)},
		{"-machine", "accel=" + accel},
		{"-cpu", "max"},
		{"-smp", fmt.Sprintf("cpus=%s", vars.Cpus)},
	}
}
