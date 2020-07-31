h, n = map(int, input().split())
xs = [0] * n
for i in range(n):
    xs[i] = list(map(int, input().split()))

dp = [10 ** 9] * (h + 1)
dp[0] = 0
for i in range(h):
    for j in range(n):
        x = min(i + xs[j][0], h)
        dp[x] = min(dp[x], dp[i] + xs[j][1])
print(dp[h])
