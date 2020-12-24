H, W, m = map(int, input().split())
xs = [0] * H
ys = [0] * W
s = set()
for _ in range(m):
    h, w = map(int, input().split())
    h -= 1
    w -= 1
    xs[h] += 1
    ys[w] += 1
    s.add((h, w))

mx = max(xs)
my = max(ys)
mxs = set()
for i in range(H):
    if xs[i] == mx:
        mxs.add(i)
mys = set()
for i in range(W):
    if ys[i] == my:
        mys.add(i)
for x in mxs:
    for y in mys:
        if not (x, y) in s:
            print(mx + my)
            exit()
print(mx + my - 1)
