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
	sh []shellCmdOperation
}

type shellCmdOperation struct {
	cmd       string
	multiline bool
}

func newShellOperator(spec OperationSpec) *shellOperation {
	operations := make([]shellCmdOperation, len(spec.Sh))
	for i := range spec.Sh {
		operations[i] = shellCmdOperation{
			cmd:       spec.Sh[i].Cmd,
			multiline: spec.Sh[i].Multiline,
		}
	}
	return &shellOperation{sh: operations}
}

func (o *shellOperation) Operate() error {
	for _, command := range o.sh {
		if err := invokeCmdOperation(command); err != nil {
			return err
		}
	}
	return nil
}

func invokeCmdOperation(command shellCmdOperation) error {
	switch command.multiline {
	case true:
		return invokeMultilineCmd(command.cmd)
	default:
		return splitLinesAndExecute(command.cmd)
	}
}

func invokeMultilineCmd(cmd string) error {
	return executeShell(strings.TrimSpace(cmd))
}

func splitLinesAndExecute(cmd string) error {
	scanner := bufio.NewScanner(strings.NewReader(cmd))
	for scanner.Scan() {
		line := scanner.Text()
		if err := executeShell(line); err != nil {
			return err
		}
	}
	return nil
}

// Template the shell commands
func (o *shellOperation) Template(vars map[string]string) error {
	var err error
	for i := range o.sh {
		o.sh[i].cmd, err = template.Execute(o.sh[i].cmd, vars)
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
