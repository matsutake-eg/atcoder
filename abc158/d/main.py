from collections import deque

s = input()
q = int(input())

d = deque()
for v in s:
    d.append(v)
r = False
for _ in range(q):
    xs = input().split()
    if xs[0] == "1":
        r = not r
    else:
        if (not r and xs[1] == "1") or (r and xs[1] == "2"):
            d.appendleft(xs[2])
        else:
            d.append(xs[2])
s = "".join(d)
print(s[::-1] if r else s)
