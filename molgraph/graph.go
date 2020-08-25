package molgraph

import (
	"fmt"
	"strconv"
)

type MolecularGraph struct {
	Atoms []Atom
	Bonds [][]float32
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
	graph.Atoms = append(graph.Atoms, atom)
	for i := 0; i < len(graph.Bonds); i++ {
		graph.Bonds[i] = append(graph.Bonds[i], 0)
	}
	graph.Bonds = append(graph.Bonds, make([]float32, len(graph.Atoms)))
}

func (graph *MolecularGraph) AddBond(firstAtomIndex, secondAtomIndex int, order float32) {
	graph.Bonds[firstAtomIndex][secondAtomIndex] = order
	graph.Bonds[secondAtomIndex][firstAtomIndex] = order
}

func (graph *MolecularGraph) AddExplicitHydrogens() {
	hAtom := Atom{
		Element: "H",
		Isotope: -1,
	}
	currentAtomIndex := len(graph.Atoms)
	for i := 0; i < len(graph.Atoms); i++ {
		for graph.Atoms[i].HCount > 0 {
			graph.AddAtom(hAtom)
			graph.AddBond(i, currentAtomIndex, 1)
			currentAtomIndex++
			graph.Atoms[i].HCount--
		}
	}
}

func (graph MolecularGraph) String() string {
	var str string
	for i := 0; i < len(graph.Atoms); i++ {
		str += graph.Atoms[i].String() + "\n"
	}
	str += "\n"
	for i := 0; i < len(graph.Atoms); i++ {
		str += strconv.Itoa(i+1) + "\t[ "
		for j := 0; j < len(graph.Bonds[i]); j++ {
			str += fmt.Sprintf("%0.1f", graph.Bonds[i][j]) + " "
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
