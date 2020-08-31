package similaritySearcher

import (
	"fmt"
	"similaritySearcher/pycall"
	"strconv"

	"gonum.org/v1/gonum/mat"
)

type Fingerprint interface {
	//MatchFingerprints(bitstring *fingerprint) float64
	TanimotoSimilarity(fp2 *fingerprint) float64
}

type fingerprint struct {
	bitstring mat.Vector
}

func GetFingerprint(smiles string) (string, error) {
	bitstring, err := pycall.Call(smiles, "./bitstring.py", "inp", "out")
	if err != nil {
		return "", err
	}

	return bitstring, nil
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

func (fp *fingerprint) TanimotoSimilarity(fp2 *fingerprint) float64 {
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

func GetTanimotoVec(fingerprints []string) []float64 {
	var fpVec []fingerprint
	var fp fingerprint
	for i := 0; i < len(fingerprints); i++ {
		fp.bitstring, _ = stringToVector(fingerprints[i])
		fmt.Println("FINGERPRINT ", fp.bitstring)
		fpVec = append(fpVec, fp)
	}

	var tanimotoVec []float64
	for i := 0; i < len(fpVec); i++ {
		for j := i; j < len(fpVec); j++ {
			if i == j {
				tanimotoVec = append(tanimotoVec, 1)
			} else {
				tanimotoVec = append(tanimotoVec, fpVec[i].TanimotoSimilarity(&fpVec[j]))
			}
		}
	}

	return tanimotoVec
}

//MatchFingerprints matching two bitstrings and returns Tanimoto index
//func (fingerprint *fingerprint) MatchFingerprints(fingerprint2 *fingerprint) float64 {
//	a := unitCounter(fingerprint.bitstring)
//	b := unitCounter(fingerprint2.bitstring)
//	c := unitCounter(fingerprint.bitstring & fingerprint2.bitstring)
//
//	return c / (a + b - c)
//}
//
//func unitCounter(n uint64) float64 {
//	bit := uint64(1)
//	count := float64(0)
//	for bit != 0 {
//		if bit&n != 0 {
//			count++
//		}
//		bit <<= 1
//	}
//
//	return count
//}
