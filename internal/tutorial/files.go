package tutorial

import (
	_ "embed"
	"log"
)

// This variable uses a Golang Directive '//go:embed' to tell to the builder to embed a file with
// a relative given path.
// During the build phase, the file content is read and stored in a byte array.
// This byte array can then be used by the code at runtime.
var (
	//go:embed static/words.txt
	data []byte
)

func Files() {
	// get static local file content by using the 'data' var from the embedded content above.
	log.Printf(string(data))
}
