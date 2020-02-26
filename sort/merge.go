package sort

import "sort"

func Merge(data sort.Interface) {
	merge(data, make([]int, data.Len()), 0)
}

func merge(data sort.Interface, perm []int, from int) {
	if len(perm) < 2 {
		return
	}
	middle := len(perm) / 2
	merge(data, perm[:middle], from)
	merge(data, perm[middle:], from+middle)

	i := 0
	l, r := 0, middle
	for ; l < middle && r < len(perm); i++ {
		if data.Less(from+l, from+r) {
			perm[i] = l
			l++
		} else {
			perm[i] = r
			r++
		}
	}
	for ; l < middle; l++ {
		perm[i] = l
		i++
	}
	for ; r < len(perm); r++ {
		perm[i] = r
		i++
	}

	for i = range perm {
		for ; i != perm[i] && perm[i] != -1 &&  perm[perm[i]] != -1; i, perm[i] = perm[i], -1 {
			data.Swap(from+i, from+perm[i])
		}
	}
}
