package bin

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func levenshteinDistance(s1, s2 string) int {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := range dp[0] {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}

func Similarity(s1, s2 string) float64 {
	l := len(s1) + len(s2)
	distance := levenshteinDistance(s1, s2)
	return 1.0 - float64(distance)/float64(l)
}
