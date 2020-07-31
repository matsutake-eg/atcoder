from itertools import accumulate

n, k = map(int, input().split())
p = list(map(int, input().split()))

xs = [0] + list(accumulate(map(lambda v: (1 + v) * v // 2 / v, p)))
ans = 0
for i in range(n - k + 1):
    ans = max(ans, xs[i + k] - xs[i])
print(ans)
