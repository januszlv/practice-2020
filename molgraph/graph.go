package molgraph

import (
	"fmt"
	"strconv"
)

type MolecularGraph struct {
	atoms []Atom
	bonds [][]float32
}

type Atom struct {
	Element string
	HCount  int
	Charge  int
	Isotope int
	IsAromatic bool
}

func DefaultAtom() Atom {
	var a Atom
	a.Element = "*"
	a.Isotope = -1
	return a
}

func (graph *MolecularGraph) AddAtom(atom Atom) {
	graph.atoms = append(graph.atoms, atom)
	for i := 0; i < len(graph.bonds); i++ {
		graph.bonds[i] = append(graph.bonds[i], 0)
	}
	graph.bonds = append(graph.bonds, make([]float32, len(graph.atoms)))
}

func (graph *MolecularGraph) AddBond(firstAtomIndex, secondAtomIndex int, order float32) {
	graph.bonds[firstAtomIndex][secondAtomIndex] = order
	graph.bonds[secondAtomIndex][firstAtomIndex] = order
}

func (graph *MolecularGraph) AddExplicitHydrogens() {
	hAtom := Atom{
		Element: "H",
		Isotope: -1,
	}
	currentAtomIndex := len(graph.atoms)
	for i := 0; i < len(graph.atoms); i++ {
		for graph.atoms[i].HCount > 0 {
			graph.AddAtom(hAtom)
			graph.AddBond(i, currentAtomIndex, 1)
			currentAtomIndex++
			graph.atoms[i].HCount--
		}
	}
}

func (graph MolecularGraph) String() string {
	var str string
	for i := 0; i < len(graph.atoms); i++ {
		str += graph.atoms[i].String() + "\n"
	}
	str += "\n"
	for i := 0; i < len(graph.atoms); i++ {
		str += strconv.Itoa(i+1) + "\t[ "
		for j := 0; j < len(graph.bonds[i]); j++ {
			str += fmt.Sprintf("%0.1f", graph.bonds[i][j]) + " "
		}
		str += "]\n"
	}
	return str
}

func (atom Atom) String() string {
	isAromatic := "false"
	if atom.IsAromatic {
		isAromatic = "true"
	}
	return "{" + atom.Element + " " + strconv.Itoa(atom.HCount) + " " + strconv.Itoa(atom.Charge) + " " + strconv.Itoa(atom.Isotope) + " " + isAromatic + "}"
}
