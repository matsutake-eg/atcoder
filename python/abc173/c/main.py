h, w, k = map(int, input().split())
c = [input() for _ in range(h)]

ans = 0
for hi in range(2 ** h):
    hs = set()
    for hj in range(h):
        if (hi >> hj) & 1:
            hs.add(hj)
    for wi in range(2 ** w):
        ws = set()
        for wj in range(w):
            if (wi >> wj) & 1:
                ws.add(wj)
        cnt = 0
        for i in range(h):
            if i in hs:
                continue
            for j in range(w):
                if j in ws:
                    continue
                if c[i][j] == "#":
                    cnt += 1
        if cnt == k:
            ans += 1
print(ans)
