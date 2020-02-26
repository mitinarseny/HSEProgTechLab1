package sort

import "sort"

func Select(s sort.Interface) {
	length := s.Len()
	for bound := 0; bound < length; bound++ {
		minInd := bound
		for j := bound; j < length; j++ {
			if s.Less(j, minInd) {
				minInd = j
			}
		}
		if bound != minInd {
			s.Swap(minInd, bound)
		}
	}
}
