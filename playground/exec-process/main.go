// spawn - need an external process accessible to a running Go process
// exec - replace current Go process with another (perhaps non-Go) one

package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// Go requires an absolute path to the binary we want to execute
	bin, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// exec requires arguments in slice form
	// first argument should be the program name
	args := []string{"ls", "-a", "-l", "-h"}

	// exec needs a set of environment variables to use
	// we provide our current environment
	env := os.Environ()

	// actual `syscall.Exec` call.
	// if successful, execution of Go process ends here and is replaced by `/bin/ls -a -l -h`
	// if there is an error, we get a return value
	execErr := syscall.Exec(bin, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
