from itertools import accumulate

n = int(input())
a = list(map(int, input().split()))

sa = list(accumulate(a))
mod = 10 ** 9 + 7
ans = 0
for i in range(n - 1):
    ans += a[i] * (sa[-1] - sa[i])
    ans %= mod
print(ans)
