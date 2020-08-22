n, k = map(int, input().split())
p = [0] + list(map(int, input().split()))
c = [0] + list(map(int, input().split()))

g = [v for v in range(n + 1)]
s = [0] * (n + 1)
r = [0] * (n + 1)
seen = set()
for i in range(1, n + 1):
    if i in seen:
        s[i] = s[g[i]]
        r[i] = r[g[i]]
        continue
    ni = i
    sm = 0
    cnt = 0
    while True:
        ni = p[ni]
        sm += c[ni]
        cnt += 1
        g[ni] = i
        seen.add(ni)
        if ni == i:
            break
    s[i] = sm
    r[i] = cnt

ans = -(10 ** 9 + 1)
for i in range(1, n + 1):
    ni = i
    sm = 0
    cnt = 0
    while True:
        ni = p[ni]
        sm += c[ni]
        cnt += 1
        ans = max(ans, sm, sm + (k - cnt) // r[i] * s[i])
        if ni == i or cnt == k:
            break
print(ans)
