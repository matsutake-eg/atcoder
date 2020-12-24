n = input()

sm = 0
for v in n:
    sm += int(v)
if sm % 9 == 0:
    print("Yes")
else:
    print("No")
