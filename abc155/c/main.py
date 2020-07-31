from collections import Counter

n = int(input())
xc = Counter()
for _ in range(n):
    xc[input()] += 1

mx = xc.most_common(1)[0][1]
ans = []
for k, v in xc.items():
    if v == mx:
        ans.append(k)
ans.sort()
print(*ans, sep="\n")
