package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] != "doctor" {
		fmt.Fprintln(os.Stderr, "Kullanım: kubebpf doctor")
		os.Exit(2)
	}

	allPassed := true

	if runtime.GOOS == "linux" {
		fmt.Println("PASS: İşletim sistemi Linux")
	} else {
		fmt.Printf("FAIL: İşletim sistemi Linux değil (%s)\n", runtime.GOOS)
		allPassed = false
	}

	if _, err := os.Stat("/sys/kernel/btf/vmlinux"); err == nil {
		fmt.Println("PASS: /sys/kernel/btf/vmlinux mevcut")
	} else {
		fmt.Println("FAIL: /sys/kernel/btf/vmlinux mevcut değil")
		allPassed = false
	}

	if _, err := os.Stat("/sys/fs/cgroup/cgroup.controllers"); err == nil {
		fmt.Println("PASS: /sys/fs/cgroup/cgroup.controllers mevcut")
	} else {
		fmt.Println("FAIL: /sys/fs/cgroup/cgroup.controllers mevcut değil")
		allPassed = false
	}

	if !allPassed {
		os.Exit(1)
	}
}
