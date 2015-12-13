number = "1113122113"
count = 0
newnumber = ""
last = ''
constant = 1.303577269
digit = 10
for i in range(0,50):
	for c in number:
		if last == c:
			count += 1
		else:
			if count != 0:
				newnumber += str(count) + last
			last = c
			count = 1
	if count != 0:
			newnumber += str(count) + last
			last = ''
			count = 0
	number = newnumber
	newnumber = ""
	digit *= constant
print len(number)
print digit
