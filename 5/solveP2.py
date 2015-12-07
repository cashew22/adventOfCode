import re

f = open('input.txt','r')
val = 0

for l in f:
	if len(re.findall(r"([a-z]{2}).*\1", l)) and re.findall(r"([a-z]).\1", l):
		val +=1
print val