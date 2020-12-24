n = int(input())

ans = (1 + n) * n // 2
for i in range(2, n + 1):
    for j in range(1, n):
        if i * j > n:
            break
        ans += i * j
print(ans)
