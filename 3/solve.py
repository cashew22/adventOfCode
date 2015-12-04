f = open('input.txt', 'r')
earth = [[0 for x in range(10000)] for x in range(10000)]
y = 5000
x = 5000
delivery = 1

earth[x][y] += 1
for l in f:
	for c in l:
		#print "yo"
		if c == '<':
			x -= 1
		elif c == '>':
			x += 1
		elif c == '^':
			y += 1
		elif c == 'v':
			y -= 1
		if earth[x][y] == 0:
			delivery += 1
		earth[x][y] += 1

print delivery
