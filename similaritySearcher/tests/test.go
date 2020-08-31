package tests

import (
	"fmt"
	"similaritySearcher"
)

func unitCounter(n uint64) float64 {
	bit := uint64(1)
	count := float64(0)
	for bit != 0 {
		if bit&n != 0 {
			count++
		}
		bit <<= 1
	}

	return count
}

func unitCounterTest(n uint64, res float64) bool {
	return unitCounter(n) == res
}

func tanimotoIdx(n, m uint64) float64 {
	a := unitCounter(n)
	b := unitCounter(m)
	c := unitCounter(n & m)

	return c / (a + b - c)
}

func tanimotoIdxTest(n, m uint64, res float64) bool {
	return tanimotoIdx(n, m) == res
}

func main() {
	fmt.Println("unitCounter test: ", unitCounterTest(13, 3))
	fmt.Println("unitCounter test: ", unitCounterTest(4294967295, 32))
	fmt.Println("unitCounter test: ", unitCounterTest(1046285, 14))
	fmt.Println("unitCounter test: ", unitCounterTest(9825375, 15))
	fmt.Println("unitCounter test: ", unitCounterTest(8, 1))

	fmt.Println("tanimotoIdx test: ", tanimotoIdxTest(13, 10, 0.25))

	fp, _ := similaritySearcher.GetFingerprint("c1ccccc1")
	fp2, _ := similaritySearcher.GetFingerprint("c1ccccc1C(=O)O")

	// fmt.Println("fp is ", fp)

	var fps []string
	fps = append(fps, fp, fp2)

	fmt.Println(similaritySearcher.GetTanimotoVec(fps))

}

//1 -> 10; 10 -> 100; 101 -> 1010
// 101 & 1 -> 1; 1101 & 1010 -> 1000
// a = 1101 (3), b = 1010 (2), c = 1000 (1)
