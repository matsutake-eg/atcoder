from collections import Counter

n = int(input())
s = [input() for _ in range(n)]

sc = Counter(s)
print(f"AC x {sc['AC']}")
print(f"WA x {sc['WA']}")
print(f"TLE x {sc['TLE']}")
print(f"RE x {sc['RE']}")
