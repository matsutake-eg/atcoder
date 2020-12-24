n = int(input())
c = input()

r = 0
for v in c:
    if v == "R":
        r += 1
ans = min(r, n - r)
w = 0
for v in c:
    if v == "W":
        w += 1
    else:
        r -= 1
    ans = min(ans, max(w, r))
print(ans)
