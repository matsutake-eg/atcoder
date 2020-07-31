from collections import Counter

n = int(input())
a = list(map(int, input().split()))

ac = Counter(a)
if len(ac) == n:
    print("YES")
else:
    print("NO")
