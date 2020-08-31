package draw

import (
	"fmt"
	"github.com/januszlv/practice-2020/molgraph"
	"os"
	"strconv"
	"testing"
)

func TestDrawMolGroup(t *testing.T) {
	_ = os.Mkdir("mol_groups", 0777)
	var molsGroups [][]molgraph.MolecularGraph
	molsGroups = append(molsGroups, []molgraph.MolecularGraph{
		{SMILES: "CC1(C)S[C@@H]2[C@H](NC(=O)[C@H](N)c3ccc(O)cc3)C(=O)N2[C@H]1C(=O)O"},
		{SMILES: "Cc1c(N(C)CS(=O)(=O)O)c(=O)n(-c2ccccc2)n1C"},
		{SMILES: "CCCCC1C(=O)N(c2ccccc2)N(c2ccccc2)C1=O"},
		{SMILES: "Cc1c(N(C)C)c(=O)n(-c2ccccc2)n1C"},
		{SMILES: "CC1(C)S[C@@H]2[C@H](NC(=O)Cc3ccccc3)C(=O)N2[C@H]1C(=O)O"},
		{SMILES: "Cc1onc(-c2ccccc2)c1C(=O)N[C@@H]1C(=O)N2[C@@H](C(=O)O)C(C)(C)S[C@H]12"},
		{SMILES: "CC1(C)SC2C(NC(=O)C(N)c3ccccc3)C(=O)N2C1C(=O)O"},
	})
	molsGroups = append(molsGroups, []molgraph.MolecularGraph{
		{SMILES: "CC(C)NCC(O)COc1ccc(CC(N)=O)cc1"},
		{SMILES: "CC(C)NCC(O)COc1ccc(COCCOC(C)C)cc1"},
		{SMILES: "CNC[C@H](O)c1cccc(O)c1"},
		{SMILES: "NCCc1ccc(O)c(O)c1"},
		{SMILES: "NCC(O)c1ccc(O)c(O)c1"},
		{SMILES: "CNC[C@H](O)c1ccc(O)c(O)c1"},
	})
	molsGroups = append(molsGroups, []molgraph.MolecularGraph{
		{SMILES: "Cc1cc(=O)n(-c2ccccc2)n1C"},
	})

	for i, molsGroup := range molsGroups {
		if err := DrawMolGroup(molsGroup, "group" + strconv.Itoa(i + 1), "mol_groups/group" + strconv.Itoa(i + 1)); err != nil {
			fmt.Println(err)
			t.Fatalf("Drawing test failed!")
		}
	}

	t.Log("Drawing test was successfully executed!")
}
