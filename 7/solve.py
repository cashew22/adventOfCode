import numpy as np
import sys

new_todo = []
wire = {}
i = 0
k = 0
last_len = 0
valfirst = 0

while k < 2:
	file = open('input.txt', 'r')
	todo = []
	if k == 1:
		todo.append("{} -> b".format(np.uint16(valfirst)))
	for line in file:
		if not line.startswith("14146") or k == 0:
			todo.append(line)

	while(len(todo) > 0):
		i += 1

		for line in todo:
		#Parse the line
			line = line.strip()
			splitted = line.split(" ")

			if len(splitted) == 3: #This is an assignation of value
				number = 1;
				try:
					val = int(splitted[0])
				except:
					number = 0;

				if not number:
					sig = splitted[0]
					key = splitted[2]

					if sig not in wire:
						new_todo.append(line)
					else:
						wire[key] = wire[sig]

				else:
					value = np.uint16(splitted[0])
					key = splitted[2]
					wire[key] = value

			elif len(splitted) == 5: #two signal and operator or a shift x AND y -> d
				sig1 = splitted[0]
				sig2 = splitted[2]
				oper = splitted[1]
				key = splitted[4]

				if oper == "AND":
					special = 1
					try:
						val = np.uint16(sig1)
					except:
						special = 0
					if special:
						if sig2 not in wire:
							new_todo.append(line)
						else:
							wire[key] = np.uint16(sig1) & wire[sig2]
					else:
						if sig1 not in wire or sig2 not in wire:
							new_todo.append(line)
						else:
							wire[key] = wire[sig1] & wire[sig2]
				elif oper == "OR":
					if sig1 not in wire or sig2 not in wire:
						new_todo.append(line)
					else:
						wire[key] = wire[sig1] | wire[sig2]
				elif oper == "LSHIFT":
					if sig1 not in wire:
						new_todo.append(line)
					else:
						wire[key] = wire[sig1] << np.uint16(sig2)
				elif oper == "RSHIFT":
					if sig1 not in wire:
						new_todo.append(line)
					else:
						wire[key] = wire[sig1] >> np.uint16(sig2)

			else: #NOT
				sig = splitted[1]
				key = splitted[3]

				if sig not in wire:
					new_todo.append(line)
				else:
					wire[key] = ~wire[sig]
		todo = new_todo[:]
		new_todo = []

	if k == 0:
		valfirst = wire['a']
		print ("The answer for first question is a = {}").format(wire['a'])
		wire.clear()
	k += 1
print ("The answer for second question is a = {}").format(wire['a'])