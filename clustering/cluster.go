package clustering

import (
	"errors"
	"sort"
)

const distThresh = 0.7

func TanimotoDistance(tanimotoIdx float64) float64 {
	return 1 - tanimotoIdx
}

func Cluster(data []float64, nPts int) ([][]int, error) {
	if len(data) > (nPts * (nPts - 1) / 2) {
		return nil, errors.New("Matrix is too long")
	}

	nbrLists := make([][]int, 14)

	dmIdx := 0
	for i := 0; i < nPts; i++ {
		for j := 0; j < i; j++ {
			dij := TanimotoDistance(data[dmIdx])
			dmIdx++
			if dij <= distThresh {
				nbrLists[i] = append(nbrLists[i], j)
				nbrLists[j] = append(nbrLists[j], i)
			}
		}
	}

	var tLists [][]int
	for x, y := range nbrLists {
		tLists = append(tLists, []int{len(y), x})
	}

	sort.Slice(tLists, func(i, j int) bool {
		diff := tLists[i][0] - tLists[j][0]
		if diff > 0 {
			return true
		} else if diff < 0 {
			return false
		} else {
			return tLists[i][1] > tLists[j][1]
		}
	})

	var res [][]int
	seen := make([]bool, nPts)

	for len(tLists) != 0 {
		var tList []int
		tList, tLists = tLists[0], tLists[1:]
		idx := tList[1]
		if seen[idx] {
			continue
		}

		tRes := []int{idx}
		for _, nbr := range nbrLists[idx] {
			if !seen[nbr] {
				tRes = append(tRes, nbr)
				seen[nbr] = true
			}
		}

		res = append(res, tRes)
	}

	return res, nil
}