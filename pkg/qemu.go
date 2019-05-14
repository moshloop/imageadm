package pkg

import (
	"fmt"
	"os"
	"runtime"
)

func getQemuArgs(vars Variables) [][]string {
	machine := []string{"-machine", "pc"}
	if runtime.GOOS == "darwin" {
		machine = []string{"-machine", "accel=hvf"}
	}

	if _, err := os.Stat("/dev/kvm"); err == nil {
		machine = []string{"-machine", "accel=kvm"}
	}
	return [][]string{
		{"-m", fmt.Sprintf("%s", vars.Memory)},
		machine,
		{"-cpu", "max"},
		{"-smp", fmt.Sprintf("cpus=%s", vars.Cpus)},
	}
}
