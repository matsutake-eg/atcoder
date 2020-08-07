n, q = map(int, input().split())
bit = [0] * (n + 1)


def fenwick_add(x, w):
    while x <= n:
        bit[x] += w
        x += x & -x


def fenwick_sum(x):
    sm = 0
    while x > 0:
        sm += bit[x]
        x -= x & -x
    return sm


c = list(map(int, input().split()))
p = [-1] * (n + 1)
ps = [[] for _ in range(n + 1)]
for i, v in enumerate(c):
    l = p[v]
    if l != -1:
        ps[l].append(i)
    p[v] = i

qs = [[] for _ in range(n + 1)]
for i in range(q):
    l, r = map(int, input().split())
    qs[l - 1].append([r - 1, i])

ans = [0] * q
for i in range(n - 1, -1, -1):
    for v in ps[i]:
        fenwick_add(v, 1)
    for v in qs[i]:
        r, qi = v
        ans[qi] = r - i + 1 - fenwick_sum(r)
print(*ans, sep="\n")
