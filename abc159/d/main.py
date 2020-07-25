from collections import Counter

n = int(input())
a = list(map(int, input().split()))

ac = Counter(a)
sm = 0
for v in ac.values():
    if v > 1:
        sm += v * (v - 1) // 2
for v in a:
    print(sm - ac[v] * (ac[v] - 1) // 2 + (ac[v] - 1) * (ac[v] - 2) // 2)
