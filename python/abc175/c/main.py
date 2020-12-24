x, k, d = map(int, input().split())

x = abs(x)
m = min(k, x // d)
k -= m
x -= m * d

if k % 2 == 0:
    print(x)
else:
    print(d - x)
