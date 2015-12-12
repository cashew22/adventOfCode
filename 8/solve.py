import re
f = open("input.txt", "r")
diff = 0
diff2 = 0
for l in f:
	diff += 2
	diff2 += 2
	for match in re.findall(r'(\\x..|\\\\|\\\")', l):
		if match.startswith("\\x"):
			diff += 3
		else:
			diff += 1
	for char in l.strip():
            if char == '\\' or char == '\"':
                diff2 +=1

print diff
print diff2

#THX marco