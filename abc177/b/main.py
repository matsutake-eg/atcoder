s = input()
t = input()

m = 0
for i in range(len(s) - len(t) + 1):
    b = s[i : i + len(t)]
    cnt = 0
    for j, v in enumerate(b):
        if v == t[j]:
            cnt += 1
    m = max(m, cnt)
print(len(t) - m)
