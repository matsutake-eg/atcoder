def prime_fac(x):
    pfs = {1}
    for i in range(2, x):
        if i * i > x:
            break
        while x % i == 0:
            pfs.add(i)
            x //= i
    pfs.add(x)
    return pfs


n = int(input())
xs = prime_fac(n)
xs.remove(1)
ans = 0
for v in xs:
    t = v
    while n % t == 0:
        n //= t
        t *= v
        ans += 1
print(ans)
