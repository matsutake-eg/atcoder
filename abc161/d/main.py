from collections import deque

k = int(input())

x = 0
q = deque(list(range(1, 10)))
for _ in range(k):
    x = q.popleft()
    d = x % 10
    y = 10 * x + d
    if d > 0:
        q.append(y - 1)
    q.append(y)
    if d < 9:
        q.append(y + 1)
print(x)
