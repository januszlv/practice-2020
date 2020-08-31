package smilespars

import (
	"log"
	"regexp"
	"github.com/januszlv/practice-2020/molgraph"
)

func init() {
	aliphaticOrganicPattern = regexp.MustCompile(aliphaticOrganicString)
	aromaticOrganicPattern = regexp.MustCompile(aromaticOrganicString)
	aromaticSymbolPattern = regexp.MustCompile(aromaticSymbolString)
	elementsSymbolPattern = regexp.MustCompile(elementsSymbolString)
	chiralPattern = regexp.MustCompile(chiralString)
	numberPattern = regexp.MustCompile(numberString)
	valence = map[string][]int{
		`B`: {3},
		`C`: {4},
		`N`: {3, 5},
		`O`: {2},
		`P`: {3, 5},
		`S`: {2, 4, 6},
		`F`: {1},
		`Cl`: {1},
		`Br`: {1},
		`I`: {1},
	}
}

// returns slice of molecular graphs
func ReadGraphsFromSMILES(text string) []molgraph.MolecularGraph {
	l := newLexer(text)
	tokens := l.run()

	p := newParser(tokens, text)
	graphs, ok := p.parse()
	if ok {
		log.Println("parsing done successfully!")
	} else {
		log.Fatalf("some errors appeared during parsing!")
	}

	for i := 0; i < len(graphs); i++ {
		graphs[i].AddExplicitHydrogens()
	}
	return graphs
}
