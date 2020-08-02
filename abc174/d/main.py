n = int(input())
c = input()

cnt1 = 0
for v in c[: n // 2 + 1]:
    if v == "W":
        cnt1 += 1
cnt2 = 0
for v in c[n // 2 - 1 :]:
    if v == "R":
        cnt2 += 1
print(min(cnt1, cnt2))
