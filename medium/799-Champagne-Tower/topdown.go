package main

func champagneTower(poured int, query_row int, query_glass int) float64 {
	cache := map[int]float64{0: float64(poured)}
	flow := dp(query_row, query_glass, cache)
	if flow > 1 {
		return 1
	}
	return flow
}

func dp(row int, glass int, cache map[int]float64) float64 {
	if glass < 0 || glass > row {
		return 0
	}
	index := row*(1+row)/2 + glass
	if flow, ok := cache[index]; ok {
		return flow
	}
	upperLeft, upperRight := dp(row-1, glass-1, cache), dp(row-1, glass, cache)
	cache[index] = max((upperLeft-1)/2, 0) + max((upperRight-1)/2, 0)
	return cache[index]
}

func max(a float64, b float64) float64 {
	if a < b {
		return b
	}
	return a
}
