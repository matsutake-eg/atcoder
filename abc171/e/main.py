from functools import reduce

n = int(input())
a = list(map(int, input().split()))

b = reduce(lambda ac, x: ac ^ x, a)
print(*map(lambda x: x ^ b, a))
