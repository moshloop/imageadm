package pkg

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"text/template"

	log "github.com/sirupsen/logrus"
)

//Exec runs the sh script and forwards stderr/stdout to the console
func Exec(sh string, args ...interface{}) error {
	cmd := exec.Command("bash", "-c", fmt.Sprintf(sh, args...))
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%s failed with %s", sh, err)
	}

	if !cmd.ProcessState.Success() {
		return fmt.Errorf("%s failed to run", sh)
	}
	return nil
}

func Interpolate(arg string, vars interface{}) string {
	tmpl, err := template.New("test").Parse(arg)
	if err != nil {
		log.Errorf("Failed to parse template %s -> %s\n", arg, err)
		return arg
	}
	buf := bytes.NewBufferString("")

	err = tmpl.Execute(buf, vars)
	if err != nil {
		log.Errorf("Failed to execute template %s -> %s\n", arg, err)
		return arg
	}
	return buf.String()

}
func InterpolateStrings(arg []string, vars interface{}) []string {
	out := make([]string, len(arg))
	for i, e := range arg {
		out[i] = Interpolate(e, vars)
	}
	return out

}
