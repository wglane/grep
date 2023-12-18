package main

import (
	"flag"
	"fmt"
	"grep/grep"
	"grep/grep/common"
	"os"
)

func main() {
	opts := makeOptions()
	args := flag.Args()

	if len(args) < 2 {
		fmt.Println("usage: grep [-n -l -i -v -x] [-g good] pattern file1 file2 ...")
		os.Exit(1)
	}

	pattern := args[0]
	files := args[1:]
	if len(files) > 1 {
		opts.Name = true
	}

	grep.Grep(files, pattern, opts)
}

func makeOptions() *common.Options {
	num := flag.Bool("n", false, "prepends line number to output")
	name := flag.Bool("l", false, "only output names of files with at least one matching line")
	ins := flag.Bool("i", false, "match using case-insensitive operation")
	inv := flag.Bool("v", false, "invert the program by collecting lines that fail to match")
	ent := flag.Bool("x", false, "only search for lines where string matches entire line")
	grep := flag.String("g", "minimal", "specifies which version of grep to use ('min', 'result', 'opts', 'in')")

	flag.Parse()

	return &common.Options{
		Num:      *num,
		NameOnly: *name,
		Ins:      *ins,
		Inv:      *inv,
		Ent:      *ent,
		Grep:     *grep,
	}
}
