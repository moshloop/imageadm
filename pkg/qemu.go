package pkg

import (
	"fmt"
)

func getQemuArgs(vars Variables) [][]string {
	return [][]string{
		{"-m", fmt.Sprintf("%s", vars.Memory)},
		{"-machine", "accel=hvf"},
		{"-cpu", "max"},
		{"-smp", fmt.Sprintf("cpus=%s", vars.Cpus)},
	}
}
