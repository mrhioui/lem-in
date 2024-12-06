package lemin

// Combinations of non-intersecting paths
func ValidPathsCombs() (validCombs [][][]string) {
	for i := range Paths {
		var currComb [][]string
		currComb = append(currComb, Paths[i])

		for j, path := range Paths {
			if i == j {
				continue
			}
			isInter := false
			for _, currCombPath := range currComb {
				if pathsIntersect(currCombPath, path) {
					isInter = true
					break
				}
			}
			if !isInter {
				currComb = append(currComb, path)
			}
		}
		validCombs = append(validCombs, currComb)
	}
	return validCombs
}

func pathsIntersect(path1, path2 []string) bool {
	for _, roomName := range path1[1 : len(path1)-1] {
		if containsRoom(path2, roomName) {
			return true
		}
	}
	return false
}

func containsRoom(path []string, roomName string) bool {
	for _, room := range path[1 : len(path)-1] {
		if room == roomName {
			return true
		}
	}
	return false
}

func SortInCombinations(validCombs [][][]string) {
	for _, cmb := range validCombs {
		SortPaths(cmb)
	}
}
