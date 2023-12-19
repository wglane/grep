package grep

import (
	"bufio"
	"fmt"
	"grep/grep/common"
	"os"
	"strconv"
	"strings"
)

func (g *Grepper) BadGrep() {
	var allMatches []string
	for _, file := range g.Files {
		allMatches = append(allMatches, matchBad(file, g.Pattern, g.Opts)...)
	}

	g.Results = allMatches
}

func matchBad(filename, pattern string, opts *common.Options) []string {
	fd, err := os.Open(filename)
	if err != nil {
		fmt.Printf("could not open %s\n", filename)
		os.Exit(1)
	}
	defer fd.Close()

	var matches []string

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
			// entire line
			if opts.NameOnly {
				matches = append(matches, filename)
				return matches
			}
			// prepend number
			if opts.Num {
				line = strconv.Itoa(lineNum) + ":" + line
			}
			// prepend filename
			if opts.Name {
				line = filename + ":" + line
			}
			matches = append(matches, line)
		}

		lineNum += 1
	}

	return matches
}
