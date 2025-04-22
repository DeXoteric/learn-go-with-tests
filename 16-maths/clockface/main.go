// Writes an SVG clockface of the current time to Stdout.
package main

import (
	"os"
	"time"

	"github.com/DeXoteric/learn-go-with-tests/16-maths/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
