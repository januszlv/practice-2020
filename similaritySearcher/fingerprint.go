package similaritySearcher

type Fingerprint interface {
	MatchFingerprints(bitstring *fingerprint) float64
}

type fingerprint struct {
	bitstring uint64
}

func GetFingerprint(molGraph *MoleculatGraph) uint64 {
	// graph to bitstring stuff
	var bitstring uint64
	return bitstring
}

//MatchFingerprints matching two bitstrings and returns Tanimoto index
func (fingerprint *fingerprint) MatchFingerprints(fingerprint2 *fingerprint) float64 {
	a := unitCounter(fingerprint.bitstring)
	b := unitCounter(fingerprint2.bitstring)
	c := unitCounter(fingerprint.bitstring & fingerprint2.bitstring)

	return c/(a + b - c)
}

func unitCounter(n uint64) float64 {
	bit := uint64(1)
	count := float64(0)
	for bit != 0 {
		if bit & n != 0 {
			count++
		}
		bit <<= 1
	}

	return count
}