package smilespars

import (
	"github.com/januszlv/practice-2020/molgraph"
	"log"
	"strconv"
	"strings"

)

type parser struct {
	tokens []token
	currentTokenIndex int
	currentGraph *molgraph.MolecularGraph
	currentAtom molgraph.Atom
	currentAtomIndex int
	atomToLinkIndex int
	branchedAtomIndexes []int
	bondOrder float32
	ringClosures []*ringClosureInfo
}

type ringClosureInfo struct {
	closureAtomIndex int
	bond float32
}

func (p parser) currentToken() token {
	if p.currentTokenIndex >= len(p.tokens) {
		log.Fatalf("out of tokens")
	}
	return p.tokens[p.currentTokenIndex]
}

func (p *parser) nextToken() {
	if p.currentTokenIndex < len(p.tokens)-1 {
		p.currentTokenIndex++
	}
}

func newParser(tokens []token) *parser {
	p := new(parser)
	p.tokens = make([]token, len(tokens))
	copy(p.tokens, tokens)
	p.ringClosures = make([]*ringClosureInfo, 100)
	return p
}

func (p *parser) parse() ([]molgraph.MolecularGraph, bool) {
	isInputCorrect := true
	var graphs []molgraph.MolecularGraph
	for i:= 1; p.currentToken().tType != endOfFile; {
		if graph, ok:= p.parseSMILES(); ok {
			graphs = append(graphs, *graph)
			log.Println("SMILES write №" + strconv.Itoa(i) + " parsed successfully!")

		} else {
			isInputCorrect = false
			log.Fatalf("SMILES write №" + strconv.Itoa(i) + " failed to be parsed!")
		}
		i++
	}
	return graphs, isInputCorrect
}

func (p *parser) parseSMILES() (*molgraph.MolecularGraph, bool) {
	p.currentGraph = new(molgraph.MolecularGraph)
	p.currentAtomIndex = 0
	p.atomToLinkIndex = 0
	p.parseChain()
	for i := 0; i < len(p.ringClosures); i++ {
		if p.ringClosures[i] != nil {
			log.Fatalf(p.currentToken().start.String() + ": unmatched ring indices at the end of SMILES")
		}
	}
	if p.currentToken().isTerminator() {
		p.fillValence(false)
		p.nextToken()
		return p.currentGraph, true
	}
	return nil, false
}


func (p *parser) parseChain() bool {
	startTokenIndex := p.currentTokenIndex
	if p.parseBranchedAtom() {
		if p.currentToken().isBond() || p.currentToken().tType == dot {
			p.bondOrder = bondToOrder(p.currentToken())
			p.nextToken()
		} else {
			p.bondOrder = -1
		}
		p.parseChain()
		return true
	}

	p.currentTokenIndex = startTokenIndex
	return false
}


func (p *parser) parseBranchedAtom() bool {
	if p.parseAtom() {
		for p.parseRingBond() {}
		for p.parseBranch() {}
		return true
	}
	return false
}

func (p *parser) parseAtom() (bool) {
	p.currentAtom = molgraph.DefaultAtom()
	if p.parseBracketAtom() {
		p.addAtom()
		return true
	}

	if p.currentToken().tType == aliphaticOrganic ||
		p.currentToken().tType == aromaticOrganic ||
		p.currentToken().tType == unknown {
			if p.currentToken().tType == aromaticOrganic {
				p.currentAtom.IsAromatic = true
			}
			p.currentAtom.Element = strings.Title(p.currentToken().attribute)
			p.addAtom()
			p.nextToken()
			return true
	}

	return false
}

func (p *parser) parseRingBond() bool {
	startTokenIndex := p.currentTokenIndex

	info := &ringClosureInfo{
		closureAtomIndex: p.currentAtomIndex-1,
		bond: -1,
	}

	if p.currentToken().isBond() {
		info.bond = bondToOrder(p.currentToken())
		p.nextToken()
	}

	if p.currentToken().tType == digit {
		closureIndex, _ := strconv.Atoi(p.currentToken().attribute)
		p.handleClosure(closureIndex, info)
		p.nextToken()
		return true
	}

	if p.currentToken().tType == percent {
		p.nextToken()
		if p.currentToken().tType == digit {
			indexString := p.currentToken().attribute
			p.nextToken()
			if p.currentToken().tType == digit {
				indexString += p.currentToken().attribute
				closureIndex, _ := strconv.Atoi(indexString)
				p.handleClosure(closureIndex, info)
				p.nextToken()
				return true
			}
		}
	}

	p.currentTokenIndex = startTokenIndex
	return false
}

func (p *parser) parseBranch() bool {
	startTokenIndex := p.currentTokenIndex

	if p.currentToken().tType == lRoundParen {
		p.branchedAtomIndexes = append(p.branchedAtomIndexes, p.atomToLinkIndex)
		p.nextToken()

		if p.currentToken().isBond() || p.currentToken().tType == dot {
			p.bondOrder = bondToOrder(p.currentToken())
			p.nextToken()
		} else {
			p.bondOrder = -1
		}

		if p.parseChain() {
			if p.currentToken().tType == rRoundParen {
				p.atomToLinkIndex = p.branchedAtomIndexes[len(p.branchedAtomIndexes) - 1]
				p.branchedAtomIndexes = p.branchedAtomIndexes[:len(p.branchedAtomIndexes) - 1]
				p.nextToken()
				return true
			}
		}
	}

	p.currentTokenIndex = startTokenIndex
	return false
}

func (p *parser) parseBracketAtom() bool {
	startTokenIndex := p.currentTokenIndex

	if p.currentToken().tType == lSquareParen {
		p.nextToken()

		p.parseIsotope()

		if p.currentToken().isSymbol() {
			if p.currentToken().isAromaticSymbol() {
				p.currentAtom.IsAromatic = true
			}
			p.currentAtom.Element = strings.Title(p.currentToken().attribute)
			p.nextToken()

			if p.currentToken().tType == chiral {
				p.nextToken()
			}

			p.parseHCount()
			p.parseCharge()
			p.parseClass()

			if p.currentToken().tType == rSquareParen {
				p.nextToken()
				return true
			} else {
				log.Fatalf(p.currentToken().start.String() + ": close square parenthesis ']' is awaited")
			}
		}
	}

	p.currentTokenIndex = startTokenIndex
	return false
}

func (p *parser) parseIsotope() bool {
	number, ok := p.parseNumber()
	if ok {
		p.currentAtom.Isotope = number
	}
	return ok
}

func (p *parser) parseHCount() bool {
	if p.currentToken().tType == hSymbol {
		if p.currentAtom.Element == "H" {
			log.Fatalf(p.currentToken().start.String() + ": a hydrogen item can't have hydrogens")
		}

		p.nextToken()
		if p.currentToken().tType == digit {
			p.currentAtom.HCount, _ = strconv.Atoi(p.currentToken().attribute)
			p.nextToken()
		} else {
			p.currentAtom.HCount = 1
		}
		return true
	}
	return false
}

func (p *parser) parseCharge() bool {
	if p.currentToken().tType == minus || p.currentToken().tType == plus {
		var chargeString string
		if p.currentToken().tType == minus {
			chargeString = "-"
		}
		p.nextToken()

		if p.currentToken().tType == digit {
			chargeString += p.currentToken().attribute
			p.nextToken()
			if p.currentToken().tType == digit {
				chargeString += p.currentToken().attribute
				p.nextToken()
			}
		} else {
			chargeString += "1"
		}

		p.currentAtom.Charge, _ = strconv.Atoi(chargeString)
		return true
	}

	if p.currentToken().tType == minusMinus {
		p.currentAtom.Charge = -2
		p.nextToken()
		return true
	}

	if p.currentToken().tType == plusPlus {
		p.currentAtom.Charge = 2
		p.nextToken()
		return true
	}

	return false
}

func (p *parser) parseClass() bool {
	if p.currentToken().tType == colon {
		if _, ok := p.parseNumber(); ok {
			return true
		}
	}
	return false
}

func (p *parser) parseNumber() (int, bool) {
	if p.currentToken().tType == digit {
		numberString := p.currentToken().attribute
		p.nextToken()
		for p.currentToken().tType == digit {
			numberString += p.currentToken().attribute
			p.nextToken()
		}
		number, _ := strconv.Atoi(numberString)
		return number, true
	}
	return 0, false
}

func bondToOrder(t token) float32 {
	switch t.tType {
	case dot:
		return 0
	case minus:
		return 1
	case slash:
		return 1
	case backslash:
		return 1
	case colon:
		return 1.5
	case equal:
		return 2
	case hash:
		return 3
	case dollar:
		return 4
	default:
		return -1
	}
}

func (p *parser) handleClosure(index int, info *ringClosureInfo) {
	if p.ringClosures[index] == nil {
		p.ringClosures[index] = info
	} else {
		startInfo := p.ringClosures[index]
		if startInfo.bond == -1 && info.bond == -1 {
			order := p.defaultBond(startInfo.closureAtomIndex, info.closureAtomIndex)
			p.currentGraph.AddBond(startInfo.closureAtomIndex, info.closureAtomIndex, order)
		} else if startInfo.bond == -1 {
				p.currentGraph.AddBond(startInfo.closureAtomIndex, info.closureAtomIndex, info.bond)
		} else if info.bond == -1 {
			p.currentGraph.AddBond(startInfo.closureAtomIndex, info.closureAtomIndex, startInfo.bond)
		} else if startInfo.bond == info.bond {
			p.currentGraph.AddBond(startInfo.closureAtomIndex, info.closureAtomIndex, startInfo.bond)
		} else {
			log.Fatalf(p.currentToken().start.String() + ": start and end of ring closure have differen bond symbols")
		}
		p.ringClosures[index] = nil
	}
}

func (p *parser) addAtom() {
	p.currentGraph.AddAtom(p.currentAtom)
	if len(p.currentGraph.Atoms) > 1 {
		if p.bondOrder != -1 {
			p.currentGraph.AddBond(p.atomToLinkIndex, p.currentAtomIndex, p.bondOrder)
		} else {
			order := p.defaultBond(p.atomToLinkIndex, p.currentAtomIndex)
			p.currentGraph.AddBond(p.atomToLinkIndex, p.currentAtomIndex, order)
		}
		p.atomToLinkIndex = p.currentAtomIndex
	}
	p.currentAtomIndex++
}

func (p parser) defaultBond(firstAtomIndex, secondAtomIndex int) float32 {
	firstAtom, secondAtom := p.currentGraph.Atoms[firstAtomIndex], p.currentGraph.Atoms[secondAtomIndex]
	if firstAtom.IsAromatic && secondAtom.IsAromatic {
		return 1.5
	}
	return 1
}

func (p *parser) fillValence(rewriteHCounts bool) {
	for index := 0; index < len(p.currentGraph.Atoms); index++ {
		if p.currentGraph.Atoms[index].HCount != 0 && !rewriteHCounts {
			continue
		}
		bondOrders := p.currentGraph.Bonds[index]
		var sumOrders float32
		for i := 0; i < len(bondOrders); i++ {
			sumOrders += bondOrders[i]
		}
		val := getValence(p.currentGraph.Atoms[index], int(sumOrders))
		p.currentGraph.Atoms[index].HCount = val - int(sumOrders)
	}
}

func getValence(atom molgraph.Atom, minimum int) int {
	possibleValencies := valence[atom.Element]
	if possibleValencies == nil {
		return 0
	}
	for _, val := range possibleValencies {
		if val >= minimum {
			return val
		}
	}

	log.Fatalf("bonds more than possible met on element: " + atom.Element)
	return 0
}