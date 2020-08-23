R, C, k = map(int, input().split())
xs = [[0] * C for _ in range(R)]
for _ in range(k):
    r, c, v = map(int, input().split())
    xs[r - 1][c - 1] = v

dp = [[[0] * (C + 1) for _ in range(R + 1)] for _ in range(4)]
for i in range(R):
    for j in range(C):
        for k in range(2, -1, -1):
            dp[k + 1][i][j] = max(dp[k + 1][i][j], dp[k][i][j] + xs[i][j])
        for k in range(4):
            dp[k][i][j + 1] = max(dp[k][i][j + 1], dp[k][i][j])
            dp[0][i + 1][j] = max(dp[0][i + 1][j], dp[k][i][j])
ans = 0
for k in range(4):
    ans = max(ans, dp[k][R - 1][C - 1])
print(ans)
