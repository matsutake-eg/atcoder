from math import gcd


def prime_fac(x):
    pfs = set()
    for i in range(2, x):
        if i * i > x:
            break
        while x % i == 0:
            pfs.add(i)
            x //= i
    if x > 1:
        pfs.add(x)
    return pfs


n = int(input())
a = list(map(int, input().split()))

g = a[0]
for v in a[1:]:
    g = gcd(g, v)

if g != 1:
    print("not coprime")
    exit()

s = set()
for v in a:
    p = prime_fac(v)
    ls, lp = len(s), len(p)
    s.update(p)
    if len(s) < ls + lp:
        print("setwise coprime")
        exit()
print("pairwise coprime")
