// our Go program may need to spawn other non-Go processes

package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {
	// simple command that takes no arguments or input and just prints something to stdout
	dateCmd := exec.Command("date")

	// `Output` method runs the command, waits for it to finish and collects
	// its standard output. it holds bytes with date info
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// `Output` and other methods of Command will return `*exec.Error` if there
	// was a problem executing the command and `*exec.ExitError` if the command
	// ran but exited with a nonzero return code.
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed executing:", err)
		case *exec.ExitError:
			fmt.Println("command exit return code =", e.ExitCode())
		default:
			panic(err)
		}
	}

	// pipe data to the external process on its stdin and collect the results
	// from its stdout
	grepCmd := exec.Command("grep", "hello")

	// explicitly grab input/output pipes
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()

	grepCmd.Start()

	// write input
	grepIn.Write([]byte("hello grep\ngoodbye grep\nhello odi"))
	grepIn.Close()

	// read resulting output
	grepBytes, _ := io.ReadAll(grepOut)

	// wait for process to exit
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// when spawning commands we need to provide an explicitly delineated command
	// and argument array
	// we can use bash -c option if we want to spwan full command with string
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
