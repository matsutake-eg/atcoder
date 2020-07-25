n, m = map(int, input().split())
x = n * (n - 1) // 2 if n > 1 else 0
y = m * (m - 1) // 2 if m > 1 else 0
print(x + y)
