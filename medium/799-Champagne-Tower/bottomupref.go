package main

func champagneTower(poured int, query_row int, query_glass int) float64 {
	dp := make([]float64, query_row+1)
	dp[0] = float64(poured)
	for row := 0; row < query_row; row++ {
		for i := row; i >= 0; i-- {
			flow := (dp[i] - 1) / 2
			if flow < 0 {
				flow = 0
			}
			dp[i+1] += flow
			dp[i] = flow
		}
	}
	if dp[query_glass] > 1 {
		return 1
	}
	return dp[query_glass]
}
