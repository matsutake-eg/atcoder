x, n = map(int, input().split())
p = list(map(int, input().split()))

a = set()
for i in range(-100, 201):
    a.add(i)
for v in p:
    a.remove(v)
a = list(a)
b = list(map(lambda v: abs(v - x), a))
print(a[b.index(min(b))])
