x, y, a, b, c = map(int, input().split())
p = list(map(int, input().split()))
q = list(map(int, input().split()))
r = list(map(int, input().split()))

p.sort(reverse=True)
q.sort(reverse=True)
s = p[:x] + q[:y] + r
s.sort(reverse=True)
print(sum(s[: x + y]))
