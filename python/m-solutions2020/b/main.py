a, b, c = map(int, input().split())
k = int(input())

cnt = 0
while b <= a:
    b *= 2
    cnt += 1
while c <= b:
    c *= 2
    cnt += 1
print("Yes" if cnt <= k else "No")
