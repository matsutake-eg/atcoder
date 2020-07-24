n, m = map(int, input().split())
h = [0] + list(map(int, input().split()))
eg = [[] for _ in range(n + 1)]
for _ in range(m):
    a, b = map(int, input().split())
    eg[a].append(b)
    eg[b].append(a)

ans = 0
for i in range(1, n + 1):
    high = True
    for t in eg[i]:
        if h[t] >= h[i]:
            high = False
            break
    if high:
        ans += 1
print(ans)
