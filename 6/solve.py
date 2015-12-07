import numpy as np

f = open('input.txt', 'r')
grid = np.zeros( (1000,1000), dtype=bool )
gridlumen = np.zeros( (1000,1000), dtype=int)
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
		gridlumen[start[0]:stop[0] + 1, start[1]:stop[1] + 1] += 2
	elif ops == "on":
		grid[start[0]:stop[0] + 1, start[1]:stop[1] + 1].fill(True)
		gridlumen[start[0]:stop[0] + 1, start[1]:stop[1] + 1] += 1
	else:
		grid[start[0]:stop[0] + 1, start[1]:stop[1] + 1].fill(False)
		gridlumen[start[0]:stop[0] + 1, start[1]:stop[1] + 1] -= 1
		gridlumen[gridlumen < 0] = 0

print "The number of light turn on is :", np.sum(grid)
print "The number of lumen use is :", np.sum(gridlumen)