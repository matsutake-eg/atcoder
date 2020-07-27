a = [0] * 3
for i in range(3):
    a[i] = list(map(int, input().split()))
n = int(input())
b = set()
for _ in range(n):
    b.add(int(input()))

if (
    set(a[0]) <= b
    or set(a[1]) <= b
    or set(a[2]) <= b
    or set(x[0] for x in a) <= b
    or set(x[1] for x in a) <= b
    or set(x[2] for x in a) <= b
    or {a[0][0], a[1][1], a[2][2]} <= b
    or {a[0][2], a[1][1], a[2][0]} <= b
):
    print("Yes")
else:
    print("No")
