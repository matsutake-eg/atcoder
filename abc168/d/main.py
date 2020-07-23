from collections import deque

n, m = map(int, input().split())
eg = [[] for _ in range(n + 1)]

for _ in range(m):
    a, b = map(int, input().split())
    eg[a].append(b)
    eg[b].append(a)

xs = [0] * (n + 1)
q = deque()
q.append(1)
seen = {1}
while len(q) > 0:
    v = q.popleft()
    for t in eg[v]:
        if t in seen:
            continue
        q.append(t)
        seen.add(t)
        xs[t] = v
print("Yes")
print(*xs[2:], sep="\n")
