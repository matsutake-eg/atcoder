n = int(input())
a = list(map(int, input().split()))

for v in a:
    if v % 2 == 0:
        if not (v % 3 == 0 or v % 5 == 0):
            print("DENIED")
            exit()
print("APPROVED")
