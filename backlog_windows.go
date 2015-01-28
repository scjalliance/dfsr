// +build windows

package dfsr

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"syscall"
)

// ErrNoReplication means that the requested rgname or rfname does not exist
var ErrNoReplication = errors.New("Replication group does not exist")

var noBacklogRE = regexp.MustCompile(`\bNo\s*Backlog\b`)
var countRE = regexp.MustCompile(`\bBacklog\s*File\s*Count:\s*(\d+)\b`)
var noReplRE = regexp.MustCompile(`\b0x80041002\b`)

// Backlog returns the backlog count
func Backlog(smem string, rmem string, rgname string, rfname string) (int, error) {
	cmd := exec.Command("dfsrdiag")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CmdLine: fmt.Sprintf(`dfsrdiag backlog /rgname:"%s" /rfname:"%s" /smem:"%s" /rmem:"%s"`, rgname, rfname, smem, rmem),
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		if noReplRE.Match(out) {
			return -1, ErrNoReplication
		}
		return -1, errors.New(string(out))
	}

	backlog := -1
	if noBacklogRE.Match(out) {
		backlog = 0
	} else if matches := countRE.FindAllSubmatch(out, 1); len(matches) > 0 && len(matches[0]) > 0 {
		var err error
		backlog, err = strconv.Atoi(string(matches[0][1]))
		if err != nil {
			fmt.Printf("%s\n", err)
			return -1, err
		}
	} else {
		return -1, errors.New(string(out))
	}
	return backlog, nil
}

// RGList returns the list of replication groups
func RGList() ([]string, error) {
	cmd := exec.Command("dfsradmin", "rg", "list", "/attr:rgname")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	var list []string
	buf := bytes.NewBuffer(out)
	scanner := bufio.NewScanner(buf)
	scanner.Scan() // shift off the unwanted first line
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		list = append(list, line)
	}
	return list, nil
}
