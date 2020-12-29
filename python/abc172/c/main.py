from itertools import accumulate

n, m, k = map(int, input().split())
a = list(map(int, input().split()))
b = list(map(int, input().split()))

sa = [0] + list(accumulate(a))
sb = [0] + list(accumulate(b))

ans = 0
p = m
for i in range(n + 1):
    for j in range(p, -1, -1):
        if sa[i] + sb[j] <= k:
            ans = max(ans, i + j)
            p = j
            break
print(ans)
