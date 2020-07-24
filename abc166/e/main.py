from collections import Counter

n = int(input())
a = list(map(int, input().split()))


xc = Counter()
yc = Counter()
for i, v in enumerate(a):
    xc[i + v] += 1
    yc[i - v] += 1
ans = 0
for v in xc:
    ans += xc[v] * yc[v]
print(ans)
