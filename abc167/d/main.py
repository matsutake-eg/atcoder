n, k = map(int, input().split())
a = list(map(int, input().split()))

b = 0
r = k + 1
seen = [-1] * n
seen[0] = 0
p = 1
for i in range(k):
    p = a[p - 1]
    if seen[p - 1] != -1:
        b = seen[p - 1]
        r = i + 1 - b
        for j in range((k - b) % r):
            p = a[p - 1]
        print(p)
        exit()
    seen[p - 1] = i + 1
else:
    print(p)
