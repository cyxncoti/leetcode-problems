package main

func champagneTower(poured int, query_row int, query_glass int) float64 {
	dp := []float64{float64(poured)}
	for row := 0; row < query_row; row++ {
		nextLevel := make([]float64, row+2)
		for i := 0; i < row+1; i++ {
			flow := (dp[i] - 1) / 2
			if flow <= 0 {
				continue
			}
			nextLevel[i] += flow
			nextLevel[i+1] += flow
		}
		dp = nextLevel
	}
	if dp[query_glass] > 1 {
		return 1
	}
	return dp[query_glass]
}
