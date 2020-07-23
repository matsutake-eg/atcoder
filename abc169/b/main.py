n = int(input())
a = list(map(int, input().split()))
a.sort()
if a[0] == 0:
    print(0)
else:
    ans = 1
    for v in a:
        if v > 10 ** 18 // ans:
            print(-1)
            exit()
        ans *= v
    print(ans)
