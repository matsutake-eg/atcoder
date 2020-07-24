from collections import deque

n, x, y = map(int, input().split())

ans = [0] * n
for i in range(1, n + 1):
    seen = [-1] * (n + 1)
    seen[i] = 0
    q = deque()
    q.append(i)
    while len(q) > 0:
        v = q.popleft()
        ans[seen[v]] += 1
        nv = v - 1
        if nv >= 1 and seen[nv] == -1:
            q.append(nv)
            seen[nv] = seen[v] + 1
        nv = v + 1
        if nv <= n and seen[nv] == -1:
            q.append(nv)
            seen[nv] = seen[v] + 1
        if v == x and seen[y] == -1:
            q.append(y)
            seen[y] = seen[v] + 1
        if v == y and seen[x] == -1:
            q.append(x)
            seen[x] = seen[v] + 1
for v in ans[1:]:
    print(v // 2)
