package grep

import (
	"bufio"
	"fmt"
	"grep/grep/common"
	"os"
)

func GoodGrep(files []string, pattern string, opts *common.Options) {
	for _, file := range files {
		matchGood(file, pattern, opts)
	}
}

func matchGood(filename, pattern string, opts *common.Options) {
	fd, err := os.Open(filename)
	if err != nil {
		fmt.Printf("could not open %s\n", filename)
		os.Exit(1)
	}

	matcher := common.NewMatcher(opts)
	lineNum := 1

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()

		if matcher.Match(line, pattern) {
			out := common.Result{File: filename, LineNum: lineNum, Text: line}.Format(opts)
			fmt.Println(out)
			if opts.NameOnly {
				return
			}
		}

		lineNum += 1
	}
}
