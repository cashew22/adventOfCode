f = open('input.txt', 'r')
earth = [[0 for x in range(10000)] for x in range(10000)]
santa_y = 5000
santa_x = 5000
robot_y = 5000
robot_x = 5000
delivery = 1

earth[santa_x][santa_y] += 2
for l in f:
	for i,c in enumerate(l):
		if i%2:
			if c == '<':
				santa_x -= 1
			elif c == '>':
				santa_x += 1
			elif c == '^':
				santa_y += 1
			elif c == 'v':
				santa_y -= 1
		else:
			if c == '<':
				robot_x -= 1
			elif c == '>':
				robot_x += 1
			elif c == '^':
				robot_y += 1
			elif c == 'v':
				robot_y -= 1
		if earth[santa_x][santa_y] == 0 or earth[robot_x][robot_y] == 0:
			delivery += 1
		earth[santa_x][santa_y] += 1
		earth[robot_x][robot_y] += 1

print delivery