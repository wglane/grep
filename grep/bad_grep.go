package grep

import (
	"bufio"
	"fmt"
	"grep/grep/common"
	"os"
	"strconv"
	"strings"
)

func BadGrep(files []string, pattern string, opts *common.Options) {
	for _, file := range files {
		matchBad(file, pattern, opts)
	}
}

func matchBad(filename, pattern string, opts *common.Options) {
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
			// entire line
			if opts.NameOnly {
				fmt.Println(filename)
				return
			}
			// prepend number
			if opts.Num {
				line = strconv.Itoa(lineNum) + ":" + line
			}
			// prepend filename
			if opts.Name {
				line = filename + ":" + line
			}
			fmt.Println(line)
		}

		lineNum += 1
	}
}
