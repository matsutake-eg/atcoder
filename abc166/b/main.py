n, k = map(int, input().split())
xs = set()
for i in range(k):
    d = int(input())
    xs = xs.union(set(map(int, input().split())))
print(n - len(xs))
