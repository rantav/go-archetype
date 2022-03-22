package operations

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/rantav/go-archetype/log"
	"github.com/rantav/go-archetype/template"
)

type shellOperation struct {
	sh []string
}

func newShellOperator(spec OperationSpec) *shellOperation {
	return &shellOperation{sh: spec.Sh}
}

func (o *shellOperation) Operate() error {
	for _, command := range o.sh {
		scanner := bufio.NewScanner(strings.NewReader(command))
		for scanner.Scan() {
			line := scanner.Text()
			if err := executeShell(line); err != nil {
				return err
			}
		}
	}
	return nil
}

// Template the shell commands
func (o *shellOperation) Template(vars map[string]string) error {
	var err error
	for i := range o.sh {
		o.sh[i], err = template.Execute(o.sh[i], vars)
		if err != nil {
			return err
		}
	}
	return nil
}

func executeShell(shellLine string) error {
	cmd := exec.Command("sh", "-c", shellLine)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	log.Infof("Running command: %s", shellLine)
	err := cmd.Run()
	if err != nil {
		log.Errorf("Error running command.\n\t STDOUT: %s \n\n\t STDERR: %s", stdout.String(), stderr.String())
		return fmt.Errorf("error running command %s: %w", shellLine, err)
	}
	log.Infof("Output: %s", stdout.String())
	return nil
}
