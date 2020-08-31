package draw

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestDrawMolGroup(t *testing.T) {
	smilesGroups := [][]string{
		{"CC1(C)S[C@@H]2[C@H](NC(=O)[C@H](N)c3ccc(O)cc3)C(=O)N2[C@H]1C(=O)O", "Cc1c(N(C)CS(=O)(=O)O)c(=O)n(-c2ccccc2)n1C", "CCCCC1C(=O)N(c2ccccc2)N(c2ccccc2)C1=O", "Cc1c(N(C)C)c(=O)n(-c2ccccc2)n1C", "CC1(C)S[C@@H]2[C@H](NC(=O)Cc3ccccc3)C(=O)N2[C@H]1C(=O)O", "Cc1onc(-c2ccccc2)c1C(=O)N[C@@H]1C(=O)N2[C@@H](C(=O)O)C(C)(C)S[C@H]12", "CC1(C)SC2C(NC(=O)C(N)c3ccccc3)C(=O)N2C1C(=O)O"},
		{"CC(C)NCC(O)COc1ccc(CC(N)=O)cc1", "CC(C)NCC(O)COc1ccc(COCCOC(C)C)cc1", "CNC[C@H](O)c1cccc(O)c1", "NCCc1ccc(O)c(O)c1", "NCC(O)c1ccc(O)c(O)c1", "CNC[C@H](O)c1ccc(O)c(O)c1"},
		{"Cc1cc(=O)n(-c2ccccc2)n1C"},
	}

	_ = os.Mkdir("mol_groups", 0777)
	for i, smilesGroup := range smilesGroups {
		if err := DrawMolGroup(smilesGroup, "group" + strconv.Itoa(i + 1), "mol_groups/group" + strconv.Itoa(i + 1)); err != nil {
			fmt.Println(err)
			t.Fatalf("Drawing test failed!")
		}
	}

	t.Log("Drawing test was successfully executed!")
}
