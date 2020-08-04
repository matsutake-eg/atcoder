k = int(input())
x = 0
for i in range(1, k + 1):
    x = x * 10 + 7
    if x % k == 0:
        print(i)
        exit()
    x %= k
print(-1)
