package grep

import (
	"bufio"
	"fmt"
	"grep/grep/common"
	"os"
	"strings"
)

func (g *Grepper) OkGrep() {
	var allRes common.Results
	for _, file := range g.Files {
		allRes = append(allRes, matchOk(file, g.Pattern, g.Opts)...)
	}

	g.Results = allRes.ToSlice(g.Opts)
}

func matchOk(filename, pattern string, opts *common.Options) common.Results {
	fd, err := os.Open(filename)
	if err != nil {
		fmt.Printf("could not open %s\n", filename)
		os.Exit(1)
	}
	defer fd.Close()

	var res common.Results

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
			res = append(res, common.Result{File: filename, LineNum: lineNum, Text: line})
			if opts.NameOnly {
				return res
			}
		}
		lineNum += 1
	}

	return res
}
