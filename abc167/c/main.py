n, m, x = map(int, input().split())
c = [0] * n
a = [0] * n
for i in range(n):
    xs = list(map(int, input().split()))
    c[i] = xs[0]
    a[i] = xs[1:]
ans = 10 ** 9
for i in range(2 ** n):
    csum = 0
    asum = [0] * m
    for j in range(n):
        if (i >> j) & 1:
            csum += c[j]
            for k in range(m):
                asum[k] += a[j][k]
    if len(list(filter(lambda v: v >= x, asum))) == m:
        ans = min(ans, csum)
print(-1 if ans == 10 ** 9 else ans)
