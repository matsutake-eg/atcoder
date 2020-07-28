from collections import deque


def init_tree(x, par):
    for i in range(x + 1):
        par[i] = i


def find(x, par):
    q = deque()
    q.append(x)
    while len(q) > 0:
        v = q.pop()
        if v == par[v]:
            return v
        q.append(par[v])


def union(x, y, par, rank):
    px, py = find(x, par), find(y, par)
    if px == py:
        return
    if rank[px] < rank[py]:
        par[px] = py
        return
    elif rank[px] == rank[py]:
        rank[px] += 1
    par[py] = px


n, m, k = map(int, input().split())
par = [0] * (n + 1)
rank = [0] * (n + 1)
init_tree(n, par)
eg = [[] for _ in range(n + 1)]
for _ in range(m):
    a, b = map(int, input().split())
    union(a, b, par, rank)
    eg[a].append(b)
    eg[b].append(a)
for _ in range(k):
    a, b = map(int, input().split())
    eg[a].append(b)
    eg[b].append(a)

xs = [0] * (n + 1)
ys = [0] * (n + 1)
for i in range(1, n + 1):
    p = find(i, par)
    xs[p] += 1
    for v in eg[i]:
        if p == find(v, par):
            ys[i] += 1

ans = [-1] * (n + 1)
for i in range(1, n + 1):
    ans[i] += xs[find(i, par)] - ys[i]
print(*ans[1:])
