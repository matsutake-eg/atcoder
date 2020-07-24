from itertools import combinations_with_replacement

n, m, q = map(int, input().split())
a = [0] * q
b = [0] * q
c = [0] * q
d = [0] * q
for i in range(q):
    a[i], b[i], c[i], d[i] = map(int, input().split())

ans = 0
for v in combinations_with_replacement(list(range(1, m + 1)), n):
    sm = 0
    for i in range(q):
        if v[b[i] - 1] - v[a[i] - 1] == c[i]:
            sm += d[i]
    ans = max(ans, sm)
print(ans)
