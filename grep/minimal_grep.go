package grep

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (g *Grepper) MinimalGrep() {
	var allMatches []string
	for _, file := range g.Files {
		allMatches = append(allMatches, matchMinimal(file, g.Pattern)...)
	}

	g.Results = allMatches
}

func matchMinimal(filename string, pattern string) []string {
	fd, err := os.Open(filename)
	if err != nil {
		fmt.Printf("could not open %s\n", filename)
		os.Exit(1)
	}
	defer fd.Close()

	var matches []string

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() { // loops through lines until EOF
		line := scanner.Text() // get text from line

		if strings.Contains(line, pattern) {
			matches = append(matches, line)
		}
	}
	return matches
}
