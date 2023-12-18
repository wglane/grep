package grep

import (
	"bufio"
	"fmt"
	"grep/grep/common"
	"os"
	"strings"
)

func MinimalGrep(files []string, pattern string, opts *common.Options) {
	for _, file := range files {
		matchMinimal(file, pattern)
	}
}

func matchMinimal(filename string, pattern string) {
	fd, err := os.Open(filename)
	if err != nil {
		fmt.Printf("could not open %s\n", filename)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(fd)

	for scanner.Scan() { // loops through lines until EOF
		line := scanner.Text() // get text from line

		if strings.Contains(line, pattern) {
			fmt.Println(line)
		}
	}
}
