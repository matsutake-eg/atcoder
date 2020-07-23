n = int(input())
ans = (n + 1) * n // 2
for i in range(2, n + 1):
    for j in range(n):
        if i * j > n:
            break
        ans += i * j
print(ans)
