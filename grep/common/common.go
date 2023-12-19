package common

import (
	"fmt"
	"strconv"
	"strings"
)

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
	File    string
	LineNum int
	Text    string
}

func (r Result) Format(opts *Options) string {

	// entire line
	if opts.NameOnly {
		return r.File
	}

	var builder strings.Builder
	// prepend filename
	if opts.Name {
		builder.WriteString(r.File)
		builder.WriteByte(':')
	}
	// prepend number
	if opts.Num {
		builder.WriteString(strconv.Itoa(r.LineNum))
		builder.WriteByte(':')
	}

	builder.WriteString(r.Text)
	return builder.String()
}

type Results []Result

func (r Results) Len() int {
	return len(r)
}

func (r Results) Less(i, j int) bool {
	if r[i].File == r[j].File {
		return r[i].LineNum < r[j].LineNum
	}

	return r[i].File < r[j].File
}

func (r Results) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r Results) Print(opts *Options) {
	for _, res := range r {
		fmt.Println(res.Format(opts))
	}
}

func (r Results) ToSlice(opts *Options) []string {
	var out []string
	for _, res := range r {
		out = append(out, res.Format(opts))
	}

	return out
}

type Matcher interface {
	Match(string, string) bool
}

type SubstringMatcher struct{}

func (m SubstringMatcher) Match(text, pattern string) bool {
	return strings.Contains(text, pattern)
}

type EntireLineMatcher struct{}

func (m EntireLineMatcher) Match(text, pattern string) bool {
	return text == pattern
}

type CaseInsensitiveMatcher struct {
	Matcher Matcher
}

func (m CaseInsensitiveMatcher) Match(text, pattern string) bool {
	text = strings.ToLower(text)
	pattern = strings.ToLower(pattern)

	return m.Matcher.Match(text, pattern)
}

type InvertedMatcher struct {
	Matcher Matcher
}

func (m InvertedMatcher) Match(text, pattern string) bool {
	return !m.Matcher.Match(text, pattern)
}

func NewMatcher(opts *Options) Matcher {
	var m Matcher

	if opts.Ent {
		m = EntireLineMatcher{}
	} else {
		m = SubstringMatcher{}
	}

	if opts.Ins {
		m = CaseInsensitiveMatcher{Matcher: m}
	}

	if opts.Inv {
		m = InvertedMatcher{Matcher: m}
	}

	return m
}
