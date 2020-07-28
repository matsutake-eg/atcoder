def combination_mod(n, k, mod=10 ** 9 + 7):
    if k > n:
        return 1

    nu, de = 1, 1
    for i in range(k):
        nu = nu * (n - i) % mod
        de = de * (i + 1) % mod
    return nu * pow(de, mod - 2, mod) % mod


n, a, b = map(int, input().split())

mod = 10 ** 9 + 7
ans = pow(2, n, mod) - 1
ans -= combination_mod(n, a)
ans = (ans + mod) % mod
ans -= combination_mod(n, b)
ans = (ans + mod) % mod
print(ans)
