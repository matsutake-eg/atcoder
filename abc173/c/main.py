h, w, k = map(int, input().split())
c = [""] * h
for i in range(h):
    c[i] = input()

ans = 0
for i in range(2 ** h):
    hst = set()
    for ii in range(h):
        if (i >> ii) & 1:
            hst.add(ii)
    for j in range(2 ** w):
        wst = set()
        for jj in range(w):
            if (j >> jj) & 1:
                wst.add(jj)
        cnt = 0
        for x in range(h):
            for y in range(w):
                if x in hst or y in wst:
                    continue
                if c[x][y] == "#":
                    cnt += 1
        if cnt == k:
            ans += 1
print(ans)
