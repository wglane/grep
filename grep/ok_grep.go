package grep

import (
	"bufio"
	"fmt"
	"grep/grep/common"
	"os"
	"strings"
)

func OkGrep(files []string, pattern string, opts *common.Options) {
	for _, file := range files {
		matchOk(file, pattern, opts)
	}
}

func matchOk(filename, pattern string, opts *common.Options) {
	fd, err := os.Open(filename)
	if err != nil {
		fmt.Printf("could not open %s\n", filename)
		os.Exit(1)
	}

	lineNum := 1
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()

		var isMatch bool

		// case-insensitive:
		if opts.Ins {
			lowerLine := strings.ToLower(line)
			lowerPattern := strings.ToLower(pattern)
			isMatch = (!opts.Ent && strings.Contains(lowerLine, lowerPattern)) || (opts.Ent && lowerLine == lowerPattern)
		} else {
			isMatch = (!opts.Ent && strings.Contains(line, pattern)) || (opts.Ent && line == pattern)
		}

		// inverted:
		if opts.Inv {
			isMatch = !isMatch
		}

		if isMatch {
			out := common.Result{File: filename, LineNum: lineNum, Text: line}.Format(opts)
			fmt.Println(out)
		}

		lineNum += 1
	}
}
