package grep

import (
	"fmt"
	"grep/grep/common"
)

type Grepper struct {
	Files   []string
	Pattern string
	Opts    *common.Options
	Results []string
}

func NewGrepper(files []string, pattern string, opts *common.Options) *Grepper {
	return &Grepper{
		Files:   files,
		Pattern: pattern,
		Opts:    opts,
	}
}

func (g *Grepper) Grep() error {
	switch g.Opts.Grep {
	case "min", "minimal":
		g.MinimalGrep()
	case "bad":
		g.BadGrep()
	case "ok":
		g.OkGrep()
	case "good":
		g.GoodGrep()
	case "great":
		g.GreatGrep()
	default:
		return fmt.Errorf("grep flag (-g) must be one of 'great', 'good', 'ok', 'bad', or 'min', got: %s", g.Opts.Grep)
	}

	return nil
}
