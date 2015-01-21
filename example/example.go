package main

import (
	"fmt"

	"go.scj.io/dfsr"
)

func main() {
	if backlog, err := dfsr.Backlog("oly-fs1-svr", "oly-fs2-svr", "Projects", "Projects"); err == nil {
		fmt.Printf("Projects backlog: %d\n", backlog)
	} else {
		fmt.Printf("Projects backlog error: %s\n", err)
	}
}
