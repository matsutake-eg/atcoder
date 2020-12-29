n = int(input())
a = list(map(int, input().split()))

ans = 0
x = a[0]
for v in a[1:]:
    if v < x:
        ans += x - v
    else:
        x = v
print(ans)
