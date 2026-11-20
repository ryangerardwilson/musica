package cli

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"
)

var lookPath = exec.LookPath
var execProcess = syscall.Exec

// Command is the musica binary. Bare musica execs r2d2 musica, which opens
// the player. Every other command runs here.
func Command(args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	if len(args) > 0 {
		return Main(args, stdin, stdout, stderr)
	}
	path, err := lookPath("r2d2")
	if err != nil {
		fmt.Fprintf(stderr, "The player opens with r2d2 musica: %v\n", err)
		return 1
	}
	if err := execProcess(path, []string{"r2d2", "musica"}, os.Environ()); err != nil {
		fmt.Fprintf(stderr, "exec r2d2 musica: %v\n", err)
		return 1
	}
	return 0
}
