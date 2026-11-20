// Package player opens the musica player. r2d2 runs it as r2d2 musica.
package player

import (
	"io"

	"github.com/ryangerardwilson/musica/internal/cli"
)

// Run runs the musica player command with args. No arguments open the saved
// library in this process. end, play, open, config, help, and
// download-worker behave as they do on the musica binary.
func Run(args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	return cli.Main(args, stdin, stdout, stderr)
}
