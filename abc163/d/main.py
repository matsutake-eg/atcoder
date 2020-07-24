n, k = map(int, input().split())
mod = 10 ** 9 + 7
ans = 0
for i in range(k, n + 2):
    mn = (i - 1) * i // 2
    mx = (n + n - i + 1) * i // 2
    ans += mx - mn + 1
    ans %= mod
print(ans)
