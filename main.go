/*
	Based on

	talk from Liz Rice at UK Docker Conf 2016
	https://www.youtube.com/watch?v=HPuvDm8IC-4

	AND

	Article by Julian Friedman
	https://www.infoq.com/articles/build-a-container-golang

	Namespacing - what it sees
		UNIX Timesharing System
		Process IDs
		File System (mount points)
		Users
		IPC
		Networking

	Control Groups - what resources can use
		CPU
		Memory
		Disk I/O
		Network
		Device permissions (/dev)
*/

package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("what?")
	}
}

func run() {
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)

	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

	fmt.Println("run()")
	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}
}

func child() {
	fmt.Printf("running %v as pid %d\n", os.Args[2:], os.Getpid())

	must(syscall.Chroot("core"))
	must(os.Chdir("/"))
	must(syscall.Mount("proc", "proc", "proc", 0, ""))

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// running command
	fmt.Println("running command")
	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
