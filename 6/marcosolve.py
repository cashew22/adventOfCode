import re
import numpy as np

a = np.zeros((1000,1000), dtype=bool)
b = np.zeros((1000,1000), dtype=int)

with open("input.txt") as textFile:
	for line in textFile:
		SuperPowerRegexParty = re.search(r'^(.*)( )(\d{1,3})(,)(\d{1,3})( through )(\d{1,3})(,)(\d{1,3})$',line.strip())
		command = SuperPowerRegexParty.group(1)
		x1 = int(SuperPowerRegexParty.group(3))
		y1 = int(SuperPowerRegexParty.group(5))
		x2 = int(SuperPowerRegexParty.group(7))
		y2 = int(SuperPowerRegexParty.group(9))
		if (command=="toggle"):			
			a[x1:x2+1, y1:y2+1] = np.logical_not(a[x1:x2+1, y1:y2+1])
			b[x1:x2+1, y1:y2+1] += 2
		elif (command=="turn off"):			
			a[x1:x2+1, y1:y2+1] = 0
			b[x1:x2+1, y1:y2+1] -= 1
			b[b < 0] = 0
		elif (command=="turn on"):			
			a[x1:x2+1, y1:y2+1] = 1
			b[x1:x2+1, y1:y2+1] += 1
print "Q1: There is ", np.sum(a), " lights open"
print "Q2: Total lumens for theses lights is ", np.sum(b)