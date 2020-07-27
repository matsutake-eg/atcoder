n, m = map(int, input().split())
s = [0] * m
c = [0] * m
for i in range(m):
    s[i], c[i] = map(int, input().split())

for i in range(0, 1000):
    x = str(i)
    if len(x) != n:
        continue
    for j in range(m):
        if s[j] > len(x) or int(x[s[j] - 1]) != c[j]:
            break
    else:
        print(x)
        exit()
print(-1)
