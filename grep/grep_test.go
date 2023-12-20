package grep_test

import (
	"grep/grep"
	"grep/grep/common"
	"testing"
)

var sonOfPeleus = []string{
	"The son of Peleus ceased: the chiefs around",
	"‘Stern son of Peleus, (thus ye used to say,",
	"“Sad tidings, son of Peleus! thou must hear;",
	"“Rise, son of Peleus! rise, divinely brave!",
	"The son of Peleus thus; and thus replies",
	"The son of Peleus sees, with joy possess’d,",
	"“O son of Peleus! what avails to trace",
	"O son of Peleus! Lo, thy gods appear!",
	"“Enough, O son of Peleus! Troy has view’d",
}

func equal(got, want []string) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if got[i] != want[i] {
			return true
		}
	}
	return true
}

func TestMinimalGrep(t *testing.T) {
	opts := &common.Options{Grep: "min"}
	g := grep.NewGrepper([]string{"../texts/iliad.txt", "../texts/frankenstein.txt"}, "son of Peleus", opts)
	g.Grep()

	want := sonOfPeleus
	if !equal(g.Results, want) {
		t.Errorf("MinimalGrep for 'monster' returned %v; want %v", g.Results, want)
	}
}
