// +build windows

package dfsr

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"syscall"
)

// Backlog returns the backlog count
func Backlog(smem string, rmem string, rgname string, rfname string) (int, error) {
	cmd := exec.Command("dfsrdiag")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CmdLine: fmt.Sprintf(`dfsrdiag backlog /rgname:"%s" /rfname:"%s" /smem:"%s" /rmem:"%s"`, rgname, rfname, smem, rmem),
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		return -1, err
	}

	backlog := -1
	noBacklogRE := regexp.MustCompile(`\bNo\s*Backlog\b`)
	countRE := regexp.MustCompile(`\bBacklog\s*File\s*Count:\s*(\d+)\b`)
	if noBacklogRE.Match(out) {
		backlog = 0
	} else if matches := countRE.FindAllSubmatch(out, 1); len(matches) > 0 && len(matches[0]) > 0 {
		var err error
		backlog, err = strconv.Atoi(string(matches[0][0]))
		if err != nil {
			fmt.Printf("%s\n", err)
			return -1, err
		}
	} else {
		return -1, errors.New(string(out))
	}
	return backlog, nil
}
