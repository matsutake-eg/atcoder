n = int(input())
a = list(map(int, input().split()))

xs = [0] * (max(a) + 1)

for v in a:
    if xs[v] == 0:
        for j in range(v, len(xs), v):
            xs[j] += 1
    else:
        xs[v] += 1
print(len(list(filter(lambda v: xs[v] == 1, a))))
