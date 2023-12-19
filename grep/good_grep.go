package grep

import (
	"bufio"
	"fmt"
	"grep/grep/common"
	"os"
)

func (g *Grepper) GoodGrep() {
	var allRes common.Results
	for _, file := range g.Files {
		allRes = append(allRes, matchGood(file, g.Pattern, g.Opts)...)
	}

	g.Results = allRes.ToSlice(g.Opts)
}

func matchGood(filename, pattern string, opts *common.Options) common.Results {
	fd, err := os.Open(filename)
	if err != nil {
		fmt.Printf("could not open %s\n", filename)
		os.Exit(1)
	}
	defer fd.Close()

	var res common.Results

	matcher := common.NewMatcher(opts)
	lineNum := 1

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()

		if matcher.Match(line, pattern) {
			res = append(res, common.Result{File: filename, LineNum: lineNum, Text: line})
			if opts.NameOnly {
				return res
			}
		}

		lineNum += 1
	}

	return res
}
