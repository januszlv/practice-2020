package pycall

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func Call(input, pyFile, inFile, outFile string) (string, error) {
	fin, err := os.Create(inFile)
	defer fin.Close()
	if err != nil {
		return "", err
	}
	fin.Write([]byte(input))
	cmd := exec.Command(pyFile, inFile, outFile)
	fmt.Println("command args:", cmd.Args)

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	fout, err := os.Open(outFile)
	if err != nil {
		return "", err
	}

	res, err := ioutil.ReadAll(fout)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
