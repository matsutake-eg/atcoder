from collections import deque

h, w = map(int, input().split())
ch, cw = map(int, input().split())
ch -= 1
cw -= 1
dh, dw = map(int, input().split())
dh -= 1
dw -= 1
s = [input() for _ in range(h)]

dx = [0, 1, 0, -1]
dy = [1, 0, -1, 0]
dxw = [-2, -1, 0, 1, 2, -2, -1, 0, 1, 2, -2, -1, 1, 2, -2, -1, 0, 1, 2, -2, -1, 0, 1, 2]
dyw = [-2, -2, -2, -2, -2, -1, -1, -1, -1, -1, 0, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2]

q = deque()
q.append((ch, cw, 0))
ans = [[10 ** 9] * w for _ in range(h)]
while len(q) > 0:
    x, y, c = q.popleft()
    for i in range(4):
        nx, ny = x + dx[i], y + dy[i]
        if (
            nx >= 0
            and nx < h
            and ny >= 0
            and ny < w
            and c < ans[nx][ny]
            and s[nx][ny] == "."
        ):
            q.appendleft((nx, ny, c))
            ans[nx][ny] = c
    for i in range(24):
        nx, ny = x + dxw[i], y + dyw[i]
        if (
            nx >= 0
            and nx < h
            and ny >= 0
            and ny < w
            and c + 1 < ans[nx][ny]
            and s[nx][ny] == "."
        ):
            q.append((nx, ny, c + 1))
            ans[nx][ny] = c + 1
v = ans[dh][dw]
print(-1 if v == 10 ** 9 else v)
