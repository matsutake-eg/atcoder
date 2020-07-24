n = int(input())
s = input()

ans = s.count("R") * s.count("G") * s.count("B")
for i in range(n):
    for j in range(i + 1, n):
        d = j - i
        if j + d < n and s[i] != s[j] and s[j] != s[j + d] and s[j + d] != s[i]:
            ans -= 1
print(ans)
