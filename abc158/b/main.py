n, a, b = map(int, input().split())
x = n // (a + b)
print(x * a + min(n % (a + b), a))
