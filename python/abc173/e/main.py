from functools import reduce

n, k = map(int, input().split())
a = list(map(int, input().split()))

mod = 10 ** 9 + 7
if n == k:
    print(reduce(lambda ac, v: ac * v % mod, a, 1))
    exit()

a.sort(key=lambda v: abs(v), reverse=True)
if len(list(filter(lambda v: v < 0, a))) == n and k % 2 == 1:
    print(reduce(lambda ac, v: ac * v % mod, a[n - k :], 1))
    exit()

ans = 1
mn = 0
for v in a[:k]:
    if v < 0:
        mn += 1
    ans *= abs(v)
    ans %= mod

if mn % 2 == 0:
    print(ans)
    exit()

mnp, mnn = 0, 0
for v in a[k - 1 :: -1]:
    if mnp != 0 and mnn != 0:
        break
    elif mnp == 0 and v > 0:
        mnp = v
    elif mnn == 0 and v < 0:
        mnn = v
mxp, mxn = 0, 0
for v in a[k:]:
    if mxp != 0 and mxn != 0:
        break
    elif mxp == 0 and v > 0:
        mxp = v
    elif mxn == 0 and v < 0:
        mxn = v
x, y = mxp, mnn
if mnp != 0 and mxn * mnn > mxp * mnp:
    x, y = mxn, mnp
ans *= pow(abs(y), mod - 2, mod)
ans %= mod
ans *= abs(x)
ans %= mod
print(ans)
