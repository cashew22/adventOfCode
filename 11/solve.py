import re
def increment(password):
	last = len(password)
	index = 1
	while(1):
		if password[last - index] == 'z':
			password[last - index] = 'a'
			index +=1
		else:
			password[last - index] = chr(ord(password[last - index])+1)
			break
	return password

password = list("hxbxxzaa")

count = 0
match = 0
while(1):
	if len(re.findall("[iol]", "".join(password))):
		pass
	elif len(re.findall("(.)\\1", "".join(password))) < 2:
		pass
	else:
		count = 0
		for i, element in enumerate(password):
			if i == len(password) - 1:
				break
			if ord(password[i + 1]) - ord(element) == 1:
				count += 1
				if count == 2:
					print "pass"
					count = 0
					match = 1
					break
			else:
				count = 0
		if match:
			break
	password = increment(password)
print "".join(password)

