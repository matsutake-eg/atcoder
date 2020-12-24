s = input()
t = input()

ans = 0
for i, sv in enumerate(s):
    if sv != t[i]:
        ans += 1
print(ans)
