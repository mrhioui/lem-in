package tools

func RemoveMatching(combined [][][]string) [][][]string {
	newCombined := [][][]string{}
	for i, cmb := range combined {
		if len(cmb) == 1 && i != 0 {
			continue
		}
		if !matchingPaths(newCombined, cmb) {
			newCombined = append(newCombined, cmb)
		}
	}
	combined = newCombined
	return combined
}

func matchingPaths(newCombined [][][]string, cmb1 [][]string) bool {
	for _, cmb2 := range newCombined {
		if len(cmb1) == len(cmb2) {
			ok := make([]bool, len(cmb1))
			for i, p1 := range cmb1 {
				p2 := cmb2[i]
				if len(p1) == len(p2) || len(p1) > len(p2) {
					ok = append(ok, true)
				} else {
					ok = append(ok, false)
				}
			}
			count := 0
			for _, b := range ok {
				if b {
					count++
				}
			}
			if count == len(cmb1) {
				return true
			}
		}
	}
	return false
}
