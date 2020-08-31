package draw

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
)

type MolGroup struct {
	GroupName string
	Group []string
}

func DrawMolGroup(smilesGroup []string, groupName string, outputFileName string) error {
	molGroup := MolGroup{groupName, smilesGroup}
	file, err := json.Marshal(molGroup)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("__python__/files/" + groupName + ".json", file, 0644)
	if err != nil {
		return err
	}
	defer os.Remove("__python__/files/" + groupName + ".json")

	cmd := exec.Command("__python__/draw.py", groupName + ".json", outputFileName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
