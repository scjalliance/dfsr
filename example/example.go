package main

import (
	"fmt"
	"os"

	"go.scj.io/dfsr"
)

func main() {
	list, err := dfsr.RGList()
	if err != nil {
		panic(err)
	}
	for _, item := range list {
		fmt.Printf("[%s]\n", item)
	}
	os.Exit(0)
	if backlog, err := dfsr.Backlog("oly-fs1-svr", "oly-fs2-svr", "Projects", "Projects"); err == nil {
		fmt.Printf("Projects backlog: %d\n", backlog)
	} else {
		fmt.Printf("Projects backlog error: %s\n", err)
	}
}
