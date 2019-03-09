package main

import "syscall"
import "os"
import "os/exec"



func main() {
	binary, lookerr := exec.LookPath("ls")
	if lookerr != nil {
		panic(lookerr)
	}

	args := []string{"ls", "-s", "-l", "-h"}
	env := os.Environ()
	


	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
