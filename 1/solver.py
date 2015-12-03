f = open('lvl.txt', 'r')
lvl = 0
char = 0
first = 0
max_l = 0
min_l = 0

for l in f:
	for c in l:
		if c == '(':
			lvl = lvl + 1
		elif c == ')':
			lvl = lvl - 1
		char = char + 1
		if lvl < 0 and first == 0:
			first = first + 1
			print "Basement at char: ", char
		if lvl > max_l:
			max_l = lvl
		if lvl < min_l:
			min_l = lvl

print lvl
print "min", min_l
print "max", max_l

	
