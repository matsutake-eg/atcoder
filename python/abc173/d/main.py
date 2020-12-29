n = int(input())
a = list(map(int, input().split()))

a.sort(reverse=True)
ans = a[0]
cnt = 1
for i in range(1, n):
    if cnt == n - 1:
        break
    ans += a[i]
    cnt += 1
    if cnt == n - 1:
        break
    ans += a[i]
    cnt += 1
print(ans)
