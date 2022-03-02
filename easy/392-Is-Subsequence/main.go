package main

func isSubsequence(s string, t string) bool {
	nS, nT := len(s), len(t)
	if nS == 0 {
		return true
	}
	for iS, iT := 0, 0; iT < nT; iT++ {
		if s[iS] == t[iT] {
			iS++
			if iS == nS {
				return true
			}
		}
	}
	return false
}

func followUp(s string, t string) bool {
	indices := make([][]int, 26)
	for i := range indices {
		indices[i] = make([]int, 0)
	}
	for i, r := range t {
		indices[r-'a'] = append(indices[r-'a'], i)
	}
	i := -1
	for _, r := range s {
		if i = binarySearch(indices[r-'a'], i, 0, len(indices[r-'a'])-1); i < 0 {
			return false
		}
	}
	return true
}

func binarySearch(indices []int, index int, start int, end int) int {
	if start > end {
		return -1
	}
	middle := (start + end) / 2
	if indices[middle] > index {
		if s := binarySearch(indices, index, start, middle-1); s >= 0 && s < indices[middle] {
			return s
		}
		return indices[middle]
	}
	return binarySearch(indices, index, middle+1, end)
}
