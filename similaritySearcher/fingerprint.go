package similaritySearcher

import (
	"fmt"
	"github.com/januszlv/practice-2020/similaritySearcher/pycall"
	"strconv"
	"gonum.org/v1/gonum/mat"
)

type Fingerprint struct {
	bitstring mat.Vector
}

func GetFingerprint(smiles string) (*Fingerprint, error) {
	bitstringStr, err := pycall.Call(smiles, "similaritySearcher/bitstring.py", "inp", "out")
	if err != nil {
		return nil, err
	}

	var fp Fingerprint
	fp.bitstring, err = stringToVector(bitstringStr)
	if err != nil {
		return nil, err
	}

	return &fp, nil
}

func stringToVector(str string) (mat.Vector, error) {
	var res []float64
	for i := 0; i < len(str); i++ {
		float, err := strconv.ParseFloat(string(str[i]), 64)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		res = append(res, float)
	}

	return mat.NewVecDense(len(res), res), nil
}

func TanimotoSimilarity(fp *Fingerprint, fp2 *Fingerprint) float64 {
	numer := mat.Dot(fp.bitstring, fp2.bitstring)
	if numer == 0.0 {
		return 0.0
	}

	denom := mat.Dot(fp.bitstring, fp.bitstring) + mat.Dot(fp2.bitstring, fp2.bitstring) - numer
	if denom == 0.0 {
		return 0.0
	}

	return numer / denom
}

func BulkTanimoto(fp1 *Fingerprint, fps []*Fingerprint) []float64 {
	var res []float64

	for _, fp2 := range fps {
		simVal := TanimotoSimilarity(fp1, fp2)
		res = append(res, simVal)
	}

	return res
}

func GetTanimotoVec(fps []*Fingerprint) []float64 {
	var tanimotoVec []float64
	for i := 0; i < len(fps); i++ {
		for j := i + 1; j < len(fps); j++ {
			tanimotoVec = append(tanimotoVec, TanimotoSimilarity(fps[i], fps[j]))
		}
	}

	return tanimotoVec
}