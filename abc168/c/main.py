from math import cos, pi, sqrt

a, b, h, m = map(float, input().split())
x = 30 * (h + m / 60)
y = 360 * m / 60
z = abs(x - y)
z = min(z, 360 - z)
print(sqrt(a * a + b * b - 2 * a * b * cos(pi * z / 180)))
