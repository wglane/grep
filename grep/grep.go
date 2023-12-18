package grep

import (
	"fmt"
	"grep/grep/common"
	"os"
)

func Grep(files []string, pattern string, opts *common.Options) {
	switch opts.Grep {
	case "min", "minimal":
		MinimalGrep(files, pattern, opts)
	case "bad":
		BadGrep(files, pattern, opts)
	default:
		fmt.Printf("grep must be one of 'good', 'ok', 'bad', or 'minimal', got %s\n", opts.Grep)
		os.Exit(1)
	}
}