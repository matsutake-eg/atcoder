n = int(input())
a = list(map(int, input().split()))
a.sort(reverse=True)
ans = a[0]
idx = 1
for i in range(1, n):
    if idx == n - 1:
        break
    ans += a[i]
    idx += 1
    if idx == n - 1:
        break
    ans += a[i]
    idx += 1
print(ans)
