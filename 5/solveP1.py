import re

f = open('input.txt','r')
val = 0

for l in f:
	m = re.findall("[aeiou].*[aeiou].*[aeiou]", l)
	if m:
		m = re.findall("(.)\\1", l)
		if m:
			if l.find("ab") >= 0 or l.find("cd") >= 0 or l.find("pq") >= 0 or l.find("xy") >= 0:
				print "Bad string"
			else:
				val+=1
				print "Good string"
print val