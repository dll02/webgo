package main

import (
	"fmt"
	"syscall"
)

func main10() {
	pid, _, _ := syscall.Syscall(39, 0, 0, 0) // 用不到的就补上 0
	fmt.Println("Process id: ", pid)
}
