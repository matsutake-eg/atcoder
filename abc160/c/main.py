k, n = map(int, input().split())
a = list(map(int, input().split()))

x = k - a[-1] + a[0]
for i in range(n - 1):
    x = max(x, a[i + 1] - a[i])
print(k - x)
