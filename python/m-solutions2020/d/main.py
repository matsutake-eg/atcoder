n = int(input())
a = list(map(int, input().split()))

ans = 1000
for i in range(n - 1):
    if a[i + 1] > a[i]:
        x = ans // a[i]
        ans -= a[i] * x
        ans += a[i + 1] * x
print(ans)
