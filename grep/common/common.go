package common

type Options struct {
	Num      bool
	Name     bool
	NameOnly bool
	Ins      bool
	Inv      bool
	Ent      bool
	Grep     string
}

type Result struct {
	file    string
	lineNum int
	text    string
}
