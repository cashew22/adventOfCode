from itertools import product
import collections

class path:
	def __init__(self, townA, townB, distance):
		self.townA = townA
		self.townB = townB
		self.distance = distance

inputpath = []
townlist = []
pathlist = []

total = 0
seen = set()
f = open("input.txt",'r')

for l in f:
	townA, scrap, townB, scrap2, distance = l.split(" ")
	inputpath.append(path(townA, townB, int(distance)))

	if townA in townlist:
		pass
	else:
		townlist.append(townA)
	if townB in townlist:
		pass
	else:
		townlist.append(townB)

distance = 0
shortest = 50000
longest = 0
for i in product(townlist, repeat=len(townlist)):
	if len(i) == len(set(i)):
		pathlist.append(i)

for path in pathlist:
	distance = 0
	for i, town in enumerate(path):
		if i == len(path) - 1:
			break
		for x in inputpath:
			if x.townA == town or x.townB == town:
				if x.townA == path[i + 1] or x.townB == path[i + 1]:
					distance += x.distance
					break

	if shortest > distance:
		shortest = distance
	if longest < distance:
		longest = distance
print shortest
print longest
