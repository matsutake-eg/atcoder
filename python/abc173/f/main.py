from functools import reduce

n = int(input())

E = reduce(lambda ac, v: ac + v * (n - v + 1), range(1, n + 1), 0)
V = 0
for _ in range(n - 1):
    a, b = map(int, input().split())
    if a > b:
        a, b = b, a
    V += a * (n - b + 1)
print(E - V)
