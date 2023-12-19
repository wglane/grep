package grep

import (
	"bufio"
	"fmt"
	"grep/grep/common"
	"os"
	"sort"
	"sync"
)

func GreatGrep(files []string, pattern string, opts *common.Options) {
	ch := make(chan common.Result)
	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go matchGreat(file, pattern, opts, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var res common.Results
	for r := range ch {
		res = append(res, r)
	}

	sort.Sort(res)
	res.Print(opts)
}

func matchGreat(filename string, pattern string, opts *common.Options, wg *sync.WaitGroup, ch chan common.Result) {
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
			ch <- common.Result{File: filename, LineNum: lineNum, Text: line}

			if opts.NameOnly {
				return
			}
		}

		lineNum += 1
	}
}
