import numpy as np

f = open('input.txt', 'r')
grid = np.zeros( (1000,1000), dtype=bool )
index = 0

for l in f:
	temp = l.split(" ")
	if temp[0] == "turn":
		ops = temp[1]
		start = temp[2].split(",")
		start = map(int, start)
		stop = temp[4].split(",")
		stop = map(int, stop)
	else:
		ops = "toggle"
		start = temp[1].split(",")
		start = map(int, start)
		stop = temp[3].split(",")
		stop = map(int, stop)

	if ops == "toggle":
		grid[start[0]:stop[0] + 1, start[1]:stop[1] + 1] = np.invert(grid[start[0]:stop[0] + 1, start[1]:stop[1] + 1], dtype=bool)
	elif ops == "on":
		mask = grid[start[0]:stop[0] + 1, start[1]:stop[1] + 1]
		mask.fill(True)
	else:
		mask = grid[start[0]:stop[0] + 1, start[1]:stop[1] + 1]
		mask.fill(False)

print np.sum(grid)