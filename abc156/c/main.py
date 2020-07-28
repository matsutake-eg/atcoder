n = int(input())
x = list(map(int, input().split()))

ans = 10 ** 9
for i in range(1, 101):
    sm = 0
    for v in x:
        sm += (v - i) ** 2
    ans = min(ans, sm)
print(ans)
