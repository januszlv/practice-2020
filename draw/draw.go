package draw

import (
	"encoding/json"
	"github.com/januszlv/practice-2020/molgraph"
	"io/ioutil"
	"os"
	"os/exec"
)

type MolGroup struct {
	GroupName string
	Group []string
}

func DrawMolGroup(molsGroup []molgraph.MolecularGraph, groupName string, outputFileName string) error {
	var groupSmiles []string
	for _, mol := range molsGroup {
		groupSmiles = append(groupSmiles, mol.SMILES)
	}

	molGroup := MolGroup{groupName, groupSmiles}
	file, err := json.Marshal(molGroup)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("__python__/" + groupName + ".json", file, 0644)
	if err != nil {
		return err
	}
	defer os.Remove("__python__/" + groupName + ".json")

	cmd := exec.Command("__python__/draw.py", groupName + ".json", outputFileName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
