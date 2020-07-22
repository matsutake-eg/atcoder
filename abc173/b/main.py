n = int(input())
a, w, t, r = 0, 0, 0, 0
for i in range(n):
    s = input()
    if s == "AC":
        a += 1
    elif s == "WA":
        w += 1
    elif s == "TLE":
        t += 1
    else:
        r += 1

print(f"AC x {a}")
print(f"WA x {w}")
print(f"TLE x {t}")
print(f"RE x {r}")
