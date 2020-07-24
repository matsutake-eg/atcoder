n, m = map(int, input().split())
a = list(map(int, input().split()))

a.sort(reverse=True)
if len(list(filter(lambda v: v * 4 * m >= sum(a), a[:m]))) == m:
    print("Yes")
else:
    print("No")
